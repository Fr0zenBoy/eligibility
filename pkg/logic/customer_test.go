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
		classes := []string{"comercial", "residencial", "industrial"}
		testCase(CustomerClasses().Allow, classes, true)
	})

	t.Run("Test CustumerClasses not allowed", func(t *testing.T){
		classes := []string{"poder publico", "rural"}
		testCase(CustomerClasses().Allow, classes, false)
	})

	t.Run("Test TarriffModality allowed", func(t *testing.T){
		tariffs := []string{"convencional", "branca"}
		testCase(TariffModality().Allow, tariffs, true)
	})

	t.Run("Test TarriffModality not allowed", func(t *testing.T){
		tariffs := []string{"azul", "verde",}
		testCase(TariffModality().Allow, tariffs, false)
	})

}
