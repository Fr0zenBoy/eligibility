package routes

import (
	"net/http"

	"github.com/Fr0zenBoy/authoraizer/controller"
	"github.com/gin-gonic/gin"
)

func AuthoraizerHandler(c *gin.Context) {
	body := controller.Request{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "Invalid inputs. please check your inputs"})
		return
	}
	c.JSON(http.StatusAccepted, body.Allowed(controller.OutPut{}))
}
