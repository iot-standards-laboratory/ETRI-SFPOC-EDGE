package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func connectionParams() map[string]interface{} {
	return map[string]interface{}{
		"wsAddr":     "ws://dsad:8000/connection/websocket",
		"consulAddr": "http://dsad:9999",
	}
}

func PostCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	fmt.Println(c.Param("any"))
	if len(c.Param("any")) <= 1 {
		ctrl, err := db.AddControllerWithJsonReader(c.Request.Body)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(http.StatusCreated, ctrl)
	} else {
		cid := c.Param("any")[1:]
		_, err := db.GetController(cid)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, connectionParams())
	}
}
