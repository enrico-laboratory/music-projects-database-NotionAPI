package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (app *Application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Route("/music-projects", func(r chi.Router) {
		r.Get("/", app.GetMusicProjects)
		r.Get("/{status}", app.GetMusicProjectsByStatus)
	})

	mux.Get("/music-project/{projectId}", app.GetMusicProjectById)
	mux.Get("/repertoire/music-project/{projectId}", app.GetRepertoireByProjectId)
	mux.Get("/cast/music-project/{projectId}", app.GetCastByProjectId)
	mux.Get("/schedule/music-project/{projectId}", app.GetScheduleByProjectId)
	mux.Get("/location/{locationId}", app.GetLocationById)

	return mux
}
