package types

const (
	// ModuleName defines the module name
	ModuleName = "kepler"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_kepler"
)

var (
	ParamsKey = []byte("p_kepler")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
