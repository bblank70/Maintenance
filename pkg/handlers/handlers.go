package handlers

import (
	"net/http"

	"github.com/bblank70/bookings/models"
	"github.com/bblank70/bookings/pkg/config"
	"github.com/bblank70/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strmap := make(map[string]string)
	strmap["test"] = "Hello again!"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	strmap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: strmap,
	})
}

// Pump renders the Pump-form page
func (m *Repository) Pump(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "pump-form.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Stage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "stage.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Repaired(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "repaired.page.tmpl", &models.TemplateData{})
}
