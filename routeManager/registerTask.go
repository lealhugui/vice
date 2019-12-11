package routeManager

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lealhugui/vice/data"
)

type RegisterTaskRequest struct {
	RunnerHost string `json:"runnerHost"`
	RunnerType string `json:"runnerType"`
	Cmd        string `json:"cmd"`
}

func RegisterTask(c *gin.Context) {
	var req RegisterTaskRequest
	err := ParsePayload(c, &req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var runners []data.Worker

	data.GlobalConn.DB.Where("host = ?", req.RunnerHost).Find(&runners)
	if len(runners) == 0 {
		_ = c.AbortWithError(http.StatusInternalServerError,
			errors.New("runner unaviable"))
	}

	s := "insert into task (id, runner_host, runner_type, cmd) values ((select max(coalesce(id,0)) + 1 from task),?, ?, ?);"
	rowsAff := data.GlobalConn.Exec(s, req.RunnerHost, req.RunnerType, req.Cmd).RowsAffected
	if rowsAff != 1 {
		_ = c.AbortWithError(http.StatusInternalServerError,
			errors.New("record not inserted"))
	}
	var resultTask data.Task
	//data.GlobalConn.Query("select * from main.task order by id desc limit 1").Scan(&resultTask)
	data.GlobalConn.DB.Last(&resultTask)
	c.JSON(http.StatusOK, resultTask)
}
