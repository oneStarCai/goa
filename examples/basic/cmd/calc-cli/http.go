package main

import (
	"net/http"
	"os"
	"time"

	"goa.design/goa"
	cli "goa.design/goa/examples/basic/gen/http/cli/calc"
	goahttp "goa.design/goa/http"
)

func doHTTP(scheme, host string, timeout int, acceptType string, debug bool) (goa.Endpoint, interface{}, error) {
	var (
		doer goahttp.Doer
	)
	{
		doer = &http.Client{Timeout: time.Duration(timeout) * time.Second}
		if acceptType != "" {
			doer = &acceptTypeDoer{doer, acceptType}
		}

		if debug {
			doer = goahttp.NewDebugDoer(doer)
			doer.(goahttp.DebugDoer).Fprint(os.Stderr)
		}
	}

	return cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
	)
}
func httpUsageCommands() string {
	return cli.UsageCommands()
}

func httpUsageExamples() string {
	return cli.UsageExamples()
}

type acceptTypeDoer struct {
	goahttp.Doer
	acceptType string
}

func (dd *acceptTypeDoer) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", dd.acceptType)
	return dd.Doer.Do(req)
}
