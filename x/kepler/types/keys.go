package types

const (
	// ModuleName defines the module name
	ModuleName = "kepler"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_kepler"

	// W3FuncKey defines the key to store the value
	W3FuncKey = "W3Func/value/"

	// W3FuncCountKey defines the key to store the value
	W3FuncCountKey = "W3Func/count/"

	// W3FuncAddressBucketKey defines the key to store the value
	W3FuncAddressBucketKey = "W3Func/address/"
)

var (
	ParamsKey = []byte("p_kepler")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
