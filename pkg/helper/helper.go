package helper

import (
	"sort"
	"strconv"
	"strings"
)

type queryParams struct {
	Key string
	Val interface{}
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
		arr  []queryParams
	)

	for k, v := range params {
		arr = append(arr, queryParams{
			Key: k,
			Val: v,
		})
	}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i].Key) > len(arr[j].Key)
	})

	for _, v := range arr {
		if v.Key != "" && strings.Contains(namedQuery, ":"+v.Key) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+v.Key, "$"+strconv.Itoa(i))
			args = append(args, v.Val)
			i++
		}
	}

	return namedQuery, args
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}
