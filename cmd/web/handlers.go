package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/dhvll/snippetshare/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	//panic("testing")
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, http.StatusOK, "home.tmpl", data)

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, http.StatusOK, "view.tmpl", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "create.tmpl", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	
	err := r.ParseForm()
	if err!= nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}	

	// retreive title and content
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	// convert string to int as we are expecting string

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	} 
	 
	id, err := app.snippets.Insert(title, content, expires)
		if err != nil {
		app.serverError(w,err)
		return
	}

	fieldErrors := make(map[string]string)	

	if strings.TrimSpace(title) == "" {
		fieldErrors["title"] = "This field is blank"
	} else if utf8.RuneCountInString(title) > 100 {
		fieldErrors["title"] = "This field can be longer than 100 characters"
	}

 	if strings.TrimSpace(content) == "" {
		fieldErrors["content"] = "This cannot be  blank"
	}

	if expires != 1 && expires != 7 && expires != 365 {
		fieldErrors["expires"] = "This field must equal 1,7, 365"
	}

	if len(fieldErrors) > 0 {
		fmt.Fprint(w, fieldErrors)
		return
	}

	id, err = app.snippets.Insert(title, content, expires)
	
	if err != nil {
		app.serverError(w, err)
		return
	}


	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

	//w.Write([]byte("Create a new snippet..."))
}
