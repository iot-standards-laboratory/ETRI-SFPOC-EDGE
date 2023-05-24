package router

import (
	"errors"
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/logger"
	"etri-sfpoc-edge/model/consulstorage"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var DB = consulstorage.DefaultDB

func handleError(c *gin.Context) {
	if r := recover(); r != nil {
		logger.Println(r)
		c.String(http.StatusBadRequest, r.(error).Error())
	}
}

func parameterCheck(payload map[string]interface{}, keys []string) error {
	for _, k := range keys {
		_, ok := payload[k]
		if !ok {
			return errors.New("invalid parameter error")
		}
	}

	return nil
}

func getURL(ctx *gin.Context, u string) (string, error) {
	uri, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if uri.Hostname() == "localhost" {
		lastIdx := strings.LastIndex(ctx.Request.Host, ":")
		if lastIdx == -1 {
			lastIdx = len(ctx.Request.Host)
		}
		uri.Host = fmt.Sprintf("%s:%s", ctx.Request.Host[:lastIdx], uri.Port())
	}

	return uri.String(), nil
}

func connectionParams(ctx *gin.Context) map[string]interface{} {
	consulUrl, err := getURL(ctx, config.Params["consulAddr"].(string))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	mqttUrl, err := getURL(ctx, config.Params["mqttAddr"].(string))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return map[string]interface{}{
		"consulAddr": consulUrl,
		"mqttAddr":   mqttUrl,
	}
}
