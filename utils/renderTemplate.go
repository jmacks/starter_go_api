package utils


func renderTemplate(writer http.ResponseWriter, tmpl string, page *Page) {
	err := templates.ExecuteTemplate(writer, tmpl+".html", page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}