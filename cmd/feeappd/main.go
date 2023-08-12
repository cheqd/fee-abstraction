package main

import (
	"os"

	app "github.com/notional-labs/fee-abstraction/v3/app"
	"github.com/notional-labs/fee-abstraction/v3/app/params"
	"github.com/notional-labs/fee-abstraction/v3/cmd/feeappd/cmd"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
