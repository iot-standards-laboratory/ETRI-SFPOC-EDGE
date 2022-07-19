package router

import (
	"bytes"
	"errors"
	"etri-sfpoc-edge/logger"
	"etri-sfpoc-edge/v1/notifier"
	"etrisfpocctnmgmt"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func handleRegisterSvc(sname string, ip net.IP, port string) (string, error) {
	svc, err := db.RegisterService(sname, fmt.Sprintf("%s%s", ip.To4(), port))
	if err != nil {
		return "", err
	}

	fmt.Printf("[registered service]sname: %s / url : %s\n", sname, fmt.Sprintf("%s%s", ip.To4(), port))

	return svc.SID, nil
}

func handleReconnectSvc(sid, sname string, ip net.IP, port string) error {
	originAddr, err := db.GetAddr(sid)
	if err != nil {
		return err
	}

	newAddr := fmt.Sprintf("%s%s", ip.To4(), port)

	fmt.Println("newAddr:", newAddr)
	if strings.Compare(originAddr, newAddr) != 0 {
		_, err := db.UpdateService(sid, newAddr)
		if err != nil {
			return err
		}
	}

	return nil
	// db.GetServices()
}

func PutService(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	sname := c.GetHeader("sname")
	if len(sname) == 0 {
		panic(errors.New("bad request - you should import sname to header"))
	}

	ip, ok := c.RemoteIP()
	if !ok {
		panic(errors.New("scanning IP failed"))
	}

	port := c.GetHeader("port")
	if len(sname) == 0 {
		panic(errors.New("bad request - you should import port to header"))
	}

	path := c.Param("any")

	var sid string
	var err error
	if len(path) == 0 {
		sid, err = handleRegisterSvc(sname, ip, port)
		if err != nil {
			panic(err)
		}
	} else {
		sid = path[1:]
		err := handleReconnectSvc(sid, sname, ip, port)
		if err != nil {
			panic(err)
		}
	}

	go func() {
		time.Sleep(time.Second * 2) // wait for running server
		box.Publish(notifier.NewStatusChangedEvent("service", sname, notifier.SubtokenStatusChanged))
	}()

	c.String(http.StatusOK, sid)

}

func PostService(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	var obj = map[string]string{}

	err := c.BindJSON(&obj)
	if err != nil {
		panic(err)
	}

	err = etrisfpocctnmgmt.CreateContainer(obj["name"])
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusCreated)
}

const lenPrefix = len("/svc/")

func SvcBroker(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	path := c.Param("any")

	var id string

	idx := strings.Index(path[lenPrefix:], "/")
	if idx == -1 {
		id = path[lenPrefix:]
	} else {
		id = path[lenPrefix : lenPrefix+idx]
	}

	ip, err := db.GetAddr(id)
	if err != nil {
		panic(err)
	}

	if strings.Contains(path, "/push/v1") {
		pushHandle(ip, path, c)
	} else {
		apiHandle(ip, path, c)
	}

}

func apiHandle(ip, path string, c *gin.Context) {
	// fmt.Println("path: ", "http://"+ip+path)
	// fmt.Println(c.Request.Method)
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	req, err := http.NewRequest(
		c.Request.Method,
		"http://"+ip+path,
		bytes.NewReader(b),
	)
	if err != nil {
		panic(err)
	}

	req.Header = c.Request.Header

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	io.Copy(c.Writer, resp.Body)
}

func pushHandle(serverAdr, path string, gin *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.Println(r)
		}
	}()
	u := url.URL{Scheme: "ws", Host: serverAdr, Path: path}
	log.Printf("connecting to %s", u.String())

	w, err := upgrader.Upgrade(gin.Writer, gin.Request, nil)
	if err != nil {
		gin.Writer.WriteHeader(http.StatusBadRequest)
		gin.Writer.Write([]byte(err.Error()))
		return
	}
	defer w.Close()

	readDone := make(chan struct{})
	go func() {
		defer func() {
			if r := recover(); r != nil {
				close(readDone)
			}
		}()
		for {
			// Read Messages
			_, _, err := w.ReadMessage()
			if c, k := err.(*websocket.CloseError); k {
				if c.Code == 1000 {
					logger.Println(err)
					panic(err)
				}
			}
		}
	}()

	r, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer r.Close()

	writeDone := make(chan struct{})
	go func() {
		defer close(writeDone)
		var obj map[string]interface{}
		for {
			err := r.ReadJSON(&obj)
			if err != nil {
				return
			}
			w.WriteJSON(obj)
		}
	}()

	select {
	case <-readDone:
		return
	case <-writeDone:
		return
	}

}
