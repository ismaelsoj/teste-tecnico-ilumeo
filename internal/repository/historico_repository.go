package repository

import (
	"context"
	"teste-tecnico-ilumeo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HistoricoRepository struct {
	dbpool *pgxpool.Pool
}

func NewHistoricoRepository(dbpool *pgxpool.Pool) *HistoricoRepository {
	return &HistoricoRepository{dbpool: dbpool}
}

func (r *HistoricoRepository) BuscarHistorico(ctx context.Context, limit int) ([]models.Historico, error) {
	query := `
        SELECT id, origin, response_status_id, created_at
        FROM historico
        ORDER BY created_at DESC
        LIMIT $1
    `

	rows, err := r.dbpool.Query(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []models.Historico

	for rows.Next() {
		var h models.Historico
		err := rows.Scan(&h.ID, &h.Origin, &h.ResponseStatusID, &h.CreatedAt)
		if err != nil {
			return nil, err
		}
		resultados = append(resultados, h)
	}

	return resultados, nil
}
