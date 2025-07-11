package regex

import (
	"regexp"
)

func Contains(regularExpression string, content []string) (int, bool) {
	r := regexp.MustCompile(regularExpression)
	for i, el := range content {
		if r.MatchString(el) {
			return i, true
		}
	}
	return -1, false
}
