package integ_test_test

import (
	"github.com/gavv/httpexpect"
	_ "github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	e := httpexpect.New(t, "http://localhost:8081")

	e.GET("/v1/podcast/my-file.mp3").
		Expect().
		Status(http.StatusOK)

	e.GET("/v1/podcast/my-file-UNKNOWN.mp3").
		Expect().
		Status(http.StatusNotFound)

	//assert.Equal(t, 1, 1, "1 should equal 1!")
}

func TestPut(t *testing.T) {
	e := httpexpect.New(t, "http://localhost:8081")

	e.PUT("/v1/podcast/my-file.mp3").
		Expect().
		Status(http.StatusOK)
}

func TestDelete(t *testing.T) {
	e := httpexpect.New(t, "http://localhost:8081")

	e.DELETE("/v1/podcast/my-file-UNKNOWN.mp3").
		Expect().
		Status(http.StatusNotFound)

	e.DELETE("/v1/podcast/my-file.mp3").
		Expect().
		Status(http.StatusOK)

	e.DELETE("/v1/podcast/my-file.mp3").
		Expect().
		Status(http.StatusNotFound)
}
