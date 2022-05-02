package router

import (
	"errors"
	"etri-sfpoc-edge/logger"
	"etri-sfpoc-edge/model"
	"etri-sfpoc-edge/model/cache"
	"etri-sfpoc-edge/notifier"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetDiscoveredDevices(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	c.JSON(http.StatusOK, db.GetDiscoveredDevices())
}

func PostDevice(c *gin.Context) {
	defer handleError(c)

	params := map[string]interface{}{}

	err := c.BindJSON(&params)
	if err != nil {
		panic(err)
	}

	did, ok := params["did"].(string)
	if !ok {
		panic(errors.New("request should have did"))
	}

	dev, err := db.QueryDevice(did)
	if err != nil {
		panic(err)
	}

	err = cache.AddDevs(dev)
	if err != nil {
		panic(err)
	}

	box.Publish(notifier.NewStatusChangedEvent("device connected", nil, notifier.SubtokenStatusChanged))
	c.Status(http.StatusCreated)
}

func PostDiscoveredDevice(c *gin.Context) {
	defer handleError(c)

	var device = &model.Device{}
	err := c.BindJSON(device)
	if err != nil {
		panic(err)
	}

	if !db.IsExistController(device.CID) {
		panic(errors.New("wrong Controller ID"))
	}

	if !db.IsExistDevice(device.DName) {
		panic(errors.New("already exist device"))
	}

	device.DID = uuid.NewString()
	db.AddDiscoveredDevice(device)
	defer db.RemoveDiscoveredDevice(device)
	box.Publish(notifier.NewStatusChangedEvent("discover device", "discover device", notifier.SubtokenStatusChanged))

	resultCh := make(chan notifier.IEvent)
	subscriber := notifier.NewChanSubscriber(
		device.DID,
		device.DID,
		notifier.SubtypeOnce,
		resultCh,
	)
	box.AddSubscriber(subscriber)
	defer box.RemoveSubscriber(subscriber)

	timer := time.NewTimer(60 * time.Second)

	cancelHandler := func(err error) {
		db.RemoveDiscoveredDevice(device)
		box.Publish(notifier.NewStatusChangedEvent(err.Error(), err.Error(), notifier.SubtokenStatusChanged))
		panic(err)
	}
	select {
	case <-resultCh:
		// add device to db
		// sendPOSTtoService(device)
		db.AddDevice(device)
		db.AddService(device.SName)
		db.RemoveDiscoveredDevice(device)
		c.JSON(http.StatusCreated, device)
		box.Publish(notifier.NewStatusChangedEvent("Added device", "Added device", notifier.SubtokenStatusChanged))
		//alarm

	case <-c.Request.Context().Done():
		cancelHandler(errors.New("request is canceled"))
	case <-timer.C:
		cancelHandler(errors.New("timeout"))
	}
}

func PutDiscoveredDevice(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	msg := map[string]string{}
	err := c.BindJSON(&msg)
	if err != nil {
		panic(err)
	}

	if box.Publish(notifier.NewStatusChangedEvent("title: permit", "body: permit", msg["did"])) {
		c.Status(http.StatusOK)
	} else {
		panic(errors.New("wrong token"))
	}
}

func GetDeviceList(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	c.JSON(http.StatusOK, cache.GetDevList())
}

func DeleteDevice(c *gin.Context) {
	defer handleError(c)

	msg := map[string]string{}
	err := c.BindJSON(&msg)
	if err != nil {
		panic(err)
	}

	cid, ok := msg["cid"]
	if !ok {
		panic(errors.New("bad request - you should import cid to request"))
	}

	did, ok := msg["did"]
	if !ok {
		panic(errors.New("bad request - you should import did to request"))
	}

	cache.RemoveDev(&model.Device{
		DID: did,
		CID: cid,
	})

	box.Publish(notifier.NewStatusChangedEvent("remove device", nil, notifier.SubtokenStatusChanged))
	c.Status(http.StatusOK)
}

func DeleteDeviceFromDB(c *gin.Context) {
	defer handleError(c)

	msg := map[string]string{}
	err := c.BindJSON(&msg)
	if err != nil {
		panic(err)
	}

	logger.Println(msg)

	did, ok := msg["did"]
	if !ok {
		panic(errors.New("bad request - you should import did to request"))
	}
	device, err := db.QueryDevice(did)
	if err != nil {
		panic(err)
	}
	db.DeleteDevice(device)
	box.Publish(notifier.NewStatusChangedEvent("remove device", nil, notifier.SubtokenStatusChanged))
	c.Status(http.StatusOK)
}
