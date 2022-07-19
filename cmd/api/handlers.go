package main

import (
	"encoding/json"
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/models/parsedmodels"
	"github.com/go-chi/chi"
	"net/http"
	"unsafe"
)

func (app *Application) GetMusicProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var model []parsedmodels.MusicProject

	err := app.service.GetMusicProjects(&model)
	if err != nil {
		app.errorLog.Println(err)
	}

	if len(model) == 0 {
		app.ClientError(w, http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(model)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) GetMusicProjectsByStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status := chi.URLParam(r, "status")
	var parsedStatus string
	switch status {
	case "ongoing":
		parsedStatus = "On Going"
	case "cancelled":
		parsedStatus = "Cancelled"
	case "completed":
		parsedStatus = "Completed"
	default:
		app.ClientError(w, http.StatusForbidden)
		return
	}

	var musicProjects []parsedmodels.MusicProject

	err := app.service.GetMusicProjectsByStatus(parsedStatus, &musicProjects)
	if err != nil {
		app.errorLog.Println(err)
	}

	if len(musicProjects) == 0 {
		app.ClientError(w, http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(musicProjects)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) GetMusicProjectById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var musicProject parsedmodels.MusicProject

	projectId := chi.URLParam(r, "projectId")
	err := app.service.GetMusicProjectById(projectId, &musicProject)
	if err != nil {
		app.errorLog.Println(err)
	}

	if unsafe.Sizeof(musicProject) == 0 {
		app.ClientError(w, http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(musicProject)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) GetRepertoireByProjectId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projectId := chi.URLParam(r, "projectId")
	var repertoire []parsedmodels.Piece

	err := app.service.GetRepertoireByProjectId(projectId, &repertoire)
	if err != nil {
		app.errorLog.Println(err)
	}

	if len(repertoire) == 0 {
		app.ClientError(w, http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(repertoire)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) GetCastByProjectId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projectId := chi.URLParam(r, "projectId")
	var cast []parsedmodels.Contact

	err := app.service.GetCastByProjectId(projectId, &cast)
	if err != nil {
		app.errorLog.Println(err)
	}

	if len(cast) == 0 {
		app.ClientError(w, http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(cast)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) GetScheduleByProjectId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projectId := chi.URLParam(r, "projectId")
	var schedule []parsedmodels.Task

	err := app.service.GetScheduleByProjectId(projectId, &schedule)
	if err != nil {
		app.errorLog.Println(err)
	}

	if len(schedule) == 0 {
		app.ClientError(w, http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(schedule)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) GetLocationById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var location parsedmodels.Location

	locationId := chi.URLParam(r, "locationId")
	err := app.service.GetLocationById(locationId, &location)
	if err != nil {
		app.errorLog.Println(err)
	}

	if unsafe.Sizeof(location) == 0 {
		app.ClientError(w, http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(location)
	if err != nil {
		app.ServerError(w, err)
	}
}
