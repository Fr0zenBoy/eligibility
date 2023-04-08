package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCo2Savings(t *testing.T){
	consumptions := []int64{3878, 9760, 5976, 2797, 2481, 5731, 7538, 4392, 7859, 4160, 6941, 4597,}
	actual := Co2Savings(consumptions) 
	expected := 5553.24
  assert.Equal(t , expected, actual)
}

func TestAverage(t *testing.T){
	t.Run("Test twelve list size", func(t *testing.T) {
		consumptions := []int64{3878, 9760, 5976, 2797, 2481, 5731, 7538, 4392, 7859, 4160, 6941, 4597,}
		actual := average(consumptions) 
		expected := int64(5509)
		assert.Equal(t , expected, actual)
	})

	t.Run("Test small list", func(t *testing.T) {
		consumptions := []int64{3878, 9760, 5976, 2797,}
		actual := average(consumptions) 
		expected := int64(5602) //5602.75 
		assert.Equal(t , expected, actual)
	})

	t.Run("Test big list", func(t *testing.T) {
		consumptions := []int64{3878, 9760, 5976, 2797, 2481, 5731, 7538, 4392, 7859, 4160, 6941, 4597, 4571, 9436, 9287, 7291}
		actual := average(consumptions) 
		expected := int64(5509)
		assert.Equal(t , expected, actual)
	})

}

func TestCheckConnection(t *testing.T) {

	type test struct {
		kind string
		consumption []int64
	}

	t.Run("Test Connection is valid", func(t *testing.T) {

		tests := []test{
			{"monofasico", []int64{432, 450, 400, 700, 411, 490}},
			{"bifasico",   []int64{523, 500, 570, 598, 501, 530}},
			{"trifasico",  []int64{759, 780, 803, 883, 1000, 500}},
		}

		for _, value := range tests {
			assert.Equal(t, ConectionIsValid(value.kind, value.consumption), true)
		}

	})

	t.Run("Test Connection is not valid", func(t *testing.T){
		
		tests := []test{
			{"monofasico", []int64{401, 398, 300, 278, 291, 230}},
			{"bifasico",   []int64{300, 396, 400, 125, 222, 193}},
			{"trifasico",  []int64{550, 686, 750, 732, 433, 399}},
		}

		for _, value := range tests {
			assert.Equal(t, ConectionIsValid(value.kind, value.consumption), false)
		}
	})
}
