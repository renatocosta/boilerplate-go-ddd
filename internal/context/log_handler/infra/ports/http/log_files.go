package http

import (
	"net/http"

	"github.com/ddd/internal/context/log_handler/app"
	"github.com/ddd/internal/context/log_handler/app/query"
	"github.com/ddd/internal/context/log_handler/infra/service"
	"github.com/ddd/pkg/integration"
	"github.com/ddd/pkg/support"
	"github.com/gin-gonic/gin"
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
	resultLogFile := service.SelectLogFileCommandDispatcher(c, &h.App, support.NewString(pathFile))

	resultHumanLogFile := service.CreateHumanLogFileCommandDispatcher(c, &h.App, resultLogFile)

	rawData := integration.PreSendCommand(resultHumanLogFile)

	integration.Dispatch(c, rawData)

	/*if err := h.App.Commands.SelectLogFile.Handle(c, command); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}*/

	c.JSON(http.StatusOK, gin.H{})
}

func (h HttpServer) AvailableLogFiles(c *gin.Context) {

	query := query.AvailableLogFiles{}
	result, err := h.App.Queries.LogFiles.Handle(c, query)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
