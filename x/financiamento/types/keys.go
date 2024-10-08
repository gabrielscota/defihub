package types

const (
	// ModuleName defines the module name
	ModuleName = "financiamento"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_financiamento"
)

var (
	ParamsKey = []byte("p_financiamento")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
