package regex

import "regexp"

func Contains(regularExpression string, content []string) bool {
	r := regexp.MustCompile(regularExpression)
	for _, el := range content {
		if r.MatchString(el) {
			return true
		}
	}
	return false
}
