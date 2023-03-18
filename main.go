package main

// 1. The transaction amount should not be above limit
// 2. No transaction should be approved when the card is blocked
// 3. The first transaction shouldn't be above 90% of the limit
// 4. There should not be more than 10 transactions on the same merchant
// 5. Merchant blacklist
// 6. There should not be more than 3 transactions on a 2 minutes interval

import (
	"fmt"
	"net/http"

	"github.com/Fr0zenBoy/authoraizer/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.POST("/api/authoraizer", func(ctx *gin.Context) {
		body := &controller.Request{}
		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		ctx.JSON(http.StatusAccepted, body.Allowed(controller.OutPut{}))
	})
	router.Run("localhost:8080")
}
