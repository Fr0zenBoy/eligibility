package logic

import (
	"fmt"
	"regexp"
)

func RegexpString(input, collate string) bool {
	r, err := regexp.Compile(fmt.Sprintf("(?i)%s", collate))
	if err != nil {
		fmt.Println(err)
	} 
	return r.MatchString(input) 
}
