package http

import (
	"net/http"

	"github.com/ddd/internal/context/log_handler/app"
	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/app/query"
	"github.com/ddd/pkg/integration"
	"github.com/ddd/pkg/support"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpServer struct {
	App app.Application
}

func (h HttpServer) SelectLogFile(c *gin.Context) {
	selectLogFileRequest := SelectLogFileRequest{}
	if err := c.ShouldBindJSON(&selectLogFileRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pathFile := support.GetFilePath("internal/context/log_handler/infra/storage/" + selectLogFileRequest.Name)

	selectLogFileCommand := command.SelectLogFileCommand{ID: uuid.New(), Path: support.NewString(pathFile)}
	resultLogFile, err := h.App.Commands.SelectLogFile.Handle(c, selectLogFileCommand)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	createHumanLogFileCommand := command.CreateHumanLogFileCommand{ID: uuid.New(), Content: resultLogFile}
	resultHumanLogFile, _ := h.App.Commands.CreateHumanLogFile.Handle(c, createHumanLogFileCommand)

	rawData := integration.PreSendCommand(resultHumanLogFile)

	integration.Dispatch(c, rawData)

	c.JSON(http.StatusCreated, gin.H{})
}

func (h HttpServer) AvailableLogFiles(c *gin.Context) {

	query := query.AvailableLogFiles{}
	result, err := h.App.Queries.LogFiles.Handle(c, query)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if result == nil || len(*result) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"result": []string{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
