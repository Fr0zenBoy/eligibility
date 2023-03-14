package logic

import "testing"

func TestTime(t *testing.T) {
	t.Run("Time Parse", func(t *testing.T){
		template := "2006-01-02 15:04:05"
		hour := "2019-06-19 21:04:00"

		result := TimeParse(hour, template)
		expected := "fon"

		if result != expected {
			t.Errorf("")
		}
	})
	
}
