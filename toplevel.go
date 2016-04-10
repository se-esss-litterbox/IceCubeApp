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
	fmt.Fprintf(w, sigCreateForm)
}

func signedout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, signedoutForm)
}

const signedoutForm = `
<p>Thanks for visiting!</p>
<a href="/">Sign in again</a>
`

const stylesForm = `
<style type="text/css">
    .fieldset-auto-width {
         display: inline-block;
    }
    ul#menu li {
    display:inline-block;
    }
</style>
`

const sigCreateForm = stylesForm +
	`<fieldset class="fieldset-auto-width">
  <legend><h2>New Signal Creator</h2></legend>
  <ul id="menu">
  <li>` + readSigCreateForm + `</li>
  <li>` + writeSigCreateForm + `</li>
  </ul>
  </fieldset>
  `

const readSigCreateForm = `
<form action="/" method="post">
  <fieldset class="fieldset-auto-width">
    <legend><h2>Read</h2></legend>
    Signal name:<br>
    <input type="text" name="signame"></input><br><br>
    Serial string:<br>
    <input type="text" name="serialcommand"></input><br><br>` +
	sigSelectForm + "<br><br>" +
	`<input type="submit" value="Create">
    </fieldset>
</form>`

const writeSigCreateForm = `
<form action="/" method="post">
  <fieldset class="fieldset-auto-width">
    <legend><h2>Write</h2></legend>
    Signal name:<br>
    <input type="text" name="signame"></input><br><br>
    Serial string:<br>
    <input type="text" name="serialcommand"></input><br><br>` +
	sigSelectForm + "<br><br>" +
	`<input type="submit" value="Create">
    </fieldset>
</form>`

const sigSelectForm = `
  <select name="sigtype">
    <option value="integer">Integer</option>
    <option value="float">Float</option>
    <option value="string">C-style str</option>
  </select>`
