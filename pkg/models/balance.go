// pkg/models/balance.go
package models

import (
	"fmt"
)

type Balance struct {
	Lovelace uint64            // Native ADA in lovelace (1 ADA = 1,000,000 lovelace)
	Assets   map[string]uint64 // Other native assets, keyed by policy ID + asset name
}

func (b *Balance) String() string {
	adaAmount := float64(b.Lovelace) / 1000000.0
	result := fmt.Sprintf("%.6f ADA", adaAmount)

	if len(b.Assets) > 0 {
		result += "\nAssets:"
		for assetID, quantity := range b.Assets {
			result += fmt.Sprintf("\n  %s: %d", assetID, quantity)
		}
	}

	return result
}
