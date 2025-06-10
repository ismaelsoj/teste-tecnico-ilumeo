package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"teste-tecnico-ilumeo/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL não está definido")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer dbpool.Close()

	historicoHandler := handlers.NewHistoricoHandler(dbpool)

	// Criar router
	r := chi.NewRouter()

	// Middlewares básicos
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Rotas
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/historico", historicoHandler.HandleHistorico)
		r.Get("/taxa-conversao", historicoHandler.HandleTaxaConversao)
	})

	/*http.HandleFunc("/historico", historicoHandler.HandleHistorico)
	http.HandleFunc("/taxa-conversao", historicoHandler.HandleTaxaConversao)*/

	port := "8080"
	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
