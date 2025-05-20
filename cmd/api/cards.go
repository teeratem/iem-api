package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/teeratem/tcgp/internal/data"
)

func (app *application) createCardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new card")
}

func (app *application) showCardHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	card := data.Card{
		ID:         id,
		CreatedAt:  time.Now(),
		Name:       "Bulbasuar",
		SetName:    "Genetic Apex",
		NumberCard: 001,
		Rarity:     "SSS",
		Artist:     "Hideo Kojima",
		Packs:      []string{"Promo A", "Genetic Apex", "Anniversary"},
		Version:    1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"card": card}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
