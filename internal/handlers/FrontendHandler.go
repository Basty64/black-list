package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func FrontendHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := struct {
			Title string
		}{
			Title: "Black list",
		}

		tmpl, err := template.ParseFiles("/Users/basty64/Programming/go/src/black-list/internal/frontend/main/index.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Ошибка выполнения шаблона", http.StatusInternalServerError)
			return
		}
	}
}
