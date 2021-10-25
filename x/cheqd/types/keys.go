package types

const (
	// ModuleName defines the module name
	ModuleName = "cheqd"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	DidMethod = ModuleName

	DidPrefix = "did"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DidKey      = "did:"
	DidCountKey = "did-count:"
)

const (
	SchemaKey      = "schema:"
	SchemaCountKey = "schema-count:"
)

const (
	CredDefKey      = "credDef:"
	CredDefCountKey = "credDef-count:"
)
