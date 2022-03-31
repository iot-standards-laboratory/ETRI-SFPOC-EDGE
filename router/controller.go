package router

import (
	"etri-sfpoc-edge/notifier"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCtrlList(c *gin.Context) {

	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	list, err := db.GetControllers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	c.JSON(http.StatusOK, list)
}

func PostCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	ctrl, err := db.AddController(c.Request.Body)
	if err != nil {
		panic(err.Error())
	}
	box.Publish(notifier.NewStatusChangedEvent("register controller", "register controller", notifier.SubtokenStatusChanged))
	c.JSON(http.StatusCreated, ctrl)
}
