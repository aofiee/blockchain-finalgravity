package types

import "strings"

// Query endpoints supported by the brewer querier
const (
	// TODO: Describe query parameters, update <action> with your query
	// Query<Action>    = "<action>"
	QueryCreateBrewer = "create-brewer"
)

// QueryResList Queries Result Payload for a query
// type QueryResList []string

// QueryResBrewerList result
type QueryResBrewerList []string

/*
// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}
*/
func (n QueryResBrewerList) String() string {
	return strings.Join(n[:], "\n")
}
