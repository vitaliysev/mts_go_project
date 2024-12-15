package main

import (
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/notification/consumer"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	topic := "message-sending"
	brokers := []string{"localhost:19092"}
	c, err := consumer.New(brokers, topic)
	if err != nil {
		log.Fatalln("Failed to create consumer:", err)
	}
	defer c.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			c.PrintMessages()
		}
	}()

	sig := <-sigChan
	fmt.Printf("Получен сигнал %v. Завершаем работу...\n", sig)
}
