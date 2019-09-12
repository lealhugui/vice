package routes

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lealhugui/vice/data"
	"net/http"
)

type registerTaskResult struct{ Id sql.NullInt64 }

type registerTaskRequest struct {
	RunnerHost string `json:"runnerHost"`
	RunnerType string `json:"runnerType"`
	Cmd        string `json:"cmd"`
}

func RegisterTask(c *gin.Context) {
	var req registerTaskRequest
	err := ParsePayload(c, &req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var r registerTaskResult

	rowsAff := data.GlobalConn.Exec("insert into task (id, runner_host, runner_type, cmd) values ((select max(coalesce(id,0)) + 1 from task),?, ?, ?);",
		req.RunnerHost, req.RunnerType, req.Cmd).RowsAffected
	if rowsAff != 1 {
		_ = c.AbortWithError(http.StatusInternalServerError,
			errors.New("registro não inserido"))
	}
	data.GlobalConn.Query("select id from main.task order by id desc limit 1").Scan(&r)
	c.JSON(http.StatusOK, r)
}
