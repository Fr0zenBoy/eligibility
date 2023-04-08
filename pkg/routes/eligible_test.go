package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Fr0zenBoy/eligibility/pkg/controllers"
)

func setUpRouter() *gin.Engine {
	return gin.Default()
}

func postRequest(request interface{}, path string) *httptest.ResponseRecorder {
	jsonValue, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}

	r := setUpRouter()
	r.POST(path, EligiableHandler)

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestEligibleHandler(t *testing.T){
	t.Run("Test Succesful Request", func(t *testing.T){
		request := controllers.Eligibible{
			NumeroDoDocumento: "14041737706",
			TipoDeConecao: "bifasico",
			ClaseDeConsumo: "comercial",
			ModalidadeTarifa: "convencional",
			HistoricoDeConsumo: []int64{3878, 9760, 5976, 2797, 2481, 5731, 7538, 4392, 7859, 4160, 6941, 4597,},
			}
		output := controllers.EligibleOutput{
			Elegivel: true,
			EconomiaAnaulDeCo2: 5553.24,
		}

		jsonResult, _ := json.Marshal(output)
		post := postRequest(request, "/api/eligible")
		assert.Equal(t, http.StatusAccepted, post.Code)
		assert.Equal(t, jsonResult, post.Body.Bytes())
	})

	t.Run("Test Binding", func(t *testing.T){
		request := controllers.Eligibible{
			NumeroDoDocumento: "14041",
			TipoDeConecao: "bi",
			ClaseDeConsumo: "comercial",
			ModalidadeTarifa: "convencional",
			HistoricoDeConsumo: []int64{3878, 9760, 5976, 2797, 2481, 5731, 7538, 4392, 7859, 4160, 6941, 4597,},
		}
		output := controllers.EligibleOutput{
			Elegivel: true,
			EconomiaAnaulDeCo2: 5553.24,
		}

		jsonResult, _ := json.Marshal(output)
		post := postRequest(request, "/api/eligible")
		assert.Equal(t, http.StatusAccepted, post.Code)
		assert.Equal(t, jsonResult, post.Body.Bytes())
	})

}
