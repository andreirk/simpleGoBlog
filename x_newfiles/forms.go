package forms

import (
	"net/http"
	"net/url"
)

// 15. Holds all information assigned to the form
type Form struct {
	url.Values        // Values from the form
	Errors     errors // Errors we track
}

// Create a new Form using form values and return a
// pointer to it
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Check if required fields are empty
func (f *Form) HasValue(tagID string, r *http.Request) bool {
	x := r.Form.Get(tagID)
	if x == "" {
		f.Errors.AddError(tagID, "Field Empty")
		return false
	}
	return true
}

// Returns true or false depending on if the form is valid
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// 15. Now create route for make-reservation
