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
		return "This field need finishing with 'fasico'"
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Errors": out,})
		}
		return
	}
	c.JSON(http.StatusAccepted, body.Resp(controllers.EligibleOutput{}))
}
