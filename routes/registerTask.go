package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lealhugui/vice/data"
	"net/http"
)

type res struct{ id int64 }

func RegisterTask(c *gin.Context) {
	var r res

	data.GlobalConn.Query("select * from teste").Scan(&r)

	//_ = c.AbortWithError(http.StatusInternalServerError, errors.New("blau"))
	c.JSON(http.StatusOK, r)
}
