package store

import (
	"AIAgentsToolHub/internal/data/model"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/knaka/go-sqlite3-fts5"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := connectDB(dbPath)
	if err != nil {
		return nil, err
	}

	if err := createFTSTable(db); err != nil {
		return nil, err
	}

	if err := seedToolsIfEmpty(db); err != nil {
		return nil, err
	}

	return db, nil
}

func connectDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to DB: %w", err)
	}
	return db, nil
}

func createFTSTable(db *sql.DB) error {
	createTable := `
	CREATE VIRTUAL TABLE IF NOT EXISTS tools USING fts5(
		tool_id,
		tool_name,
		tool_location,
		tool_description,
		tool_input,
		tool_output,
		tool_type
	);
	`
	_, err := db.Exec(createTable)
	if err != nil {
		return fmt.Errorf("error creating FTS5 table: %w", err)
	}
	return nil
}

func seedToolsIfEmpty(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT count(*) FROM tools;").Scan(&count)
	if err != nil {
		return fmt.Errorf("count error: %w", err)
	}
	if count > 0 {
		return nil
	}

	log.Println("Seeding tools table...")

	seedSQL := `
	INSERT INTO tools (tool_id, tool_name, tool_location, tool_description, tool_input, tool_output, tool_type)
	VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	tools := []struct {
		Name, Location, Description, Input, Output string
	}{
		{"add", "./bin/add", "adds two integers together", `{"x": "int", "y": "int"}`, `{"sum": "int"}`},
		{"multiply", "./bin/multiply", "multiply two integers together", `{"x": "int", "y": "int"}`, `{"multiply": "int"}`},
	}

	for _, tool := range tools {
		_, err := db.Exec(seedSQL,
			uuid.New().String(),
			tool.Name,
			tool.Location,
			tool.Description,
			tool.Input,
			tool.Output,
			model.CLITool,
		)
		if err != nil {
			return fmt.Errorf("seeding tool %s failed: %w", tool.Name, err)
		}
	}

	return nil
}

func SearchTools(db *sql.DB, query string) ([]model.Tool, error) {

	rows, err := db.Query("SELECT tool_name, tool_description, tool_input, tool_output, bm25(tools) as score FROM tools WHERE tools MATCH ?", query)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}
	defer rows.Close()

	var results []model.Tool
	for rows.Next() {
		var t model.Tool
		if err := rows.Scan(&t.Name, &t.Description, &t.Input, &t.Output, &t.Score); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		results = append(results, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return results, nil
}
