package models

import "database/sql"

type Ingredient struct {
	Name  string         `json:"name" xml:"Name" yaml:"name"`
	Type  string         `json:"type" xml:"Type" yaml:"type"`
	Image sql.NullString `json:"image" xml:"Image" yaml:"image"`
}

type IngredientDose struct {
	Name     Ingredient `json:"name" xml:"Name" yaml:"name"`
	Quantity string     `json:"quantity" xml:"Quantity" yaml:"quantity"`
}

type Cocktail struct {
	Name        string         `json:"name" xml:"Name" yaml:"name"`
	Image       sql.NullString `json:"image" xml:"Image" yaml:"image"`
	Ingredients []Ingredient   `json:"ingredients" xml:"Ingredients" yaml:"ingredients"`
	IsAlcoholic bool           `json:"isAlcoholic" xml:"IsAlcoholic" yaml:"isAlcoholic"`
	Tags        []string       `json:"tags" xml:"Tags" yaml:"tags"`
}

type Glass struct {
	Name           string         `json:"name" xml:"Name" yaml:"name"`
	Image          sql.NullString `json:"image" xml:"Image" yaml:"image"`
	CapacityLiters float32        `json:"capacityLiters" xml:"CapacityLiters" yaml:"capacityLiters"`
}

type Recipe struct {
	Name          Cocktail         `json:"name" xml:"Name" yaml:"name"`
	CanBeServedIn []Glass          `json:"canBeServedIn" xml:"CanBeServedIn" yaml:"canBeServedIn"`
	Doses         []IngredientDose `json:"doses" xml:"Doses" yaml:"doses"`
	Text          string           `json:"text" xml:"Text" yaml:"text"`
	Difficulty    int8             `json:"difficulty" xml:"Difficulty" yaml:"difficulty"`
}
