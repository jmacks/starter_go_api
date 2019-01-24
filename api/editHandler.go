package api

import "net/http"

func editHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(writer, "edit", page)
}
