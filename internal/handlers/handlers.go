package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bblank70/bookings/internal/config"
	"github.com/bblank70/bookings/internal/forms"
	"github.com/bblank70/bookings/internal/render"
	"github.com/bblank70/bookings/models"
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

//     // PostPump writes to the pump form
// func (m *Repository) PostPump(w http.ResponseWriter, r *http.Request) {
// 	maintanenceDate := r.Form.Get("MaintanenceDate")
// 	person := r.Form.Get("Person")
// 	equipmentID := r.Form.Get("EquipmentID")
// 	checkValveCleaned := r.Form.Get("CheckValveCleaned")
// 	replacedSeals := r.Form.Get("SealsReplaced")
// 	pumpHeadPiston := r.Form.Get("PumpHeadPiston")
// 	checkValveReplaced := r.Form.Get("CheckValvesReplaced")
// 	timeRequired := r.Form.Get("TimeRequired")
// 	newEfficiency := r.Form.Get("MeasuredEfficiency")
// 	pumpNotes := r.Form.Get("PumpNotes")
// 	printstr := fmt.Sprintf(maintanenceDate, person, equipmentID, checkValveCleaned,
// 		replacedSeals, pumpHeadPiston, checkValveReplaced, timeRequired,
// 		newEfficiency, pumpNotes)
// 	w.Write([]byte(printstr))
// }

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

// Pump renders the Pump-form page
func (m *Repository) Pump(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "pump-form.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostPump handles the posting of pump form
func (m *Repository) PostPump(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	pump := models.Pump{
		EquipmentID:         r.Form.Get("EquipmentID"),
		Person:              r.Form.Get("Person"),
		CheckValveCleaned:   r.Form.Get("CheckValveCleaned"),
		SealsReplaced:       r.Form.Get("SealsReplaced"),
		PumpHeadPiston:      r.Form.Get("PumpHeadPiston"),
		CheckValvesReplaced: r.Form.Get("CheckValvesReplaced"),
		TimeRequired:        r.Form.Get("TimeRequired"),
		MeasuredEfficiency:  r.Form.Get("MeasuredEfficiency"),
		MaintanenceDate:     r.Form.Get("MaintanenceDate"),
	}
	form := forms.New(r.PostForm)

	form.Has("EquipmentID", r)

	// will reload form with data if the data is incomplete.
	if !form.Valid() {
		data := make(map[string]interface{})
		data["pump"] = pump
		render.RenderTemplate(w, r, "pump-form.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}
