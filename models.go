package main

type MealSection int

const (
	Protein MealSection = iota
	Starch
	Vegetable
)

type MealPart struct {
	Name   string `json: "name"`
	Recipe string `json: "recipe"`
}

type Proteins struct {
	Proteins []MealPart `json: "proteins"`
}

type Meal struct {
	protein   MealPart
	side      MealPart
	vegetable MealPart
}
