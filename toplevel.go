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
	http.HandleFunc("/home", home)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, "/home")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	http.Redirect(w, r, "/home", http.StatusFound)
}

func home(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, `<h1>Welcome home, %s! (<a href="%s">sign out</a>)</h1>`, u, url)
	fmt.Fprintf(w, readSigCreateForm)
}

func signedout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, signedoutForm)
}

const signedoutForm = `
<p>Thanks for visiting!</p>
<a href="/">Sign in again</a>
`

const readSigCreateForm = `
<style type="text/css">
    .fieldset-auto-width {
         display: inline-block;
    }
</style>
<form action="/" method="post">
  <fieldset class="fieldset-auto-width">
    <legend>New Read Signal</legend>
    Signal name:<br>
    <input type="text" name="signame"></input><br><br>
    Serial command string:<br>
    <input type="text" name="serialcommand"></input><br><br>
    <input type="submit" value="Create">
    </fieldset>
</form>
`
