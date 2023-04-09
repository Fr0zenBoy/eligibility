package routes

import (
	"errors"
	"net/http"

	"github.com/Fr0zenBoy/eligibility/pkg/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Fild string `json:"fild"`
	Message string `json:"message"`
}

func getErrorMessage(f validator.FieldError) string {
	switch f.Tag() {
	case "required":
		return "This field is required"
	case "endswith":
		return "This field must end with a specific string"
	case "numeric":
		return "This field only accepts numeric characters"
	case "min":
		return "This field need require a mininal numeber of values"
	case "max":
		return "This field have a limit of values"
	case "alpha":
		return "This field only accepts alpha characters"
	case "ascii":
		return "This field only accepts ascii characters"
	case "lt":
		return "this field needs to have fewer characters"
	}
	return "Unknown Error"
}

func EligiableHandler(c *gin.Context) {
	body := controllers.Eligibible{}
	if err := c.ShouldBindJSON(&body); err != nil {
		var v validator.ValidationErrors
		if errors.As(err, &v){
			out := make([]ErrorMessage, len(v))
			for i, f := range v {
				out[i] = ErrorMessage{f.Field(), getErrorMessage(f)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out,})
		}
		return
	}
	c.JSON(http.StatusAccepted, body.Resp())
}
