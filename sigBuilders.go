package icecubeapp

import (
	"net/http"

	"appengine"
	"appengine/datastore"
)

func createReadSig(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

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
	}
	if r.FormValue("sigtype") == "integer" {
		p.DataType = "%d"
	} else if r.FormValue("sigtype") == "float" {
		p.DataType = "%f"
	} else if r.FormValue("sigtype") == "string" {
		p.DataType = "%s"
	}
	key := datastore.NewIncompleteKey(c, "WriteSig", iceCubeKey(c))
	_, err := datastore.Put(c, key, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
