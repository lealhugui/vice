package routes

import (
	"net/http"

	"github.com/lealhugui/vice/data"

	"github.com/gin-gonic/gin"
)

type RegisterWorkerRequest struct {
	Host string `json:"host"`
}

func RegisterWorker(c *gin.Context) {
	var req RegisterWorkerRequest
	err := ParsePayload(c, &req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var worker data.Worker
	data.GlobalConn.DB.First(&worker, "host = ?", req.Host)
	if worker.Id != 0 {
		data.GlobalConn.DB.Model(&worker).Update("IsActive", true)
	} else {
		worker.Host = req.Host
		worker.IsActive = true
		data.GlobalConn.DB.Create(&worker)
	}

	c.JSON(http.StatusOK, worker)

}
