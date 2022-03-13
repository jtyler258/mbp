package model

import (
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	Name string      `json:"name"     bson:"name"`
	Quantity float64 `json:"quantity" bson:"quantity,omitempty"`
	Units string     `json:"units"    bson:"units,omitempty"`
}

type Step struct {
	Title string          `json:"title"       bson:"title"`
	Text string           `json:"text"        bson:"text,omitempty"`
	Ingredients []*string `json:"ingredients" bson:"ingredients"`
}

type Recipe struct {
	ID string                 `json:"id"          bson:"_id"`	
	Title string              `json:"title"       bson:"title,omitempty"`
	Ingredients []*Ingredient `json:"ingredients" bson:"ingredients,omitempty"`
	Steps []*Step             `json:"steps"       bson:"steps,omitempty"`
	CreatedAt time.Time       `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at"  bson:"updated_at,omitempty"`
}

type CreateIngredientInput struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Units    string  `json:"units"`
}

type CreateRecipeInput struct {
	Title       string                   `json:"title"`
	Ingredients []*CreateIngredientInput `json:"ingredients"`
	Steps       []*CreateStepInput       `json:"steps"`
}

type CreateStepInput struct {
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Ingredients []*string `json:"ingredients"`
}

type UpdateRecipeInput struct {
	CreateRecipeInput
	ID string `json:"id"`
}

func NewRecipe(input *CreateRecipeInput) *Recipe {
	recipe := &Recipe{
		ID: uuid.NewString(),
		Title: input.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	for _, i := range input.Ingredients {
		recipe.Ingredients = append(recipe.Ingredients, (*Ingredient)(i))
	}
	for _, s := range input.Steps {
		recipe.Steps = append(recipe.Steps, (*Step)(s))
	}

	return recipe
}

func NewRecipeFromUpdate(input *UpdateRecipeInput) *Recipe {
	recipe := NewRecipe(&input.CreateRecipeInput)
	recipe.ID = input.ID
	recipe.CreatedAt = time.Time{}

	return recipe
}