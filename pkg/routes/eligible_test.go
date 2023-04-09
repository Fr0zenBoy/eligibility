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
			TipoDeConexao: "bifasico",
			ClasseDeConsumo: "comercial",
			ModalidadeTarifaria: "convencional",
			HistoricoDeConsumo: []int64{3878, 9760, 5976, 2797, 2481, 5731, 7538, 4392, 7859, 4160, 6941, 4597,},
			}
		output := struct{
			Elegivel bool `json:"elegivel" binding:"required"`
			EconomiaAnaulDeCo2 float64 `json:"economiaAnaulDeCO2"`
		}{
			Elegivel: true,
			EconomiaAnaulDeCo2: 5553.24,
		}

		jsonResult, _ := json.Marshal(output)
		post := postRequest(request, "/api/eligible")
		assert.Equal(t, http.StatusAccepted, post.Code)
		assert.Equal(t, jsonResult, post.Body.Bytes())
	})

	t.Run("Test not eligible", func(t *testing.T){
		request := controllers.Eligibible{
			NumeroDoDocumento: "14041737706",
			TipoDeConexao: "trifasico",
			ClasseDeConsumo: "Poder Publico",
			ModalidadeTarifaria: "azul",
			HistoricoDeConsumo: []int64{550, 686, 750, 732, 433, 399},
		}
		output := struct{
			Elegivel bool `json:"elegivel" binding:"required"`
			RazoesInelegibilidade []string `json:"razoesInelegibilidade"`
		}{
			Elegivel: false,
			RazoesInelegibilidade: []string{
				"Classe de consumo não aceita",
				"Modalidade tarifária não aceita",
				"Consumo muito baixo para tipo de conexão",
			},
		}

		jsonResult, _ := json.Marshal(output)
		post := postRequest(request, "/api/eligible")
		assert.Equal(t, http.StatusAccepted, post.Code)
		assert.ElementsMatch(t, jsonResult, post.Body.Bytes())
	})

	t.Run("Test Bindings errors", func(t *testing.T){
		request := controllers.Eligibible{
			NumeroDoDocumento: "140",
			TipoDeConexao: "bi",
			ClasseDeConsumo: "muitasemuitascoisas",
			ModalidadeTarifaria: "bra nca",
			HistoricoDeConsumo: []int64{
				3878,
				9760,
			},
		} 

		output := struct{
			Errors []ErrorMessage `json:"errors"`
		}{
			Errors: []ErrorMessage{
				{Fild: "NumeroDoDocumento", Message: "This field need require a mininal numeber of values"},
				{Fild: "TipoDeConexao", Message: "This field must end with a specific string"},
				{Fild: "ClasseDeConsumo", Message: "this field needs to have fewer characters"},
				{Fild: "ModalidadeTarifaria", Message: "This field only accepts alpha characters"},
				{Fild: "HistoricoDeConsumo", Message: "This field need require a mininal numeber of values"},
			},
		}

		jsonResult, _ := json.Marshal(output)
		post := postRequest(request, "/api/eligible")
		assert.Equal(t, http.StatusBadRequest, post.Code)
		assert.Equal(t, jsonResult, post.Body.Bytes())
	})

}
