package main

import (
	"AIAgentsToolHub/internal/data/store"
	"flag"
	"log/slog"
	"os"
)

type config struct {
	port int

	db struct {
		dsn string
	}
}

type application struct {
	config config
	logger *slog.Logger
}

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

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
		config: cfg,
	}

	db, err := store.InitDB(app.config.db.dsn)
	if err != nil {
		app.LogError(err.Error())
	}

	defer db.Close()

	tools, err := store.SearchTools(db, "  multiply")
	if err != nil {
		app.LogError(err.Error())
	}

	for _, t := range tools {
		app.logger.Info("Tool found",
			"Name", t.Name,
			"Description", t.Description,
			"Input", t.Input,
			"Output", t.Output,
			"Score", t.Score,
		)
	}

}
