package forms

import (
	"net/http"
	"net/url"
)

// Form struct holds values and errors for our web forms
type Form struct {
	url.Values
	Errors errors
}

// returns true if no errrors in the form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New creates a new form struct which holds the page data
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks to see if field has contents
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field) //get method from errors.go
	if x == "" {
		f.Errors.Add(field, "This field cannot be empty")
		return false
	}
	return true
}
