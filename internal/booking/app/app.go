package app

import (
	"context"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitaliysev/mts_go_project/internal/booking/api/booking_http"
	"github.com/vitaliysev/mts_go_project/internal/booking/closer"
	"github.com/vitaliysev/mts_go_project/internal/booking/config"
	"github.com/vitaliysev/mts_go_project/internal/booking/logger"
	metric "github.com/vitaliysev/mts_go_project/internal/booking/metrics"
	"github.com/vitaliysev/mts_go_project/internal/booking/redpanda/admin"
	"github.com/vitaliysev/mts_go_project/internal/tracing"

	//	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	desc "github.com/vitaliysev/mts_go_project/pkg/booking_v1"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	topic := "message-sending"
	brokers := []string{"localhost:19092"}

	adm, err := admin.New(brokers)
	if err != nil {
		log.Fatal(err)
	}
	defer adm.Close()
	ok, err := adm.TopicExists(topic)
	if !ok {
		err = adm.CreateTopic(topic)
		if err != nil {
			log.Fatal(err)
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		err := runPrometheus()
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("failed to run GRPC server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
		a.initTracing,
		a.initMetrics,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
func (a *App) initMetrics(ctx context.Context) error {
	err := metric.Init(ctx)
	return err
}
func (a *App) initTracing(ctx context.Context) error {
	err := tracing.NewTracer("http://localhost:14268/api/traces", "Booking-service")
	return err
}
func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}
func runPrometheus() error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	prometheusServer := &http.Server{
		Addr:    "localhost:2300",
		Handler: mux,
	}

	log.Printf("Prometheus server is running on %s", "localhost:2300")

	err := prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterBookingV1Server(a.grpcServer, a.serviceProvider.GRPCBookingImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := http.NewServeMux()

	// Используем реализацию из booking_http
	bookingHandler := a.serviceProvider.HTTPBookingImpl(ctx)

	mux.HandleFunc("/booking/v1/create", func(w http.ResponseWriter, r *http.Request) {
		a.handleBookingCreate(w, r, bookingHandler)
	})

	mux.HandleFunc("/booking/v1/listCl", func(w http.ResponseWriter, r *http.Request) {
		a.handleBookingGetCl(w, r, "/booking/v1/listCl", bookingHandler)
	})

	mux.HandleFunc("/booking/v1/listHo", func(w http.ResponseWriter, r *http.Request) {
		a.handleBookingGetHo(w, r, "/booking/v1/listHo", bookingHandler)
	})

	mux.HandleFunc("/booking/v1/signin", func(w http.ResponseWriter, r *http.Request) {
		a.handleAuthAccess(w, r, bookingHandler)
	})

	mux.HandleFunc("/booking/v1/login", func(w http.ResponseWriter, r *http.Request) {
		a.handleAuthAccess(w, r, bookingHandler)
	})

	mux.HandleFunc("/booking/v1/getrefr", func(w http.ResponseWriter, r *http.Request) {
		a.handleAuthAccess(w, r, bookingHandler)
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: mux,
	}

	return nil
}

func (a *App) handleBookingGetCl(w http.ResponseWriter, r *http.Request, path string, handler *booking_http.Implementation) {
	if r.Method == http.MethodGet {
		var req booking_http.GetBookingRequest

		json.NewDecoder(r.Body).Decode(&req)
		resp, err := handler.Get(r.Context(), &req, path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) handleBookingGetHo(w http.ResponseWriter, r *http.Request, path string, handler *booking_http.Implementation) {
	if r.Method == http.MethodGet {
		var req booking_http.GetBookingRequest

		json.NewDecoder(r.Body).Decode(&req)
		resp, err := handler.Get(r.Context(), &req, path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) handleBookingCreate(w http.ResponseWriter, r *http.Request, handler *booking_http.Implementation) {
	if r.Method == http.MethodPost {
		var req booking_http.CreateBookingRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Вызов метода создания через HTTP-сервис
		resp, err := handler.Create(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Формирование успешного ответа
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) handleAuthAccess(w http.ResponseWriter, r *http.Request, handler *booking_http.Implementation) {
	switch r.Method {
	case http.MethodGet:
		var req booking_http.LoginClientRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Вызов метода создания через HTTP-сервис
		resp, err := handler.Login(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Формирование успешного ответа
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	case http.MethodPost:
		var req booking_http.SigninClientRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Вызов метода создания через HTTP-сервис
		resp, err := handler.Signin(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Формирование успешного ответа
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	case http.MethodPatch:
		var req booking_http.GetRefreshTokenRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Вызов метода создания через HTTP-сервис
		resp, err := handler.GetRefresh(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Формирование успешного ответа
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.httpServer.Addr)
	logger.Init(getCore(getAtomicLevel()))
	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
