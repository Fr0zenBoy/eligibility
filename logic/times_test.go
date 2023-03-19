package logic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const layout string = "2006-01-02 15:04:05"

func TestTimes(t *testing.T) {

	t.Run("Test parse string in time", func(t *testing.T){

		hourParse  := TimeParse("2019-06-19 21:04:00", layout)
		hour2Parse := TimeParse("2019-06-19 21:05:00", layout)
		assert.WithinDuration(t , hour2Parse, hourParse, time.Duration( 1 * time.Minute))
	})

	t.Run("Test get difference beetween two times in seconds", func(t *testing.T){
		
		hourParse  := TimeParse("2023-03-14 18:04:00", layout)
		hour2Parse := TimeParse("2023-03-14 18:05:00", layout)
		actual := TimeDiff(hour2Parse, hourParse) 
		assert.Equal(t, 60.0, actual)
	})

	t.Run("Test parse string in time and get the time difference", func(t *testing.T) {
		
		firstTIme := "2023-03-14 18:04:00"
		secondidTime := "2023-03-14 14:04:00"
		result := ParseAndGetTimeDiff(firstTIme, secondidTime, layout)
		var expected float64 = 14400

		assert.Equal(t , result, expected)
	})
	
}
