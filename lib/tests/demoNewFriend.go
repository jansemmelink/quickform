package main

import (
	"fmt"

	"github.com/jansemmelink/quickform/lib/demo"
	"github.com/jansemmelink/quickform/lib/forms"
	"github.com/jansemmelink/quickform/lib/users"
)

//TestDemoForm ...
func main() {
	user := users.New()
	submittedData := prompt(user, demo.NewFriendForm())

	fmt.Printf("User: %s\n", submittedData.User())
	fmt.Printf("Submitted form: %s\n", submittedData.Form().Title())
	fmt.Printf("With Data:\n")
	for _, sfd := range submittedData.Fields() {
		fmt.Printf("\t%s: %v\n", sfd.Field().Caption(), sfd.Value())
	}

}

func prompt(u users.IUser, f forms.IForm) forms.IFormData {
	fd := forms.NewFormData(u, f)

	fmt.Printf("\n")
	fmt.Printf("=======================================================\n")
	fmt.Printf("FORM: %s\n", f.Title())
	for fieldNr, field := range f.Fields() {
		fmt.Printf("%d) %s ?", fieldNr, field.Caption())
		var value string
		fmt.Scanf("%s", &value)
		fd.Add(field, value)
	}
	fmt.Printf("=======================================================\n")
	return fd
}
