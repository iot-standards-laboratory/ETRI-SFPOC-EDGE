package router

import (
	"etri-sfpoc-edge/model/cache"
	"etri-sfpoc-edge/notifier"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCtrlList(c *gin.Context) {

	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	c.JSON(http.StatusOK, cache.GetCtrls())

}

func PostCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if len(c.Param("any")) <= 1 {
		ctrl, err := db.AddController(c.Request.Body)
		if err != nil {
			panic(err.Error())
		}
		cache.AddCtrls(ctrl)
		box.Publish(notifier.NewStatusChangedEvent("register controller", "register controller", notifier.SubtokenStatusChanged))
		c.JSON(http.StatusCreated, ctrl)
	} else {
		cid := c.Param("any")[1:]
		ctrl, err := db.GetController(cid)
		if err != nil {
			panic(err)
		}
		cache.AddCtrls(ctrl)
		box.Publish(notifier.NewStatusChangedEvent("register controller", "register controller", notifier.SubtokenStatusChanged))
		c.JSON(http.StatusOK, ctrl)
	}
}
