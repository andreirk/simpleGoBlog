package handlers

import (
	"log"
	"net/http"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/forms"
	"web3/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter,
	r *http.Request) {

	m.App.Session.Put(r.Context(),
		"userid", "derekbanas")

	render.RenderTemplate(w, r, "home.page.tmpl",
		&models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter,
	r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "about.page.tmpl",
		&models.PageData{StrMap: strMap})
}

func (m *Repository) LoginHandler(w http.ResponseWriter,
	r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "login.page.tmpl",
		&models.PageData{StrMap: strMap})
}
func (m *Repository) MakePostHandler(w http.ResponseWriter,
	r *http.Request) {
	render.RenderTemplate(w, r, "make-post.page.tmpl",
		&models.PageData{
			Form: forms.New(nil),
		})
}

// Handler for posting articles using post
func (m *Repository) PostMakePostHandler(w http.ResponseWriter,
	r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	// blog_title := r.Form.Get("blog_title")
	// blog_article := r.Form.Get("blog_article")
	// w.Write([]byte(blog_title))
	// w.Write([]byte(blog_article))

	article := models.Article{
		BlogTitle:   r.Form.Get("blog_title"),
		BlogArticle: r.Form.Get("blog_article"),
	}

	form := forms.New(r.PostForm)
	form.HasValue("blog_title", r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["article"] = article

		render.RenderTemplate(w, r, "make-post.page.tmpl",
			&models.PageData{
				Form: form,
				Data: data,
			})
		return
	}
}

func (m *Repository) PageHandler(w http.ResponseWriter,
	r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "page.page.tmpl",
		&models.PageData{StrMap: strMap})
}
