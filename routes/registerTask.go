package routes

import (
	"errors"
	"net/http"

	"github.com/prometheus/common/log"

	"github.com/gin-gonic/gin"
	"github.com/lealhugui/vice/data"
)

func RegisterTask(c *gin.Context) {
	_, err := data.GlobalConn.Query("select * from test")
	if err != nil {
		log.Error(err)
	}
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("blau"))
	//c.JSON(http.StatusOK, rs)
}
