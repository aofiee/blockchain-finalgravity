package recipes

import (
	"github.com/aofiee/finalgravity/x/recipes/keeper"
	"github.com/aofiee/finalgravity/x/recipes/types"
)

//Const
const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	QuerierRoute      = types.QuerierRoute
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

	NewMsgCreateRecipes                       = types.NewMsgCreateRecipes
	NewMsgCreateWithdrawCoinsFromModuleWallet = types.NewMsgCreateWithdrawCoinsFromModuleWallet
)

//Keeper GenesisState Params MsgCreateRecipes is ...
type (
	Keeper                                 = keeper.Keeper
	GenesisState                           = types.GenesisState
	Params                                 = types.Params
	MsgCreateRecipes                       = types.MsgCreateRecipes
	MsgCreateWithdrawCoinsFromModuleWallet = types.MsgCreateWithdrawCoinsFromModuleWallet
)
