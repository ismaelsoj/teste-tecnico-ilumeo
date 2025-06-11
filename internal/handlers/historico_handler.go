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

func (h *HistoricoHandler) HandleTaxaConversao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	historicos, err := h.repo.GetTaxaConversaoPorCanalETempo(ctx)
	if err != nil {
		http.Error(w, "Erro ao buscar dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historicos)
}

func (h *HistoricoHandler) HandleTaxaConversaoFiltrada(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	canal := r.URL.Query().Get("canal")

	strStart := r.URL.Query().Get("start")
	var start *time.Time
	var end *time.Time
	if strStart != "" {
		parsedStart, errStart := time.Parse("YYYY-MM-ddTHH:mm:ss", strStart)
		if errStart != nil {
			http.Error(w, "Data Inválida: "+errStart.Error(), http.StatusInternalServerError)
			// return
		}
		start = &parsedStart

	}

	strEnd := r.URL.Query().Get("end")
	if strEnd != "" {
		parsedEnd, errEnd := time.Parse("YYYY-MM-ddTHH:mm:ss", strEnd)
		if errEnd != nil {
			http.Error(w, "Data Inválida: "+errEnd.Error(), http.StatusInternalServerError)
			// return
		}
		end = &parsedEnd
	}

	historicos, err := h.repo.GetTaxaConversaoFiltrada(ctx, canal, start, end)
	if err != nil {
		http.Error(w, "Erro ao buscar dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historicos)
}
