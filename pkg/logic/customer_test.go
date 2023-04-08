package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllow(t *testing.T){
	testCase := func(f func(s string) bool, listString []string, expected bool) {

		for _, values := range listString {

			actual := f(values)
			assert.Equal(t, expected, actual)
		}
	}

	t.Run("Test CustumerClasses allowed", func(t *testing.T){
		classes := []string{"Comercial", "Residencial", "Industrial"}
		testCase(CustomerClasses().Allow, classes, true)
	})

	t.Run("Test CustumerClasses not allowed", func(t *testing.T){
		classes := []string{"Poder Publico", "Rural"}
		testCase(CustomerClasses().Allow, classes, false)
	})

	t.Run("Test TarriffModality allowed", func(t *testing.T){
		tariffs := []string{"Convencional", "Branca"}
		testCase(TariffModality().Allow, tariffs, true)
	})

	t.Run("Test TarriffModality not allowed", func(t *testing.T){
		tariffs := []string{"Azul", "Verde",}
		testCase(TariffModality().Allow, tariffs, false)
	})

}
