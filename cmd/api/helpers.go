package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	app.infoLog.Println("Client error with status of", status)
	w.WriteHeader(status)

	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf("Error %d %s", status, http.StatusText(status))

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
