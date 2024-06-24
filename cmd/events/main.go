package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "github.com/vinicius3g/golang/internal/events/infra/http"
	"github.com/vinicius3g/golang/internal/events/infra/service"
	"github.com/vinicius3g/golang/internal/events/infra/service/repository"
	"github.com/vinicius3g/golang/internal/events/usecase"
)

func main() {
	// Configuração do banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Repositório
	eventRepo, err := repository.NewMySQLEventRepository(db)
	if err != nil {
		panic(err)
	}
	// URLs base específicas para cada parceiro
	partnerBaseURLs := map[int]string{
		1: "https://partner1.com",
		2: "https://partner2.com",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventUsecase := usecase.NewListEventsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	buyTicketUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)
	createSpotsUseCase := usecase.NewCreateSpotsUseCase(eventRepo)
	createEventUseCase := usecase.NewCreateEventUseCase(eventRepo)

	// Handlers HTTP
	eventsHandler := httpHandler.NewEventHandler(
		listEventUsecase,
		listSpotsUseCase,
		getEventUseCase,
		buyTicketUseCase,
		createSpotsUseCase,
		createEventUseCase,
	)

	r := http.NewServeMux()
	// r.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /events", eventsHandler.CreateEvent)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)
	r.HandleFunc("POST /events/{eventID}/spots", eventsHandler.CreateSpots)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Canal para escutar sinais do sistema operacional
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Recebido sinal de interrupção, iniciando o graceful shutdown
		log.Println("Recebido sinal de interrupção, iniciando o graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Erro no graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	// Iniciando o servidor HTTP
	log.Println("Servidor HTTP rodando na porta 8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("Servidor HTTP finalizado")
}
