package controllers

import (
	"github.com/mrlsd/echo-cms/libs"
	"net/http"
	"testing"
)

func TestGetMain(t *testing.T) {
	e := libs.HttpClient(t)
	e.GET("/admin/").
		Expect().
		Status(http.StatusOK).
		Body().
		Contains("31")
}

func TestGetLogin(t *testing.T) {
	e := libs.HttpClient(t)
	e.GET("/admin/login/").
		Expect().
		Status(http.StatusOK).
		Body().
		Contains("login to")

	e.GET("/admin/login").
		Expect().
		Status(http.StatusNotFound)
}
