package main

import (
	"errors"
	"fmt"
	"github.com/halviet/snippetbox/internal/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	//files := []string{
	//	"./ui/html/base.tmpl",
	//	"./ui/html/partials/nav.tmpl",
	//	"./ui/html/pages/home.tmpl",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.l.Error(
	//		err.Error(),
	//		slog.String("method", r.Method),
	//		slog.String("path", r.URL.RequestURI()),
	//	)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}
	//
	//err = ts.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	app.serverError(w, r, err)
	//}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the snippet creation page"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\\nClimb Mount Fuji,\\nBut slowly, slowly!\\n\\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
