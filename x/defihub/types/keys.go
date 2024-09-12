package types

const (
	// ModuleName defines the module name
	ModuleName = "defihub"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_defihub"
)

var (
	ParamsKey = []byte("p_defihub")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
