package brewer

import (
	"github.com/aofiee/finalgravity/x/brewer/keeper"
	"github.com/aofiee/finalgravity/x/brewer/types"
)

//Const
const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

// functions & variable aliases
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

//Keeper GenesisState Params MsgCreateBrewer is ...
type (
	Keeper          = keeper.Keeper
	GenesisState    = types.GenesisState
	Params          = types.Params
	MsgCreateBrewer = types.MsgCreateBrewer
)
