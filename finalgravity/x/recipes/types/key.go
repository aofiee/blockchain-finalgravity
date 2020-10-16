package types

const (
	// ModuleName is the name of the module
	ModuleName = "recipes"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

//RecipesPrefix const
const (
	RecipesPrefix       = "recipes-"
	RecipesWalletPrefix = "recipes-wallet-"
)
