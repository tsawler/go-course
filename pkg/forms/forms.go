package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form creates a custom Form struct, which anonymously embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New creates a new form validator
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if field is in post
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

// MinLength method to check that a specific field in the form
func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", d))
	}
}

// Valid method which returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
