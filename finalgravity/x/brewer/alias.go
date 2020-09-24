package brewer

import (
	"github.com/aofiee/finalgravity/x/brewer/keeper"
	"github.com/aofiee/finalgravity/x/brewer/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc

	NewMsgCreateBrewer = types.NewMsgCreateBrewer
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params
	MsgCreateBrewer = types.MsgCreateBrewer
)
