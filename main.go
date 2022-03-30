package main

import (
	"fmt"
	"log"
	"time"
	"upm/udevs_go_auth_service/pkg/security"
)

func main() {
	m := map[string]interface{}{
		"id": "1",
	}
	token, err := security.GenerateJWT(m, time.Duration(time.Hour*2), "secret")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(token)

	res, err := security.ParseClaims(token, "secret")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}
