package forms

type errors map[string][]string

// Add appends to errors for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get gets the first error message in errors
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
