package pressurise

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PressuriseApp interface {
	RegisterPages(handlers map[string]func() http.HandlerFunc)
	Route(path string, handler http.HandlerFunc)
	Run(address string) error
}

type pressuriseApp struct {
	Router *chi.Mux
}

// registers pages generated by pressurise cli,
// handlers argument is the generated map
func (a *pressuriseApp) RegisterPages(handlers map[string]func() http.HandlerFunc) {
	for k, v := range handlers {
		a.Router.Handle(k, v())
	}
}

// alias for app.Router.Handle
func (a *pressuriseApp) Route(pattern string, handler http.HandlerFunc) {
	a.Router.Handle(pattern, handler)
}

// runs the app on the selected address
func (a *pressuriseApp) Run(address string) error {
	return http.ListenAndServe(address, a.Router)
}

// creates a new pressurise app
func NewPressurise() PressuriseApp {
	return &pressuriseApp{
		Router: chi.NewRouter(),
	}
}
