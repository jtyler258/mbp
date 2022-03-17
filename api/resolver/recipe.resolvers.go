package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jtyler258/mbp/api/model"
)

func (r *mutationResolver) CreateRecipe(ctx context.Context, input model.CreateRecipeInput) (*model.Recipe, error) {
	return r.RecipeService.Create(&ctx, &input)
}

func (r *mutationResolver) UpdateRecipe(ctx context.Context, input model.UpdateRecipeInput) (*model.Recipe, error) {
	return r.RecipeService.Update(&ctx, &input)
}

func (r *mutationResolver) DeleteRecipe(ctx context.Context, id string) (*model.Recipe, error) {
	return r.RecipeService.Delete(&ctx, id)
}

func (r *queryResolver) Recipe(ctx context.Context, id string) (*model.Recipe, error) {
	return r.RecipeService.Get(&ctx, id)
}

func (r *queryResolver) Recipes(ctx context.Context) ([]*model.Recipe, error) {
	return r.RecipeService.List(&ctx)
}
