package models

import "database/sql"

type Ingredient struct {
	Name  string         `json:"name" xml:"Name" yaml:"name"`
	Type  string         `json:"type" xml:"Type" yaml:"type"`
	Image sql.NullString `json:"image" xml:"Image" yaml:"image"`
}

type IngredientDose struct {
	Name     Ingredient
	Quantity string
}

type Cocktail struct {
	Name        string
	Image       string
	Ingredients []Ingredient
	IsAlcoholic bool
	Tags        []string
}

type Glass struct {
	Name           string
	Image          string
	CapacityLiters float32
}

type Recipe struct {
	Name          Cocktail
	CanBeServedIn []Glass
	Doses         []IngredientDose
	Text          string
	Difficulty    int8
}
