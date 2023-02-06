package check

import "net/http"

func NewCheckerMap(c *http.Client) map[string]Checker {
	return map[string]Checker{
		"status_code": NewStatusCodeChecker(c),
		"text":        NewTextChecker(c),
	}
}
