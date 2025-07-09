package main

import (
	"net/http"
)

func (app *application) GetToolByDescriptionHandler(w http.ResponseWriter, r *http.Request) {

	description := r.URL.Query().Get("description")
	ctx := r.Context()

	tools, opErr := app.toolService.GetToolByDescription(ctx, description)
	if opErr != nil {
		app.serverSideErrorResponse(w, r, opErr)
		return
	}

	err := app.writeJSON(w, http.StatusFound, tools, nil)
	if err != nil {
		app.serverSideErrorResponse(w, r, err)
		return
	}
}
