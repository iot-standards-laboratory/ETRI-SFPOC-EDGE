package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type contextKey int

var ginContextKey contextKey

func newRouter() *gin.Engine {
	rt := gin.New()

	store := cookie.NewStore([]byte("secret_string"))
	// rt.Use(sessions.Sessions("session_name", store))
	rt.Use(sessions.SessionsMany([]string{"session", "access_key"}, store))
	rt.GET("/", func(c *gin.Context) {
		// s := sessions.Default(c)
		fmt.Println("^^")
		s := sessions.DefaultMany(c, "access_key")

		c.Writer.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		c.Writer.Header().Set("Pragma", "no-cache")
		// c.Writer.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
		// c.Writer.Header().Set("X-Accel-Expires", "0")

		s.Set("rtp", "edge")
		s.Set("name", "godopu")
		s.Save()

		s = sessions.DefaultMany(c, "session")
		s.Set("session", "session")
		s.Save()
		c.String(http.StatusOK, "Hello world")
		// c.Redirect(http.StatusPermanentRedirect, "/home")
		// if user == nil {
		// 	c.String(http.StatusOK, "wrong")
		// } else {
		// 	c.String(http.StatusOK, user.(string))
		// }
	})

	rt.GET("/svc", func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		c.Writer.Header().Set("Pragma", "no-cache")

		s := sessions.Default(c)

		s.Set("rtp", "svc")
		s.Save()
		c.Redirect(http.StatusPermanentRedirect, "/home")
	})

	rt.GET("/home", func(c *gin.Context) {
		s := sessions.Default(c)
		rtp := s.Get("rtp")
		if strings.Compare(rtp.(string), "edge") == 0 {
			// c.HTML(http.StatusOK, "chat.html", nil)
			c.String(http.StatusOK, "home")
		} else {
			// c.HTML(http.StatusOK, "chat2.html", nil)
			c.String(http.StatusOK, "home-2")
		}
	})

	// rt.Use(GinContextToContextMiddleware())
	rt.GET("/chat", func(c *gin.Context) {
		s := sessions.Default(c)
		user := s.Get("user")
		if user == nil {
			c.String(http.StatusOK, "wrong")
		} else {
			c.String(http.StatusOK, user.(string))
		}
	})

	rt.GET("/login", func(c *gin.Context) {
		b_body, _ := ioutil.ReadAll(c.Request.Body)

		s := sessions.Default(c)
		s.Set("user", string(b_body))
		s.Save()
		c.Status(http.StatusOK)
	})

	return rt
}

func InjectRoutePath(c *gin.Context) {
	s := sessions.Default(c)
	s.Set("rtp", "edge")
	c.Next()
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(ginContextKey)
		ctx := context.WithValue(c.Request.Context(), ginContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(ginContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}
	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func main() {
	newRouter().Run(":9988")
}
