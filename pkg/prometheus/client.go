package prometheus

import (
	"os"

	clientApi "github.com/prometheus/client_golang/api"
	APIV1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

var api APIV1.API
var client clientApi.Client

func init() {
	addr := os.Getenv("PROMETHEUS_ADDR")
	if addr == "" {
		panic("\"PROMETHEUS_ADDR\" should be set in env")
	}

	var err error
	client, err = clientApi.NewClient(clientApi.Config{
		Address:      addr,
		RoundTripper: clientApi.DefaultRoundTripper,
	})
	if err != nil {
		panic(err)
	}

	api = APIV1.NewAPI(client)
	return
}
