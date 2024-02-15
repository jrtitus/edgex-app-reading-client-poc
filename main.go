package main

import (
	"context"
	"fmt"

	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
)

func main() {
	sdk, ok := pkg.NewAppService("service-key")
	if !ok {
		fmt.Println(">>>>>>>>>> no sdk available <<<<<<<<<<")
		return
	}

	lc := sdk.LoggingClient()

	ec := sdk.EventClient()
	if ec == nil {
		lc.Error(">>>>>>>>>> event client is nil <<<<<<<<<<")
	}

	resp, _ := ec.EventCount(context.Background())
	lc.Info(fmt.Sprintf("Event Count: %d", resp.Count))

	rc := sdk.ReadingClient()
	if rc == nil {
		lc.Error(">>>>>>>>>> reading client is nil <<<<<<<<<<")
		return
	}

	resp, _ = rc.ReadingCount(context.Background())
	lc.Info(fmt.Sprintf("Reading Count: %d", resp.Count))

	sdk.Run()
}
