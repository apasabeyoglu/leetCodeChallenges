package addBinary

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {
	var A, _ = new(big.Int).SetString(a, 2)
	var B, _ = new(big.Int).SetString(b, 2)
	return fmt.Sprintf("%b", new(big.Int).Add(A, B))
}
