package test

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	ht "upm/udevs_go_auth_service/api/http"
	"upm/udevs_go_auth_service/genproto/auth_service"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

var s int64

func TestCreateRole(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 150; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			createRole(t)
		}()
	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createRole(t *testing.T) string {
	response := &ht.Response{}
	request := &auth_service.AddRoleRequest{
		ClientTypeId: "5a3818a9-90f0-44e9-a053-3be0ba1e2c01",
		Name:         faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPost, "/role", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)
	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	} else {
		s++
	}

	return ""
}
