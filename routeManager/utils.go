package routeManager

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func ParsePayload(c *gin.Context, bind interface{}) error {
	bytesPayload, _ := c.GetRawData()
	return json.Unmarshal(bytesPayload, &bind)
}
