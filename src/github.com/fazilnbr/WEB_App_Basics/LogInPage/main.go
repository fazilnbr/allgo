package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var client *redis.Client

var store = sessions.NewCookieStore([]byte("super-secret-password"))

var temp *template.Template

type Page struct {
	Status  bool
	Header1 interface{}
	Valid   bool
}

var P = Page{
	Status: false,
}

// var user map [string]string
type usr struct {
	username, pswd string
}

var user = map[string]string{
	"un":       "fa_z_il_nbr",
	"password": "123456",
}
var uer usr

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	fmt.Println(client)
	temp = template.Must(template.ParseGlob("template/*.html"))
	// fmt.Println(temp)
	r := mux.NewRouter()
	r.HandleFunc("/", home_handler).Methods("GET")
	r.HandleFunc("/login", login_handler).Methods("GET")
	r.HandleFunc("/tologin", logout_handler).Methods("GET")
	r.HandleFunc("/", check_handler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

func Middleware(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session-name")

	if session.Values["id"] == nil {
		return false
	}
	P.Header1 = session.Values["id"]
	return true

}

func home_handler(w http.ResponseWriter, r *http.Request) {
	// ok:=Middleware(w,r)
	session, err := store.Get(r, "session-name")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("on home", session)
	u := session.Values
	if session.Values != nil {
		temp.ExecuteTemplate(w, "home.html", u["name"])
	} else {
		temp.ExecuteTemplate(w, "home.html", nil)
	}
}

func login_handler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("fff", session)
	// session.Options.MaxAge = -1
	session.Save(r, w)
	fmt.Println("dlog ", session)

	u := session.Values
	if session.Values != nil {
		temp.ExecuteTemplate(w, "home.html", u["name"])
	} else {
		temp.ExecuteTemplate(w, "login.html", nil)
	}

}

func logout_handler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("fff", session)
	session.Options.MaxAge = -1
	session.Save(r, w)
	u := session.Values
	if session.Values != nil {
		temp.ExecuteTemplate(w, "login.html", nil)
	} else {
		temp.ExecuteTemplate(w, "home.html", u["name"])
	}
}

func check_handler(w http.ResponseWriter, r *http.Request) {

	//fetching data from form

	uer.username = r.FormValue("username")
	uer.pswd = r.FormValue("password")

	//session creations

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := session.Values
	//form validation

	//if case
	if user["un"] == uer.username && user["password"] == uer.pswd {

		session.Options = &sessions.Options{
			Path:   "/",
			MaxAge: 86400 * 7,
			//MaxAge:   5,
			HttpOnly: true,
		}

		if uer.username != "" {
			session.Values["name"] = uer.username
		}

		//session saving is important

		session.Save(r, w)
		fmt.Println("on login", session)

		//give direction to home
		temp.ExecuteTemplate(w, "home.html", u["name"])
	} else {
		temp.ExecuteTemplate(w, "login.html", nil)
	}

}
