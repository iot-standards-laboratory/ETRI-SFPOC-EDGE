package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if len(c.Param("any")) <= 1 {
		ctrl, err := db.AddControllerWithJsonReader(c.Request.Body)
		if err != nil {
			panic(err.Error())
		}
		c.JSON(http.StatusCreated, ctrl)
	} else {
		cid := c.Param("any")[1:]
		ctrl, err := db.GetController(cid)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, ctrl)
	}
}
