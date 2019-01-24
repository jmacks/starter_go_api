package api

import "net/http"

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(writer, request, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(writer, "view", page)
}
