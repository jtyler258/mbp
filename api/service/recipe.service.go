package service

import (
	"context"

	"github.com/jtyler258/mbp/api/db"
	"github.com/jtyler258/mbp/api/model"
)

type RecipeService struct {
	repository *db.RecipeRepository
}

func NewRecipeService(repository *db.RecipeRepository) *RecipeService {
	return &RecipeService{
		repository: repository,
	}
}

func (r *RecipeService) Create(ctx context.Context, recipe *model.CreateRecipeInput) (*model.Recipe, error) {
	return r.repository.Insert(ctx, model.NewRecipe(recipe))
}

func (r *RecipeService) List(ctx context.Context) ([]*model.Recipe, error) {
	return r.repository.List(ctx)
}

func (r *RecipeService) Get(ctx context.Context, id string) (*model.Recipe, error) {
	return r.repository.Get(ctx, id)
}

func (r *RecipeService) Update(ctx context.Context, recipe *model.UpdateRecipeInput) (*model.Recipe, error) {
	return r.repository.Update(ctx, model.NewRecipeFromUpdate(recipe))
}

func (r *RecipeService) Delete(ctx context.Context, id string) (*model.Recipe, error) {
	return r.repository.Delete(ctx, id)
}