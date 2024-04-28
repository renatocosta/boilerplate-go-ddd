package http

import (
	"encoding/json"
	"net/http"

	"github.com/ddd/internal/context/log_handler/app"
	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/app/query"
	"github.com/ddd/pkg/support"
	"github.com/google/uuid"
)

type HttpServer struct {
	App app.Application
}

func (h HttpServer) SelectLogFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	name := r.Context().Value("name").(string)

	pathFile := support.GetFilePath("internal/context/log_handler/infra/storage/" + name)

	selectLogFileCommand := command.SelectLogFileCommand{ID: uuid.New(), Path: support.NewString(pathFile)}
	_, err := h.App.Commands.SelectLogFile.Handle(r.Context(), selectLogFileCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h HttpServer) AvailableLogFiles(w http.ResponseWriter, r *http.Request) {

	query := query.AvailableLogFiles{}
	result, err := h.App.Queries.LogFiles.Handle(r.Context(), query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	if result == nil || len(*result) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)
}
