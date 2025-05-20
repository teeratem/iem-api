package data

import "time"

type Card struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	SetName    string    `json:"setname"`
	NumberCard int32     `json:"numbercard"`
	Rarity     string    `json:"rarity"`
	Artist     string    `json:"artist"`
	Packs      []string  `json:"packs,omitzero"`
	Version    int32     `json:"version"`
}
