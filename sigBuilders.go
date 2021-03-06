package icecubeapp

import (
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

func createSig(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("rw") == "r" {
		createReadSig(w, r)
		return
	}
	if r.FormValue("rw") == "w" {
		createWriteSig(w, r)
		return
	}
}

func createReadSig(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if r.PostFormValue("updateOrDelete") == "Delete" {
		var sigs []ReadSig
		//sig := r.FormValue("signame")
		key := r.FormValue("key")
		q := datastore.NewQuery("ReadSig").Filter("KeyStr =", key).KeysOnly()
		keys, _ := q.GetAll(c, sigs)
		datastore.Delete(c, keys[0])
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.PostFormValue("updateOrDelete") == "Update" {
		var sigs []ReadSig
		//sig := r.FormValue("signame")
		key := r.FormValue("key")
		q := datastore.NewQuery("ReadSig").Filter("KeyStr =", key).KeysOnly()
		keys, _ := q.GetAll(c, sigs)
		datastore.Delete(c, keys[0])
	}

	p := ReadSig{
		SigName:   r.FormValue("signame"),
		SerialStr: r.FormValue("serialcommand"),
		When:      time.Now(),
	}
	if r.FormValue("sigtype") == "integer" {
		p.DataType = "%d"
	} else if r.FormValue("sigtype") == "float" {
		p.DataType = "%f"
	} else if r.FormValue("sigtype") == "string" {
		p.DataType = "%s"
	}
	key := datastore.NewIncompleteKey(c, "ReadSig", iceCubeKey(c, u.String()))
	completeKey, err := datastore.Put(c, key, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.Key = completeKey
	p.KeyStr = completeKey.String()
	_, err = datastore.Put(c, completeKey, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func createWriteSig(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if r.PostFormValue("updateOrDelete") == "Delete" {
		var sigs []WriteSig
		//sig := r.FormValue("signame")
		key := r.FormValue("key")
		q := datastore.NewQuery("WriteSig").Filter("KeyStr =", key).KeysOnly()
		keys, _ := q.GetAll(c, sigs)
		datastore.Delete(c, keys[0])
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.PostFormValue("updateOrDelete") == "Update" {
		var sigs []WriteSig
		//sig := r.FormValue("signame")
		key := r.FormValue("key")
		q := datastore.NewQuery("WriteSig").Filter("KeyStr =", key).KeysOnly()
		keys, _ := q.GetAll(c, sigs)
		datastore.Delete(c, keys[0])
	}

	p := WriteSig{
		SigName:   r.FormValue("signame"),
		SerialStr: r.FormValue("serialcommand"),
		When:      time.Now(),
	}
	if r.FormValue("sigtype") == "integer" {
		p.DataType = "%d"
	} else if r.FormValue("sigtype") == "float" {
		p.DataType = "%f"
	} else if r.FormValue("sigtype") == "string" {
		p.DataType = "%s"
	}
	key := datastore.NewIncompleteKey(c, "WriteSig", iceCubeKey(c, u.String()))
	completeKey, err := datastore.Put(c, key, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.Key = completeKey
	p.KeyStr = completeKey.String()
	_, err = datastore.Put(c, completeKey, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
