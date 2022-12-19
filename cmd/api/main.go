package api

import (
	"GPIOapi/internal/gpio"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Structure struct {
	Data      interface{} `json:"data"`
	Error     interface{} `json:"error"`
	Timestamp int         `json:"timestamp"`
}

func Template(c *gin.Context, d gin.H) Structure {
	e := Structure{
		Data:      d,
		Error:     nil,
		Timestamp: int(time.Now().Unix()),
	}

	if c != nil {
		c.JSON(200, e)
	}

	return e
}

func TemplateError(c *gin.Context, err Error) Structure {
	t := Template(nil, nil)
	t.Error = err
	c.JSON(err.Code, t)
	return t
}

func GetGPIORouter() func(c *gin.Context) {
	return func(c *gin.Context) {
		pin := c.Param("pin")
		gp, err := gpio.New(pin)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		pinID, err := gp.Pin()
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		di, err := gp.Direction()
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		v, err := gp.Value()
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		Template(c, gin.H{
			"name":      pin,
			"id":        pinID,
			"exported":  gp.IsExported(),
			"direction": di,
			"value":     v,
		})
	}
}

func Export() func(c *gin.Context) {
	return func(c *gin.Context) {
		pin := c.Param("pin")
		gp, err := gpio.New(pin)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		err = gp.Export()
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		direction := c.DefaultQuery("direction", "out")
		err = gp.SetDirection(direction)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		value := c.DefaultQuery("value", "1")
		err = gp.SetValue(value)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		Template(c, gin.H{
			"name": pin,
		})
	}
}

func Unexport() func(c *gin.Context) {
	return func(c *gin.Context) {
		pin := c.Param("pin")
		gp, err := gpio.New(pin)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		_ = gp.SetValue("1") //Try to set value to 0

		err = gp.Unexport()
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		Template(c, gin.H{
			"name": pin,
		})
	}
}

func Direction() func(c *gin.Context) {
	return func(c *gin.Context) {
		pin := c.Param("pin")
		gp, err := gpio.New(pin)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		d, err := gp.Direction()
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		Template(c, gin.H{
			"name":      pin,
			"direction": d,
		})
	}
}

func DirectionPOST() func(c *gin.Context) {
	return func(c *gin.Context) {
		pin := c.Param("pin")
		gp, err := gpio.New(pin)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		d := c.Query("data")
		if d == "" {
			TemplateError(c, Error{Code: 400, Message: "direction is required"})
			return
		}

		err = gp.SetDirection(d)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		Template(c, gin.H{
			"name":      pin,
			"direction": d,
		})
	}
}

func ValueGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		pin := c.Param("pin")
		gp, err := gpio.New(pin)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		v, err := gp.Value()
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		Template(c, gin.H{
			"name":  pin,
			"value": v,
		})
	}
}

func ValuePOST() func(c *gin.Context) {
	return func(c *gin.Context) {
		pin := c.Param("pin")
		gp, err := gpio.New(pin)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		v := c.Query("data")
		if v == "" {
			TemplateError(c, Error{Code: 400, Message: "value is required"})
			return
		}

		err = gp.SetValue(v)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		Template(c, gin.H{
			"name":  pin,
			"value": v,
		})
	}
}

func SwitchPOST() func(c *gin.Context) {
	return func(c *gin.Context) {
		pin := c.Param("pin")
		gp, err := gpio.New(pin)
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		value, err := gp.Value()
		if err != nil {
			TemplateError(c, Error{Code: 400, Message: err.Error()})
			return
		}

		fmt.Println(value)

		switch value {
		case "1":
			gp.SetValue("0")
			break
		case "0":
			gp.SetValue("1")
			break
		default:
			gp.SetValue("1")
		}
	}
}
