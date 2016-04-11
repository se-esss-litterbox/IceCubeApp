package icecubeapp

import (
	"fmt"
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
	fmt.Fprintf(w, `Source code found <a href="%s">on GitHub</a><br>`, "https://github.com/se-esss-litterbox/IceCubeApp")
	fmt.Fprintf(w, `<a href="%s">sign out</a><br>`, url)
	fmt.Fprint(w, sigCreateForm)

	fmt.Fprint(w, `<h3>Protocol file</h3>`)
	if err := readSigsProtoTemplate.Execute(w, readSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := writeSigsProtoTemplate.Execute(w, writeSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, `<h3>EPICS DB file</h3>`)
	if err := readSigsDBTemplate.Execute(w, readSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := writeSigsDBTemplate.Execute(w, writeSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func signedout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, signedoutForm)
}
