package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"teste-tecnico-ilumeo/internal/repository"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HistoricoHandler struct {
	repo *repository.HistoricoRepository
}

func NewHistoricoHandler(dbpool *pgxpool.Pool) *HistoricoHandler {
	return &HistoricoHandler{
		repo: repository.NewHistoricoRepository(dbpool),
	}
}

func (h *HistoricoHandler) HandleHistorico(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	historicos, err := h.repo.BuscarHistorico(ctx, 1000)
	if err != nil {
		http.Error(w, "Erro ao buscar dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historicos)
}
