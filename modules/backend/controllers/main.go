package controllers

import (
	"github.com/labstack/echo"
	"html/template"
	"net/http"

	"fmt"
	"github.com/go-playground/form"
	"github.com/leebenson/conform"
	"github.com/mrlsd/echo-cms/config"
	"net/url"
	"time"
)

func GetMain(c echo.Context) error {
	head := config.GetSiteHeader("backend")
	head.SetTitle("Main+")
	data := struct {
		Name template.HTML
		Head config.SiteHeader
	}{
		Name: "Test <b>BOLD</b>",
		Head: head,
	}
	return c.Render(http.StatusOK, "admin/main", data)
}

func GetForm(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/form", nil)
}

func PostForm(c echo.Context) error {
	type InterfaceStruct struct {
		ID   int
		Name string
	}

	type Company struct {
		NewName    string `formam:"newname" form:"newname" conform:"trim"`
		Public     bool      `formam:"public" form:"public"`
		Website    url.URL   `formam:"website" form:"website"`
		Foundation time.Time `formam:"foundation" form:"foundation"`
		MapExample map[string]string
		Name       string `conform:"trim"`
		Location   struct {
			Country string
			City    string
			Age     uint64
		}
		Products []struct {
			Name string
			Type string
			Age  uint64
		}
		Founders  []string
		Employees int64

		Interface interface{}
	}

	values, _ := c.FormParams()
	c2 := Company{
		Interface: &InterfaceStruct{}, // its is possible to access to the fields although it's an interface field!
	}
	var decoder *form.Decoder
	decoder = form.NewDecoder()
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
        return time.Parse("2006-01-02", vals[0])
    }, time.Time{})
	err := decoder.Decode(&c2, values)
	if err != nil {
		derr := err.(form.DecodeErrors)
		for k, p := range derr {
			println(k)
			println(p.Error())
		}
		fmt.Printf("C2 err: %#v\n\n", derr)
	}
	conform.Strings(&c2)
	fmt.Printf("%v\n", c2)
	println("===================")

	return c.Render(http.StatusOK, "admin/form", nil)
}

func GetLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/login", nil)
}
