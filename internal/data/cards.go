package data

import "time"

type Card struct {
	ID           int64     // Unique integer ID for the product
	CreatedAt    time.Time // Timestamp for when the product is added to our database
	Name         string
	Brand        string
	Collab       string
	Year         int32
	Connectivity string
	Impedence    string
	NozzleSize   string
	Cable        string
	Genres       []string
	Version      int32
}
