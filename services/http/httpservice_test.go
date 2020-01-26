package http_test

import (
	myHttp "Edwardz43/godevframework/services/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGET(t *testing.T) {
	service := myHttp.GetInstance()

	msg, err := service.GET("http://google.com")

	assert.Nil(t, err)
	assert.NotEmpty(t, msg)
}
