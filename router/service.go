package router

import (
	"errors"
	"etri-sfpoc-edge/notifier"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServiceList(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	l, err := db.GetServices()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, l)
}

func GetServiceInfo(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	sname := c.Request.Header.Get("sname")
	if len(sname) == 0 {
		panic(errors.New("wrong request - you should include sname to header"))
	}

	sid, err := db.GetSID(sname)
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, sid)
}

func PutService(c *gin.Context) {
	var obj = map[string]string{}

	err := c.BindJSON(&obj)
	if err != nil {
		panic(err)
	}

	svc, err := db.UpdateService(obj["name"], obj["addr"])
	if err != nil {
		panic(err)
	}

	box.Publish(notifier.NewStatusChangedEvent("service is registered", "service is registered", notifier.SubtokenStatusChanged))
	c.JSON(http.StatusOK, svc)
}
