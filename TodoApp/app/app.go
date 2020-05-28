package app

import (
	"learngo/TodoApp/model"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var rd *render.Render = render.New()
var store *sessions.CookieStore

// AHandler is AHandler
type AHandler struct {
	http.Handler
	db model.DBHandler
}

var getSessionID = func(r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}

	// Set some session values.
	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

func (a *AHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)
	list := a.db.GetTodos(sessionID)
	rd.JSON(w, http.StatusOK, list)
}

func (a *AHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)
	name := r.FormValue("name")
	todo := a.db.AddTodo(sessionID, name)

	rd.JSON(w, http.StatusCreated, todo)
}

type success struct {
	Success bool `json:"success"`
}

func (a *AHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, success{true})
	} else {
		rd.JSON(w, http.StatusOK, success{false})
	}

}

func (a *AHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, success{true})
	} else {
		rd.JSON(w, http.StatusOK, success{false})
	}

}

// Close Function
func (a *AHandler) Close() {
	a.db.Close()
}

// CheckSignin is Signin Check
func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// if request URL is /signin.html, then next()
	if strings.Contains(r.URL.Path, "/signin") || strings.Contains(r.URL.Path, "/auth/google") {
		next(w, r)
		return
	}

	// if user already signed in
	sessionID := getSessionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}

	// if not user sign in
	// redirect signin.html
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

// MakeHandler is app handler
func MakeHandler(filepath string) *AHandler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(CheckSignin),
		negroni.NewStatic(http.Dir("public")),
	)

	n.UseHandler(r)

	a := &AHandler{
		Handler: n,
		db:      model.NewDBHandler(filepath),
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	googleOauthConfig = oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleLoginCallback)
	r.HandleFunc("/", a.indexHandler)
	return a
}
