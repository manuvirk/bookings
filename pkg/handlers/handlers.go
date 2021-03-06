package handlers

import (
	"github.com/manuvirk/bookings/pkg/config"
	"github.com/manuvirk/bookings/pkg/model"
	"github.com/manuvirk/bookings/pkg/render"

	"net/http"
)

var Repo *Repository

//reposiory is the reposiory type
type Repository struct {
	App *config.AppConfig
}

//new repo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//newhandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

//home is the the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &model.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform some logic and send data
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &model.TemplateData{
		StringMap: stringMap,
	})
}
