package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"html/template"
	"os"
	"snippetbox.ktykhanskyi.net/internal/models"
	_"github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	dsn := flag.String("dsn", "root:reset123@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL Data source name")
	flag.Parse()


	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug,}))

	db, err := openDB(*dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
