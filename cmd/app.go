package cmd

import (
	"GPIOapi/cmd/api"
	"github.com/gin-gonic/gin"
)

func Start() error {
	ro := gin.Default()
	ro.Use(CORSMiddleware())
	gin.SetMode(gin.ReleaseMode)

	ro.GET("/gpio/:pin", api.GetGPIORouter())

	//Raw function
	ro.POST("/gpio/:pin/export", api.Export())     //OK
	ro.POST("/gpio/:pin/unexport", api.Unexport()) //OK

	ro.GET("/gpio/:pin/direction", api.Direction())
	ro.POST("/gpio/:pin/direction", api.DirectionPOST())

	ro.POST("/gpio/:pin/value", api.ValuePOST())
	ro.GET("/gpio/:pin/value", api.ValueGET())

	//Custom function
	ro.POST("/gpio/:pin/function/switch", api.SwitchPOST())
	err := ro.Run(":8080")
	if err != nil {
		return err
	}

	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
