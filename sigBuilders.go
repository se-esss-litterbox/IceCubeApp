package icecubeapp

import (
	"net/http"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

func createReadSig(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	p := ReadSig{
		SigName:   r.FormValue("signame"),
		SerialStr: r.FormValue("serialcommand"),
	}
	if r.FormValue("sigtype") == "integer" {
		p.DataType = "%d"
	} else if r.FormValue("sigtype") == "float" {
		p.DataType = "%f"
	} else if r.FormValue("sigtype") == "string" {
		p.DataType = "%s"
	}
	key := datastore.NewIncompleteKey(c, "ReadSig", iceCubeKey(c, u.String()))
	_, err := datastore.Put(c, key, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func createWriteSig(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if r.PostFormValue("action") == "Delete" {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	p := WriteSig{
		SigName:   r.FormValue("signame"),
		SerialStr: r.FormValue("serialcommand"),
	}
	if r.FormValue("sigtype") == "integer" {
		p.DataType = "%d"
	} else if r.FormValue("sigtype") == "float" {
		p.DataType = "%f"
	} else if r.FormValue("sigtype") == "string" {
		p.DataType = "%s"
	}
	key := datastore.NewIncompleteKey(c, "WriteSig", iceCubeKey(c, u.String()))
	_, err := datastore.Put(c, key, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
