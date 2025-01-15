package recipesvc

import (
	"context"

	"connectrpc.com/connect"
	recipev1alpha1 "github.com/JoshuaWilkes/react-native-example-app/pkg/gen/proto/recipe/v1alpha1"
	"github.com/JoshuaWilkes/react-native-example-app/pkg/gen/proto/recipe/v1alpha1/recipev1alpha1connect"
)

type Service struct {
}

var _ recipev1alpha1connect.RecipeServiceHandler = &Service{}

func (s *Service) QueryRecipes(context.Context, *connect.Request[recipev1alpha1.QueryRecipesRequest]) (*connect.Response[recipev1alpha1.QueryRecipesResponse], error) {

	res := &recipev1alpha1.QueryRecipesResponse{}

	return connect.NewResponse(res), nil
}

func (s *Service) TopRecipes(context.Context, *connect.Request[recipev1alpha1.TopRecipesRequest]) (*connect.Response[recipev1alpha1.TopRecipesResponse], error) {

	res := &recipev1alpha1.TopRecipesResponse{}

	return connect.NewResponse(res), nil
}
