package logic

import (
	"fmt"
	"regexp"
)

func RegexMatchMenchantName(merchantInTransaction, merchantInDenyList string) bool {
	match, err := regexp.Compile(fmt.Sprintf("(?i)%s", merchantInDenyList))
	if err != nil {
		fmt.Println(err)
	}
	return match.MatchString(merchantInTransaction)
}
