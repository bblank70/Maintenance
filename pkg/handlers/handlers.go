package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strmap := make(map[string]string)
	strmap["test"] = "Hello again!"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	strmap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: strmap,
	})
}

// Pump renders the Pump-form page
func (m *Repository) Pump(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "pump-form.page.tmpl", &models.TemplateData{})
}

// PostPump writes to the pump form
func (m *Repository) PostPump(w http.ResponseWriter, r *http.Request) {
	maintanenceDate := r.Form.Get("MaintanenceDate")
	person := r.Form.Get("Person")
	equipmentID := r.Form.Get("EquipmentID")
	checkValveCleaned := r.Form.Get("CheckValveCleaned")
	replacedSeals := r.Form.Get("SealsReplaced")
	pumpHeadPiston := r.Form.Get("PumpHeadPiston")
	checkValveReplaced := r.Form.Get("CheckValvesReplaced")
	timeRequired := r.Form.Get("TimeRequired")
	newEfficiency := r.Form.Get("MeasuredEfficiency")
	pumpNotes := r.Form.Get("PumpNotes")
	printstr := fmt.Sprintf(maintanenceDate, person, equipmentID, checkValveCleaned,
		replacedSeals, pumpHeadPiston, checkValveReplaced, timeRequired,
		newEfficiency, pumpNotes)
	w.Write([]byte(printstr))
}

type jsonResponse struct {
	MaintanenceDate     string  `json:maintanenceDate`
	Person              string  `json:person`
	EquipmentID         string  `json:equipmentID`
	CheckValveCleaned   bool    `json:checkValveCleaned`
	ReplacedSeals       bool    `json:replacedSeals`
	PumpHeadPiston      bool    `json:pumpHeadPiston`
	CheckValvesReplaced bool    `json:checkvalvesreplaced`
	TimeRequired        float64 `json:timeRequired`
	NewEfficiency       float64 `json:newEfficiency`
	PumpNotes           string  `json:pumpNotes`
}

func (m *Repository) PumpJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		MaintanenceDate:     "testdate",
		Person:              "Mark",
		EquipmentID:         "PMPO100",
		CheckValveCleaned:   true,
		ReplacedSeals:       true,
		PumpHeadPiston:      true,
		CheckValvesReplaced: true,
		TimeRequired:        1.7,
		NewEfficiency:       85.3,
		PumpNotes:           "This is a dumb note",
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Stage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "stage.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Repaired(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "repaired.page.tmpl", &models.TemplateData{})
}
