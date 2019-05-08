package component

import (
	"net/http"
	"strings"
)

// StaticEndpoint .
type StaticEndpoint struct {
	Core
}

// SetupStatic .
func SetupStatic(core Core) {
	p := new(StaticEndpoint)
	p.Core = core

	p.Router.Get("/v1", p.IndexGET)
	p.Router.Get("/static/*", p.StaticGET)
}

// IndexGET .
// swagger:route GET /v1 healthcheck Ready
//
// API is ready.
//
// Responses:
//   200: OKResponse
func (p StaticEndpoint) IndexGET(w http.ResponseWriter, r *http.Request) (int, error) {
	return p.Response.OK(w, "ready")
}

// StaticGET .
func (p StaticEndpoint) StaticGET(w http.ResponseWriter, r *http.Request) (int, error) {
	if strings.HasSuffix(r.URL.Path, "/") {
		return http.StatusNotFound, nil
	}

	http.ServeFile(w, r, r.URL.Path[1:])
	return http.StatusOK, nil
}