package repository

import (
	"context"
	"strconv"
	"strings"
	"teste-tecnico-ilumeo/internal/models"
	"time"

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

func (r *HistoricoRepository) GetTaxaConversaoPorCanalETempo(ctx context.Context) ([]models.TaxaConversaoPorCanalETempo, error) {
	query := `
        SELECT
            DATE(created_at) AS data,
            origin AS canal,
            (SUM(CASE WHEN response_status_id = 6 THEN 1 ELSE 0 END)::float / COUNT(*)) * 100 AS taxa_conversao
        FROM historico
        GROUP BY data, canal
        ORDER BY data ASC, canal ASC
    `

	rows, err := r.dbpool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []models.TaxaConversaoPorCanalETempo

	for rows.Next() {
		var item models.TaxaConversaoPorCanalETempo
		var data time.Time

		err := rows.Scan(&data, &item.Canal, &item.TaxaConversao)
		if err != nil {
			return nil, err
		}

		item.Data = data
		resultados = append(resultados, item)
	}

	return resultados, nil
}

func (r *HistoricoRepository) GetTaxaConversaoFiltrada(ctx context.Context, canal string, start, end *time.Time) ([]models.TaxaConversaoPorCanalETempo, error) {
	query := `
		SELECT
			DATE(created_at) AS data,
			origin AS canal,
			(SUM(CASE WHEN response_status_id = 1 THEN 1 ELSE 0 END)::float / COUNT(*)) * 100 AS taxa_conversao
		FROM historico
	`

	// Montar cláusulas WHERE dinâmicas
	var conditions []string
	var args []interface{}
	argId := 1

	if canal != "" {
		conditions = append(conditions, "origin = $"+strconv.Itoa(argId))
		args = append(args, canal)
		argId++
	}
	if start != nil {
		conditions = append(conditions, "created_at >= $"+strconv.Itoa(argId))
		args = append(args, *start)
		argId++
	}
	if end != nil {
		conditions = append(conditions, "created_at <= $"+strconv.Itoa(argId))
		args = append(args, *end)
		argId++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += `
		GROUP BY data, canal
		ORDER BY data ASC, canal ASC
	`

	rows, err := r.dbpool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []models.TaxaConversaoPorCanalETempo

	for rows.Next() {
		var item models.TaxaConversaoPorCanalETempo
		var data time.Time

		err := rows.Scan(&data, &item.Canal, &item.TaxaConversao)
		if err != nil {
			return nil, err
		}

		item.Data = data
		resultados = append(resultados, item)
	}

	return resultados, nil
}
