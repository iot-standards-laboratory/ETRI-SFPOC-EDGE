package router

import (
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/controller/state"
	"etri-sfpoc-edge/docs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewInitRouter() *gin.Engine {
	docs.SwaggerInfo.Title = "ETRI Smartfarm PoC API"
	docs.SwaggerInfo.Description = "Edge server for ETRI Smartfarm PoC API"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/"

	assetEngine := gin.New()
	assetEngine.Static("/", "./www")

	swaggerEngine := gin.New()
	swaggerEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r := gin.New()
	r.Any("/*any", func(c *gin.Context) {
		defer handleError(c)
		w := c.Writer
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		path := c.Param("any")

		if path == "/loading" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"page": "/init",
			})
		} else if strings.HasPrefix(path, "/swagger") {
			swaggerEngine.HandleContext(c)
		} else if path == "/init" && c.Request.Method == "POST" {
			POST_Init(c)
		} else if path == "/mqtt" {
			// mqtt
			mqttReverseProxy.ServeHTTP(c.Writer, c.Request)
		} else if len(path) <= 1 || strings.HasSuffix(c.Request.Header.Get("Referer"), "/") {
			assetEngine.HandleContext(c)
		} else {
			c.Status(http.StatusBadRequest)
		}
	})

	return r
}

// {object} welcomeModel

// Init godoc
// @Summary Summary를 적어 줍니다.
// @Description 자세한 설명은 이곳에 적습니다.
// @Param test query int false "test parameter"
// @Param body body object{mqttAddr=string,consulAddr=string} true "User ID and comma separated roles"
// @Accept  json
// @Success 200
// @Router /init [post]
func POST_Init(c *gin.Context) {
	payload := map[string]interface{}{}
	err := c.BindJSON(&payload)
	if err != nil {
		panic(err)
	}

	err = parameterCheck(payload, []string{"consulAddr", "mqttAddr"})
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)

	config.Params["consulAddr"] = payload["consulAddr"]
	config.Params["mqttAddr"] = payload["mqttAddr"]
	state.Put(state.STATE_INITIALIZED)
}

// Init godoc
// @Summary get params to load page.
// @Description get params to load page.
// @success 200 {object} object{page=string}
// @Router /loading [get]
func GET_Loading(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"page": "/init",
	})
}
