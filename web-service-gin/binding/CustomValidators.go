package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 如果你的时间字符串总是以 2006 - 01 - 02 这样的标准格式出现，将 time_format 改为 "2006 - 01 - 02"，这样就可以正确解析常见的 ISO 8601 格式的时间字符串。

type Booking struct {
	CheckIn  time.Time `form:"check_in" binging:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func CustomValidators(c *gin.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
}
