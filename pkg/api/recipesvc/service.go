package recipesvc

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	recipev1alpha1 "github.com/JoshuaWilkes/react-native-example-app/pkg/gen/proto/recipe/v1alpha1"
	"github.com/JoshuaWilkes/react-native-example-app/pkg/gen/proto/recipe/v1alpha1/recipev1alpha1connect"
)

type Service struct {
}

var _ recipev1alpha1connect.RecipeServiceHandler = &Service{}

func (s *Service) QueryRecipes(ctx context.Context, req *connect.Request[recipev1alpha1.QueryRecipesRequest]) (*connect.Response[recipev1alpha1.QueryRecipesResponse], error) {
	// Extract filters from the request
	filter := strings.ToLower(req.Msg.Search)

	// Filter recipes based on title and tags
	filteredRecipes := []*recipev1alpha1.Recipe{}

	if filter == "" {
		res := &recipev1alpha1.QueryRecipesResponse{
			Recipes: recipes,
		}
		return connect.NewResponse(res), nil
	}

	for _, recipe := range recipes {
		if strings.Contains(strings.ToLower(recipe.Title), strings.ToLower(filter)) {
			filteredRecipes = append(filteredRecipes, recipe)
			continue
		}

		for _, recipeTag := range recipe.Tags {
			if strings.Contains(strings.ToLower(recipeTag), strings.ToLower(filter)) {
				filteredRecipes = append(filteredRecipes, recipe)
				continue
			}
		}

		for _, ingredient := range recipe.Ingredients {
			if strings.Contains(strings.ToLower(ingredient), strings.ToLower(filter)) {
				filteredRecipes = append(filteredRecipes, recipe)
				continue
			}
		}

	}

	res := &recipev1alpha1.QueryRecipesResponse{
		Recipes: filteredRecipes,
	}

	return connect.NewResponse(res), nil
}

func (s *Service) TopRecipes(context.Context, *connect.Request[recipev1alpha1.TopRecipesRequest]) (*connect.Response[recipev1alpha1.TopRecipesResponse], error) {

	res := &recipev1alpha1.TopRecipesResponse{
		// for this mock API just return the first 2 recipes
		Recipes: recipes[:2],
	}

	return connect.NewResponse(res), nil
}

var recipes = []*recipev1alpha1.Recipe{
	{
		Id:          1,
		Title:       "Simple Pasta",
		Ingredients: []string{"Pasta", "Olive Oil", "Garlic", "Salt", "Pepper"},
		Body:        "Cook the pasta in boiling water. Drain and toss with olive oil, minced garlic, salt, and pepper. Serve warm.",
		Tags:        []string{"vegetarian", "quick", "italian"},
	},
	{
		Id:          2,
		Title:       "Basic Omelette",
		Ingredients: []string{"Eggs", "Salt", "Pepper", "Butter"},
		Body:        "Beat the eggs with salt and pepper. Heat butter in a pan, pour in the eggs, and cook until set. Fold and serve.",
		Tags:        []string{"breakfast", "quick", "protein"},
	},
	{
		Id:          3,
		Title:       "Grilled Cheese Sandwich",
		Ingredients: []string{"Bread", "Cheese", "Butter"},
		Body:        "Butter the bread slices, place cheese between them, and grill in a pan until golden brown and cheese is melted.",
		Tags:        []string{"snack", "quick", "comfort food"},
	},
	{
		Id:          4,
		Title:       "Fruit Salad",
		Ingredients: []string{"Apple", "Banana", "Orange", "Honey"},
		Body:        "Chop all fruits, mix them in a bowl, and drizzle with honey. Serve chilled.",
		Tags:        []string{"healthy", "dessert", "vegan"},
	},
	{
		Id:          5,
		Title:       "Peanut Butter Toast",
		Ingredients: []string{"Bread", "Peanut Butter", "Banana"},
		Body:        "Spread peanut butter on toast and top with banana slices. Enjoy.",
		Tags:        []string{"breakfast", "quick", "snack"},
	},
}
