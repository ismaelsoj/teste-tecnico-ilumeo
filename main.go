package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Historico representa um registro da tabela
type Historico struct {
	ID               int64  `json:"id"`
	Origin           string `json:"origin"`
	ResponseStatusID int    `json:"response_status_id"`
}

func main() {
	// Pega a string de conexão do ambiente
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL não está definido")
	}

	// Cria o pool de conexões sem um contexto com timeout (ctx global não recomendado)
	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer dbpool.Close()

	// Roteia o endpoint
	http.HandleFunc("/historico", func(w http.ResponseWriter, r *http.Request) {
		// Contexto por request com timeout de 10 segundos
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		historicos, err := buscarHistorico(ctx, dbpool)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao buscar dados: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(historicos)
	})

	// Sobe o servidor HTTP
	port := "8080"
	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// buscarHistorico consulta todos os registros da tabela historico
func buscarHistorico(ctx context.Context, dbpool *pgxpool.Pool) ([]Historico, error) {
	rows, err := dbpool.Query(ctx, "SELECT id, origin, response_status_id FROM historico LIMIT 1000")
	if err != nil {
		log.Printf("Erro na query: %s", err)
		return nil, err
	}
	defer rows.Close()

	var resultados []Historico

	for rows.Next() {
		var h Historico
		err := rows.Scan(&h.ID, &h.Origin, &h.ResponseStatusID)
		if err != nil {
			return nil, err
		}
		resultados = append(resultados, h)
	}

	return resultados, nil
}
