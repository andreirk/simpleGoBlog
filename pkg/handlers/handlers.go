package handlers

import (
	"fmt"
	"log"
	"net/http"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/dbdriver"
	"web3/pkg/forms"
	"web3/pkg/render"
	"web3/pkg/repository"
	"web3/pkg/repository/dbrepo"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(ac *config.AppConfig, db *dbdriver.DB) *Repository {
	return &Repository{
		App: ac,
		DB:  dbrepo.NewPostgresRepo(db.SQL, ac),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter,
	r *http.Request) {

	// 29 Get an article from database
	// id, uid, title, content, err := m.DB.GetAnArticle()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("ID : ", id)
	// fmt.Println("User ID : ", uid)
	// fmt.Println("Title : ", title)
	// fmt.Println("Content : ", content)

	// 29 Get 3 posts
	// will hold struct with arrays of values
	var artList models.ArticleList
	artList, err := m.DB.Get3Articles()
	if err != nil {
		log.Println(err)
		return
	}

	for i := range artList.Content {
		fmt.Println(artList.Content[i])
	}

	m.App.Session.Put(r.Context(),
		"userid", "derekbanas")

	// Used to pass articleList to the template
	data := make(map[string]interface{})
	data["articleList"] = artList

	// Render template with data
	render.RenderTemplate(w, r, "home.page.tmpl",
		&models.PageData{
			Data: data,
		})

	// render.RenderTemplate(w, r, "home.page.tmpl",
	// 	&models.PageData{})
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

	// 28 If user isn't logged in redirect to login
	if !m.App.Session.Exists(r.Context(), "user_id") {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}

	var emptyArticle models.Article
	data := make(map[string]interface{})
	data["article"] = emptyArticle

	render.RenderTemplate(w, r, "make-post.page.tmpl",
		&models.PageData{
			Form: forms.New(nil),
			Data: data,
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

	// 29 Get user id of logged in user
	uID := (m.App.Session.Get(r.Context(), "user_id")).(int)

	// 26 Update this as a Post model
	article := models.Post{
		Title:   r.Form.Get("blog_title"),
		Content: r.Form.Get("blog_article"),
		UserID:  int(uID), // 29 Add user id to article
	}

	form := forms.New(r.PostForm)

	form.HasRequired("blog_title", "blog_article")

	form.MinLength("blog_title", 5, r)
	form.MinLength("blog_article", 5, r)

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

	// 26 After form validation write to
	// DB
	err = m.DB.InsertPost(article)
	if err != nil {
		log.Fatal(err)
	}

	m.App.Session.Put(r.Context(), "article", article)
	http.Redirect(w, r, "/article-received",
		http.StatusSeeOther)
}

func (m *Repository) ArticleReceived(w http.ResponseWriter, r *http.Request) {
	article, ok := m.App.Session.Get(r.Context(),
		"article").(models.Article)
	if !ok {
		log.Println("Can't get data from session")

		m.App.Session.Put(r.Context(), "error", "Can't get data from session")

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}
	data := make(map[string]interface{})
	data["article"] = article

	render.RenderTemplate(w, r, "article-received.page.tmpl",
		&models.PageData{
			Data: data,
		})
}

func (m *Repository) PageHandler(w http.ResponseWriter,
	r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "page.page.tmpl",
		&models.PageData{StrMap: strMap})
}

// 27 Function handles posts to login
func (m *Repository) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Prevents session fixation attacks
	// by renewing the token
	_ = m.App.Session.RenewToken(r.Context())

	// Parse the form
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Get email and password
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	// Verify the form has an email and
	// password and email is valid
	form := forms.New(r.PostForm)
	form.HasRequired("email", "password")
	form.IsEmail("email")

	// If not provided
	if !form.Valid() {
		// Redirect back to login page
		render.RenderTemplate(w, r, "login.page.tmpl", &models.PageData{
			Form: form,
		})
		return
	}

	// Authenticate user
	id, _, err := m.DB.AuthenticateUser(email, password)
	if err != nil {

		// Log error
		m.App.Session.Put(r.Context(), "error", "Invalid Email or Password")

		log.Fatal("Error logging in")

		// Redirect back to login screen
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Login the user by storing the id
	// in the session
	m.App.Session.Put(r.Context(), "user_id", id)

	m.App.Session.Put(r.Context(), "flash", "Valid Login")

	// Redirect back to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)

	fmt.Println("Logged in")
}

// 28 Handler for logout page
func (m *Repository) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Delete session
	_ = m.App.Session.Destroy(r.Context())

	// Prevents session fixation attacks
	// by renewing the token
	_ = m.App.Session.RenewToken(r.Context())

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
