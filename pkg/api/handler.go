package api

import (
	"log/slog"
	"net/http"

	"github.com/JoshuaWilkes/react-native-example-app/pkg/api/recipesvc"
	"github.com/JoshuaWilkes/react-native-example-app/pkg/gen/proto/recipe/v1alpha1/recipev1alpha1connect"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
)

func (a *API) Handler() http.Handler {
	router := chi.NewRouter()

	logger := httplog.NewLogger("api", httplog.Options{
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
	})

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(httplog.RequestLogger(logger))

	router.Mount(recipev1alpha1connect.NewRecipeServiceHandler(&recipesvc.Service{}))

	return router
}
