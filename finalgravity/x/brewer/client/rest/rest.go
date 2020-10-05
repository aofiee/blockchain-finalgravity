package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	brewerID = "BrewerID"
)

// RegisterRoutes registers blog-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	// this line is used by starport scaffolding
	r.HandleFunc(fmt.Sprintf("/%s/create", storeName), listBrewerHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/id/{%s}", storeName, brewerID), getBrewerByIDHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/module/wallet", storeName), getBrewerWallet(cliCtx, storeName)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/create", storeName), createBrewerHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/module/create/withdraw", storeName), createModuleWithdrawHandler(cliCtx)).Methods("POST")

}
