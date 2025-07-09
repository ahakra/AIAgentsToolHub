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

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create an FTS5 virtual table
	// This to enable full text search
	createTable := `
	CREATE VIRTUAL TABLE IF NOT EXISTS tools USING fts5(
		tool_id ,
		tool_name,
		tool_location,
		tool_description,
		tool_input,
		tool_output,
		tool_type
	);
	`

	_, err = db.Exec(createTable)
	if err != nil {
		return nil, fmt.Errorf("error creating FTS5 table: %w", err)
	}

	// Check if table is empty
	// To init db with some tools
	var count int
	err = db.QueryRow("SELECT count(*) FROM tools;").Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("count error: %w", err)
	}

	if count == 0 {
		log.Println("Seeding tools table...")

		seed_add := `
		INSERT INTO tools (tool_id,tool_name, tool_location, tool_description, tool_input, tool_output, tool_type)
		VALUES (?,?, ?, ?, ?, ?, ?);
		`

		_, err := db.Exec(seed_add,
			uuid.New(),
			"add",
			"./bin/add",
			"adds two integers together",
			`{"x": "int", "y": "int"}`,
			`{"sum": "int"}`,
			model.CLITool,
		)
		if err != nil {
			return nil, fmt.Errorf("seeding error: %w", err)
		}
		seed_multiply := `
		INSERT INTO tools (tool_id,tool_name, tool_location, tool_description, tool_input, tool_output, tool_type)
		VALUES (?,?, ?, ?, ?, ?, ?);
		`

		_, err = db.Exec(seed_multiply,
			uuid.New(),
			"multiply",
			"./bin/multiply",
			"multiply two integers together",
			`{"x": "int", "y": "int"}`,
			`{"multiply": "int"}`,
			model.CLITool,
		)
		if err != nil {
			return nil, fmt.Errorf("seeding error: %w", err)
		}
	}

	return db, nil
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
