package beer

import (
	"github.com/aofiee/finalgravity/x/beer/keeper"
	"github.com/aofiee/finalgravity/x/beer/types"
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

	NewMsgCreateBeer                          = types.NewMsgCreateBeer
	NewMsgCreateWithdrawCoinsFromModuleWallet = types.NewMsgCreateWithdrawCoinsFromModuleWallet
)

//Keeper GenesisState Params MsgCreateBeer is ...
type (
	Keeper                                 = keeper.Keeper
	GenesisState                           = types.GenesisState
	Params                                 = types.Params
	MsgCreateBeer                          = types.MsgCreateBeer
	MsgCreateWithdrawCoinsFromModuleWallet = types.MsgCreateWithdrawCoinsFromModuleWallet
)
