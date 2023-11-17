package forms

// 15. This map has a string key with a slice of strings
type errors map[string][]string

// This function has access to everything in errors
// If there is a form error this will add it to the errors
// map. It is passed the field (id for input tag) where
// error occurs. Message is the error we assign.
func (e errors) AddError(tagID, message string) {
	e[tagID] = append(e[tagID], message)
}

// Get the error message assigned to a field
func (e errors) GetError(tagID string) string {
	es := e[tagID]
	// If the length is = to zero we return nothing
	if len(es) == 0 {
		return ""
	} else {
		// Otherwise we return the error
		return es[0]
	}
}
