package middlewares

import (
	"encoding/json"
	"net/http"
	common "piteroni/online-store-api"

	"github.com/rs/cors"
)

func SetupCORS(handler http.Handler) (http.Handler, error) {
	o, err := common.Env("ALLOW_ORIGINS")
	if err != nil {
		return nil, err
	}

	origins := []string{}
	err = json.Unmarshal([]byte(o), &origins)
	if err != nil {
		return nil, err
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})

	return c.Handler(handler), nil
}
