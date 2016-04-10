package icecubeapp

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/signedout", signedout)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	url, _ := user.LogoutURL(c, "/signedout")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func signedout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, signedoutForm)
}

const signedoutForm = `
<html>
  <head>
    <title>IceCube App</title>
  </head>
  <body>
    <p>Thanks for visiting!</p>
    <a href="/">Sign in again</a>
  </body>
</html>
`
