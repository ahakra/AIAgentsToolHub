package toolrepo

import (
	"AIAgentsToolHub/internal/data/model"
	"context"
	"database/sql"
	"fmt"
)

type ToolRepo struct {
	db *sql.DB
}

func NewToolRepo(db *sql.DB) *ToolRepo {
	return &ToolRepo{
		db: db,
	}
}

func (tr *ToolRepo) GetToolByDescription(ctx context.Context, description string) ([]model.Tool, error) {

	rows, err := tr.db.QueryContext(ctx, "SELECT tool_id, tool_name, tool_description, tool_input, tool_output, bm25(tools) as score FROM tools WHERE tools MATCH ?", description)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}
	defer rows.Close()

	var results []model.Tool
	for rows.Next() {
		var t model.Tool
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Input, &t.Output, &t.Score); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		results = append(results, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return results, nil

}
