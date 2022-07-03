package models

type Ingredient struct {
	Name  string
	Type  string
	Image string
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
