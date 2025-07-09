package main

import (
	"AIAgentsToolHub/internal/data/store"
	"AIAgentsToolHub/internal/data/toolrepo"
	"AIAgentsToolHub/internal/service"
	"flag"
	"log/slog"
	"os"
)

type config struct {
	port int

	db struct {
		dsn string
	}
	jsonConfig struct {
		maxByte            int
		allowUnknownFields bool
	}
}

type application struct {
	config      config
	logger      *slog.Logger
	toolService *service.ToolService
}
type responseData map[string]any

func main() {
	// input := map[string]interface{}{
	// 	"x": 5,
	// 	"y": 8,
	// }

	// output, err := runner.CLIRunTool("./bin/add", input)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Result:", output)

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "./database.db", "SQLITE3 DSN")

	flag.IntVar(&cfg.jsonConfig.maxByte, "Maxbytes", 1_048_576, "MAX Bytes for JSON Body")
	flag.BoolVar(&cfg.jsonConfig.allowUnknownFields, "AllowUnknownFields", false, "Allow unknown fields in JSON Body")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
		config: cfg,
	}

	db, err := store.InitDB(app.config.db.dsn)
	if err != nil {
		app.LogError("failed to init db", err)
	}

	defer db.Close()

	toolRepo := toolrepo.NewToolRepo(db)
	toolService := service.NewToolService(toolRepo)

	app.toolService = &toolService

}
