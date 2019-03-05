package forms

import "sync"

//NewForm creates a new form
func NewForm(title string) IForm {
	return &form{
		title:  title,
		fields: make([]IFormField, 0),
	}
}

//IForm is a form that user can complete and submit
type IForm interface {
	Title() string
	WithField(caption string, field IField) IForm
	Fields() []IFormField
	//Buttons() []IButton
}

//form implements IForm
type form struct {
	title  string
	mutex  sync.Mutex
	fields []IFormField
}

func (f *form) WithField(caption string, field IField) IForm {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.fields = append(f.fields, formField{caption: caption, field: field})
	return f
}

func (f *form) Title() string {
	return f.title
}

func (f *form) Fields() []IFormField {
	return f.fields
}

//IFormField is a field in a specific form
type IFormField interface {
	Caption() string
	Field() IField
}

type formField struct {
	caption string
	field   IField
}

func (ff formField) Caption() string {
	return ff.caption
}

func (ff formField) Field() IField {
	return ff.field
}
