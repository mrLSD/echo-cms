package libs

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
	"time"
)

func HttpClient(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		// prepend this url to all requests
		BaseURL: "http://localhost:3000",

		// use http.Client with a cookie jar and timeout
		Client: &http.Client{
			Jar:     httpexpect.NewJar(),
			Timeout: time.Second * 1,
		},

		// use fatal failures
		Reporter: httpexpect.NewRequireReporter(t),

		// use verbose logging
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}
