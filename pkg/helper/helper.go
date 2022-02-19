package helper

import (
	"strconv"
	"strings"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" {
			oldsize := len(namedQuery)
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))

			if oldsize != len(namedQuery) {
				args = append(args, v)
				i++
			}
		}
	}

	return namedQuery, args
}
