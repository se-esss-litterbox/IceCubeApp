package icecubeapp

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

func home(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	url, _ := user.LogoutURL(c, "/signedout")
	baseTempl.Parse(fmt.Sprintf(`{{define "logout"}}<a href="%s">sign out</a><br>{{end}}`, url))

	// Get ReadSigs from the persistent datastore
	q := datastore.NewQuery("ReadSig").Ancestor(iceCubeKey(c, u.String())).Order("SigName").Limit(10)
	readSigs := make([]ReadSig, 0, 10)
	if _, err := q.GetAll(c, &readSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get WriteSigs from the persistent datastore
	q = datastore.NewQuery("WriteSig").Ancestor(iceCubeKey(c, u.String())).Order("SigName").Limit(10)
	writeSigs := make([]WriteSig, 0, 10)
	if _, err := q.GetAll(c, &writeSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sigs := Signal{
		Read:  readSigs,
		Write: writeSigs,
	}

	if err := baseTempl.Execute(w, sigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
