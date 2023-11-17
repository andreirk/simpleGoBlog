package models

import "web3/pkg/forms"

type PageData struct {
	StrMap    map[string]string
	IntMap    map[string]int
	FltMap    map[string]float32
	DataMap   map[string]interface{}
	CSRFToken string
	Warning   string
	Error     string

	// 15. Create a pointer to our form so form data
	// will be available everywhere
	Form *forms.Form

	// 15. Create a place to store any type of data
	Data map[string]interface{}
}
