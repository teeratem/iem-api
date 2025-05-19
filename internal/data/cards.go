package data

import "time"

type Card struct {
	ID         int64     // Unique integer ID for the product
	CreatedAt  time.Time // Timestamp for when the product is added to our database
	Name       string
	SetName    string
	NumberCard int32
	Rarity     string
	Artist     string
	Packs      []string
	Version    int32
}
