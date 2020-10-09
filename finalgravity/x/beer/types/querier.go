package types

//const
const (
	QueryListBeer      = "list-beer"
	QueryGetBeerWallet = "get-wallet"
)

/*
Below you will be able how to set your own queries:


// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}

*/
