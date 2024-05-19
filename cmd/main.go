package main

import (
	"context"
	exchangeratenotifierapi "exchange-rate-notifier-api"
	"exchange-rate-notifier-api/pkg/client"
	"exchange-rate-notifier-api/pkg/handler"
	"exchange-rate-notifier-api/pkg/repository"
	"exchange-rate-notifier-api/pkg/scheduler"
	"exchange-rate-notifier-api/pkg/service"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Exchange rate notifier API
// @version 1.0
// @description API server for notifying exchange rate
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	dbConfig := repository.DBConfig{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Username:   os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("DB_SSL_MODE"),
		DriverName: os.Getenv("DB_DRIVER_NAME"),
	}
	mailConfig := service.MailConfig{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
	db, err := repository.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories, mailConfig)
	clients := client.NewClient()
	handlers := handler.NewHandler(services)
	scheduler.NewExchangeRateNotificationScheduler(repositories.Subscription, clients.ExchangeRate, services.Mail).StartJob()
	server := new(exchangeratenotifierapi.Server)
	go func() {
		if err := server.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("failed to start server: %s", err.Error())
		}
	}()

	log.Print("Exchange rate notifier started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Exchange rate notifier shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("error occurred while shutting down server: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Errorf("error occurred while closing db connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
