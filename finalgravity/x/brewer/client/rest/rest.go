package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers blog-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/brewer/create", listBrewerHandler(cliCtx, "brewer")).Methods("GET")
	r.HandleFunc("/brewer/create", createBrewerHandler(cliCtx)).Methods("POST")
}
