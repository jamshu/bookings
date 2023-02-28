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
	return x != ""

}