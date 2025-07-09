package main

import (
	"net/http"
	"runtime"
)

func (app *application) LogError(msg string, err error) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		app.logger.Error(
			msg,
			"file", file,
			"line", line,
			"error", err,
		)
	} else {
		app.logger.Error(msg, "error", err)
	}
}

func (app *application) serverSideErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	_, file, line, ok := runtime.Caller(1)
	msg := "internal server error"

	if ok {
		app.logger.Error(
			msg,
			"file", file,
			"line", line,
			"error", err,
		)
	} else {
		app.logger.Error(msg, "error", err)
	}
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, data any) {
	res := responseData{"error": data}
	err := app.writeJSON(w, status, res, nil)
	if err != nil {
		app.LogError("failed to write JSON response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
