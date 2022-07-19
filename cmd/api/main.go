package main

import (
	"errors"
	"fmt"
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/notion"
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/services"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serverPort    = 5002
	serverNetwork = "tcp"
)

type config struct {
	server struct {
		port    int
		network string
	}
	env string
}

type Application struct {
	config        config
	service       services.Service
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func main() {
	var cfg config
	cfg.server.port = serverPort
	cfg.server.network = serverNetwork
	cfg.env = "dev"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &Application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	notionToken, err := getNotionToken()
	if err != nil {
		log.Fatal(err)
	}

	nc, err := notion.NewClient(notionToken)
	if err != nil {
		log.Fatal(err)
	}

	s := services.NewService(nc)
	app.service = s

	// TEST
	//var model []parsedmodels.MusicProject
	//
	//s.GetMusicProjects(&model)
	//
	//log.Println(model)

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}

func getNotionToken() (string, error) {
	notionToken := os.Getenv("NOTION_TOKEN")
	if notionToken == "" {
		return "", errors.New("notion token is empty")
	}

	return notionToken, nil
}

func (app *Application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.server.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Println(fmt.Sprintf("Starting HTTP server in %s mode on port %d", app.config.env, app.config.server.port))

	return srv.ListenAndServe()
}
