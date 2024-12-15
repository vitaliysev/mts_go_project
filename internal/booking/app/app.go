package app

import (
	"context"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/vitaliysev/mts_go_project/internal/booking/api/booking_http"
	"github.com/vitaliysev/mts_go_project/internal/booking/closer"
	"github.com/vitaliysev/mts_go_project/internal/booking/config"
	"github.com/vitaliysev/mts_go_project/internal/booking/logger"
	metric "github.com/vitaliysev/mts_go_project/internal/booking/metrics"
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	_ "github.com/vitaliysev/mts_go_project/statik/booking/statik"
	"io"

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
	swaggerServer   *http.Server
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
	//topic := "message-sending"
	//brokers := []string{"localhost:19092"}
	//
	////adm, err := admin.New(brokers)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer adm.Close()
	//ok, err := adm.TopicExists(topic)
	//if !ok {
	//	err = adm.CreateTopic(topic)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
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
	go func() {
		defer wg.Done()

		err := a.runSwaggerServer()
		if err != nil {
			log.Fatalf("failed to run Swagger server: %v", err)
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
		a.initSwaggerServer,
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

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/api.booking.swagger.json", serveSwaggerFile("/api.booking.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:    a.serviceProvider.SwaggerConfig().Address(),
		Handler: mux,
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Serving swagger file: " + path)

		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Open swagger file: %s", path)

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		log.Printf("Read swagger file: %s", path)

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Write swagger file: %s", path)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Served swagger file: %s", path)
	}
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

func (a *App) runSwaggerServer() error {
	log.Printf("Swagger server is running on %s", a.serviceProvider.SwaggerConfig().Address())

	err := a.swaggerServer.ListenAndServe()
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

// @title Booking API
// @version 1.0
// @description This API provides booking-related operations such as creating a booking, listing clients, etc.
// @host localhost:8081
// @BasePath /
// @securityDefenitions.apikey ApiKeyAuth
// @in header
// @name auth
func (a *App) initHTTPServer(ctx context.Context) error {
	mux := http.NewServeMux()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

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
		a.handleAuth(w, r, bookingHandler)
	})

	mux.HandleFunc("/booking/v1/login", func(w http.ResponseWriter, r *http.Request) {
		a.handleAccess(w, r, bookingHandler)
	})

	mux.HandleFunc("/booking/v1/getrefr", func(w http.ResponseWriter, r *http.Request) {
		a.handleGetRefr(w, r, bookingHandler)
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

// @description GetBookingRequest contains a info for get
type GetBookingRequest struct {
	ID           int64  `json:"id"`
	Access_token string `json:"access_token"`
	Path         string
}

// handleBookingGetCl Получение бронирований отелей клиента.
// @Summary Получение бронирований отелей клиента
// @SecurityApiKeyAuth
// @Description Получение бронирований отелей клиента используя HTTP API.
// @Tags Booking
// @Accept json
// @Produce json
// @Param bookingBody body GetBookingRequest true "Booking Data"
// @Success 200 {object} GetBookingResponse "Bookings listed successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /booking/v1/listCl [post]
func (a *App) handleBookingGetCl(w http.ResponseWriter, r *http.Request, path string, handler *booking_http.Implementation) {
	if r.Method == http.MethodPost {
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

// handleBookingGetHo Получение бронирований отелей владельца.
// @Summary Получение бронирований отелей владельца
// @SecurityApiKeyAuth
// @Description Получение бронирований отелей владельца используя HTTP API.
// @Tags Booking
// @Accept json
// @Produce json
// @Param bookingBody body GetBookingRequest true "Booking Data"
// @Success 200 {object} GetBookingResponse "Bookings listed successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /booking/v1/listHo [post]
func (a *App) handleBookingGetHo(w http.ResponseWriter, r *http.Request, path string, handler *booking_http.Implementation) {
	if r.Method == http.MethodPost {
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

// @description GetBookingResponse contains a list of booking information.
type GetBookingResponse struct {
	Info []*model.Book `json:"info"`
}
type CreateBookingRequest struct {
	Info         model.BookInfo `json:"info"`
	Access_token string         `json:"access_token"`
}

// handleBookingCreate Создание нового бронирования.
// @Summary Создание нового бронирования
// @SecurityApiKeyAuth
// @Description Создание нового бронирования используя HTTP API.
// @Tags Booking
// @Accept json
// @Produce json
// @Param bookingBody body CreateBookingRequest true "Booking Data"
// @Success 200 {object} CreateBookingResponse "Booking created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /booking/v1/create [post]
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

type CreateBookingResponse struct {
	ID       int64  `json:"id"`
	Cost     int64  `json:"cost"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Period   int64  `json:"period"`
}

// ErrorResponse represents a standard error response.
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Invalid request"`
}

type SigninClientRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SigninClientResponse struct {
	Refresh_token string `json:"refresh_token"`
}

// handleAuth Регистрация.
// @Summary Регистрация
// @Description Регистрация используя HTTP API.
// @Tags Booking
// @Accept json
// @Produce json
// @Param bookingBody body SigninClientRequest true "Auth Data"
// @Success 200 {object} SigninClientResponse "Authentication is succesful"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /booking/v1/signin [post]
func (a *App) handleAuth(w http.ResponseWriter, r *http.Request, handler *booking_http.Implementation) {

	if r.Method == http.MethodPost {
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
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type GetRefreshTokenRequest struct {
	Refresh_token string `json:"refresh_token"`
}

type GetRefreshTokenResponse struct {
	Refresh_token string `json:"refresh_token"`
}

// handleAuth Обновление RefreshToken.
// @Summary Обновление RefreshToken
// @SecurityApiKeyAuth
// @Description Обновление RefreshToken используя HTTP API.
// @Tags Booking
// @Accept json
// @Produce json
// @Param bookingBody body GetRefreshTokenRequest true "Auth Data"
// @Success 200 {object} GetRefreshTokenResponse "Refresh token got succesfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /booking/v1/getrefr [patch]
func (a *App) handleGetRefr(w http.ResponseWriter, r *http.Request, handler *booking_http.Implementation) {
	if r.Method == http.MethodPatch {
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
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleAccess Вход.
// @Summary Обновление RefreshToken
// @SecurityApiKeyAuth
// @Description Вход используя HTTP API.
// @Tags Booking
// @Accept json
// @Produce json
// @Param bookingBody body LoginClientRequest true "Access Data"
// @Success 200 {object} LoginClientResponse "Access if succesful"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /booking/v1/login [post]
func (a *App) handleAccess(w http.ResponseWriter, r *http.Request, handler *booking_http.Implementation) {
	if r.Method == http.MethodPost {
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
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type LoginClientRequest struct {
	Username      string `json:"login"`
	Password      string `json:"password"`
	Refresh_token string `json:"refresh_token"`
}

type LoginClientResponse struct {
	Access_token string `json:"access_token"`
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
