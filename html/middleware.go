package html

import (
	"context"
	"net/http"
)

type Middlewares struct {
}

func (m *Middlewares) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// could be struct
		var viewData = make(map[string]interface{})

		// for every request set something, like UserID, UserName etc.
		viewData["global"] = "this is global string for lets say authorized user"

		// get custom navigation menu for authorized user or set default
		viewData["nav"] = []MenuItem{
			{Name: "Home", Path: "/"},
			{Name: "About", Path: "/about"},
			{Name: "Part", Path: "/part"},
		}

		// set view data to request context for child handlers
		ctx := context.WithValue(r.Context(), "view-data", viewData)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

type MenuItem struct {
	Name string
	Path string
}
