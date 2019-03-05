package forms

import (
	"github.com/jansemmelink/quickform/lib/users"
)

//NewFormData creates empty set of submitted form data
func NewFormData(user users.IUser, form IForm) IFormData {
	return &formData{
		user:   user,
		form:   form,
		fields: make([]IFieldData, 0),
	}
}

//IFormData represents the data submitted in a form
type IFormData interface {
	User() users.IUser
	//Timestamp()
	Form() IForm
	Add(field IFormField, value interface{})
	Fields() []IFieldData
}

//IFieldData is value submitted for one field in a form
type IFieldData interface {
	Field() IFormField
	Value() interface{}
}

type formData struct {
	user   users.IUser
	form   IForm
	fields []IFieldData
}

func (fd formData) User() users.IUser {
	return fd.user
}

func (fd formData) Form() IForm {
	return fd.form
}

func (fd *formData) Add(f IFormField, v interface{}) {
	fd.fields = append(fd.fields, &fieldData{formField: f, value: v})
}

func (fd formData) Fields() []IFieldData {
	return fd.fields
}

type fieldData struct {
	formField IFormField
	value     interface{}
}

func (fd fieldData) Field() IFormField {
	return fd.formField
}

func (fd fieldData) Value() interface{} {
	return fd.value
}
