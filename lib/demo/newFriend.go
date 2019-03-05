package demo

import (
	"github.com/jansemmelink/quickform/lib/forms"
)

//NewFriendForm makes a simple form asking name and surname of a new friend
func NewFriendForm() forms.IForm {
	all := forms.NewFields()
	std := all.Add("std")
	person := std.Add("person")
	personName := person.Add("name")
	personSurname := person.Add("surname")

	return forms.NewForm("New Friend").
		WithField("Name", personName).
		WithField("Surname", personSurname)
	//WithSubmit("Add")
}
