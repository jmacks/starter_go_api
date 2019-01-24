package api

import "net/http"

func saveHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/save/"):]
	body := request.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, request, "/view/"+title, http.StatusFound)
}
