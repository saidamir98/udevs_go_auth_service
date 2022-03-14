package util

import (
	"regexp"
)

// IsValidUUID ...
// func IsValidUUID(uuid string) bool {
// 	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
// 	return r.MatchString(uuid)
// }
func IsValidUUID(uuid string) bool {
    r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
    return r.MatchString(uuid)
}

// a1ca1301-4da9-424d-a9e2-578ae6dcde01
// a1ca1301-4da9-424d-a9e2-578ae6dcde03
