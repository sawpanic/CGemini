package risk

import "strings"

// This file contains the logic for correlation controls.

// GetSector returns the sector for a given pair.
// In a real application, this would be more sophisticated.
func GetSector(pair string) string {
	// For now, all pairs are in the "Crypto" sector.
	return "Crypto"
}

// GetEcosystem returns the ecosystem for a given pair.
// In a real application, this would be more sophisticated.
func GetEcosystem(pair string) string {
	// For now, we'll use the base asset as the ecosystem.
	parts := strings.Split(pair, "/")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
