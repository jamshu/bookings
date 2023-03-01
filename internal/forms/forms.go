package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func(f *Form) Has(field string,request *http.Request) bool {
	x := request.Form.Get(field)
	if x == "" {
		f.Errors.Add(field,"This field is required")
		return false
	}
	return true

}

func(f *Form) Valid() bool {
	return len(f.Errors) == 0
}