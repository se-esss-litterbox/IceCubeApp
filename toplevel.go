package icecubeapp

import (
	"fmt"
	"html/template"
	"net/http"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

// ReadSig is a read signal
type ReadSig struct {
	SigName   string
	SerialStr string
	DataType  string
}

// WriteSig is a write signal
type WriteSig struct {
	SigName   string
	SerialStr string
	DataType  string
}

// testingKey returns the key used for all testing entries.
func iceCubeKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "IceCube", "iceCube_testing", 0, nil)
}

func init() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/signedout", signedout)
	http.HandleFunc("/home", home)
	http.HandleFunc("/createReadSig", createReadSig)
	http.HandleFunc("/createWriteSig", createWriteSig)
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
	// Only allowed to proceed when signed in
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

	q := datastore.NewQuery("ReadSig").Ancestor(iceCubeKey(c)).Order("SigName").Limit(10)
	readSigs := make([]ReadSig, 0, 10)
	if _, err := q.GetAll(c, &readSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q = datastore.NewQuery("WriteSig").Ancestor(iceCubeKey(c)).Order("SigName").Limit(10)
	writeSigs := make([]WriteSig, 0, 10)
	if _, err := q.GetAll(c, &writeSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url, _ := user.LogoutURL(c, "/signedout")
	fmt.Fprintf(w, `<a href="%s">sign out</a><br>`, url)
	fmt.Fprint(w, sigCreateForm)

	if err := definedSigsTemplate.Execute(w, readSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := definedSigsTemplate.Execute(w, writeSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createReadSig(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	p := ReadSig{
		SigName:   r.FormValue("signame"),
		SerialStr: r.FormValue("serialcommand"),
		DataType:  r.FormValue("sigtype"),
	}
	key := datastore.NewIncompleteKey(c, "ReadSig", iceCubeKey(c))
	_, err := datastore.Put(c, key, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func createWriteSig(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	p := WriteSig{
		SigName:   r.FormValue("signame"),
		SerialStr: r.FormValue("serialcommand"),
		DataType:  r.FormValue("sigtype"),
	}
	key := datastore.NewIncompleteKey(c, "WriteSig", iceCubeKey(c))
	_, err := datastore.Put(c, key, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func signedout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, signedoutForm)
}

var definedSigsTemplate = template.Must(template.New("sigs").Parse(
	`<br>
    {{range .}}
      <b>{{.SigName}}</b>
      {{.SerialStr}}
      {{.DataType}}<br>
    {{end}}
    `))

const signedoutForm = `
<p>Thanks for visiting!</p>
<a href="/">Sign in again</a>
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
<form action="/createReadSig" method="post">
  <fieldset class="fieldset-auto-width">
    <legend><h2>Read</h2></legend>
    Signal name:<br>
    <input type="text" name="signame" required></input><br><br>
    Serial string:<br>
    <input type="text" name="serialcommand" required></input><br><br>` +
	sigSelectForm + "<br><br>" +
	`<input type="submit" value="Create">
    </fieldset>
</form>`

const writeSigCreateForm = `
<form action="/createWriteSig" method="post">
  <fieldset class="fieldset-auto-width">
    <legend><h2>Write</h2></legend>
    Signal name:<br>
    <input type="text" name="signame" required></input><br><br>
    Serial string:<br>
    <input type="text" name="serialcommand" required></input><br><br>` +
	sigSelectForm + "<br><br>" +
	`<input type="submit" value="Create">
    </fieldset>
</form>`

const sigSelectForm = `
  Data Type:<br>
  <select name="sigtype">
    <option value="integer">Integer</option>
    <option value="float">Float</option>
    <option value="string">C-style str</option>
  </select>`

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
