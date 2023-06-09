package models

import "github.com/bblank70/bookings/internal/forms"

// TemplateData holds data to route into the templates from the handlers
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
