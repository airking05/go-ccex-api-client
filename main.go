package main

import (
	"fmt"
	"github.com/fxpgr/go-ccex-api-client/api/public"
)

func main() {
	cli, err := public.NewClient("hitbtc")
	if err != nil {
		panic(err)
	}
	fmt.Println(cli.RateMap())
}
