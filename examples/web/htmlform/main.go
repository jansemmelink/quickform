package main

import (
	"fmt"
	"net/http"

	"github.com/jansemmelink/quickform/lib/demo"
	"github.com/jansemmelink/quickform/lib/forms"
	"github.com/jansemmelink/quickform/lib/users"
)

func main() {
	http.ListenAndServe("localhost:8000", router{})
}

type router struct{}

var (
	newFriendForm     forms.IForm
	theOneAndOnlyUser users.IUser
)

func init() {
	newFriendForm = demo.NewFriendForm()
	theOneAndOnlyUser = users.New()
}

func (r router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		res.Write([]byte(htmlForm(theOneAndOnlyUser, newFriendForm)))
	case http.MethodPost:
		req.ParseForm()
		submittedData := forms.NewFormData(theOneAndOnlyUser, newFriendForm)
		for _, formField := range newFriendForm.Fields() {
			value := req.Form.Get(formField.Field().Name())
			submittedData.Add(formField, value)
		}
		//store the submission
		fmt.Printf("With Data:\n")
		for _, sfd := range submittedData.Fields() {
			fmt.Printf("\t%s: %v\n", sfd.Field().Caption(), sfd.Value())
		}

		//write thank you note:
		{
			head := tag("head", "")

			form := fmt.Sprintf("<h1>Thank you</h1>")
			form += fmt.Sprintf("<p>%s submitted with ", newFriendForm.Title())
			for _, sfd := range submittedData.Fields() {
				form += fmt.Sprintf(" %s=%v", sfd.Field().Caption(), sfd.Value())
			}
			form += "</p>"
			body := tag("body", form)

			html := tag("html", head, body)
			res.Write([]byte(html))
		}

	default:
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func htmlForm(u users.IUser, f forms.IForm) string {
	head := tag("head", "")

	form := "<form method=\"post\" action=\"\">"
	form += fmt.Sprintf("<h1>%s</h1>", f.Title())
	form += "<table>"
	for _, ff := range f.Fields() {
		form += fmt.Sprintf("<tr><td>%s</td><td><input type=\"text\" name=\"%s\"/></td></tr>", ff.Caption(), ff.Field().Name())
	}
	form += fmt.Sprintf("<tr><td colspan=\"2\"><input type=\"submit\" value=\"Submit\"></td></tr>")
	form += "</table>"
	form += "</form>"
	body := tag("body", form)

	html := tag("html", head, body)
	return html
}

func tag(name string, contents ...string) string {
	tagged := fmt.Sprintf("<%s>\n", name)
	for _, c := range contents {
		tagged += c + "\n"
	}
	tagged += fmt.Sprintf("</%s>\n", name)
	return tagged
}
