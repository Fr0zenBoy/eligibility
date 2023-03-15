package logic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {

	template := "2006-01-02 15:04:05"

	t.Run("Time Parse", func(t *testing.T){

		hourParse  := TimeParse("2019-06-19 21:04:00", template)
		hour2Parse := TimeParse("2019-06-19 21:05:00", template)
		assert.WithinDuration(t , hour2Parse, hourParse, time.Duration( 1 * time.Minute))
	})

	t.Run("Time diff", func(t *testing.T){
		
		hourParse  := TimeParse("2023-03-14 18:04:00", template)
		hour2Parse := TimeParse("2023-03-14 18:05:00", template)
		actual := TimeDiff(hour2Parse, hourParse) 
		assert.Equal(t, 60.0, actual)
	})
	
}
