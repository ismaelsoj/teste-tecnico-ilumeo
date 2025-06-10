package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"teste-tecnico-ilumeo/internal/handlers"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL não está definido")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer dbpool.Close()

	historicoHandler := handlers.NewHistoricoHandler(dbpool)

	http.HandleFunc("/historico", historicoHandler.HandleHistorico)
	http.HandleFunc("/taxa-conversao", historicoHandler.HandleTaxaConversao)

	port := "8080"
	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
