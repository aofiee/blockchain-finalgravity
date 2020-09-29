package rest

import (
	"net/http"

	"github.com/aofiee/finalgravity/x/brewer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type createBrewerRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Creator      string       `json:"creator" yaml:"creator"`
	BrewerID     string       `json:"BrewerID" yaml:"BrewerID"`
	TypeOfBrewer string       `json:"TypeOfBrewer" yaml:"TypeOfBrewer"`
	Address      string       `json:"Address" yaml:"Address"`
	Telephone    string       `json:"Telephone" yaml:"Telephone"`
	Email        string       `json:"Email" yaml:"Email"`
	Story        string       `json:"Story" yaml:"Story"`
	LogoURL      string       `json:"LogoURL" yaml:"LogoURL"`
	CoverURL     string       `json:"CoverURL" yaml:"CoverURL"`
	Founded      string       `json:"Founded" yaml:"Founded"`
	Founder      string       `json:"Founder" yaml:"Founder"`
	SiteURL      string       `json:"SiteURL" yaml:"SiteURL"`
}

func createBrewerHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBrewerRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgCreateBrewer(creator, req.TypeOfBrewer, req.Address, req.Telephone, req.Email, req.Story, req.LogoURL, req.CoverURL, req.Founded, req.Founder, req.SiteURL)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
