package db

import (
	"context"
	"errors"

	"github.com/jtyler258/mbp/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeRepository struct {
	database *Database
	collection *mongo.Collection
}

func NewRecipeRepository(database *Database) (*RecipeRepository) {
	return &RecipeRepository{
		collection: database.client.Database("mbp").Collection("recipes"),
		database: database,
	}
}

func (r *RecipeRepository) Insert(ctx *context.Context, doc *model.Recipe) (*model.Recipe, error) {
	_, err := r.collection.InsertOne(*ctx, doc)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (r *RecipeRepository) Update(ctx *context.Context, doc *model.Recipe) (*model.Recipe, error) {
	_, err := r.collection.UpdateByID(*ctx, doc.ID, bson.D{{"$set", doc}})
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, doc.ID)
}

func (r *RecipeRepository) Get(ctx *context.Context, id string) (*model.Recipe, error) {
	var recipe *model.Recipe
	r.collection.FindOne(*ctx, bson.D{{"_id", id}}).Decode(&recipe)

	return recipe, nil
}

func (r *RecipeRepository) List(ctx *context.Context) ([]*model.Recipe, error) {
	cursor, err := r.collection.Find(*ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(*ctx)
	var results []*model.Recipe

	for cursor.Next(*ctx) {
		var result *model.Recipe
		err = cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *RecipeRepository) Delete(ctx *context.Context, id string) (*model.Recipe, error) {
	recipe, err := r.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	res, err := r.collection.DeleteOne(*ctx, bson.D{{"_id", id}})
	if err != nil || res.DeletedCount < 1 {
		return nil, errors.New("Error deleting document")
	}

	return recipe, nil
}