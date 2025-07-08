package main

import (
	"runtime"
)

func (app *application) LogError(msg string) {
	// Get caller info
	_, file, line, ok := runtime.Caller(1)
	if ok {
		app.logger.Error("[ERROR] %s (called from %s:%d)", msg, file, line)
	} else {
		app.logger.Error("[ERROR] %s (caller unknown)", msg)
	}
}
