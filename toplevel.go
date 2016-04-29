package icecubeapp

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

var baseTempl = template.Must(template.ParseFiles(
	"base.html",
	"srcTempl.html",
	"protoTempl.tmp",
	"dbTempl.tmp",
	"arduinoTempl.tmp",
))
var protoTempl = template.Must(template.ParseFiles("protoFile.tmp", "protoTempl.tmp"))
var dbTempl = template.Must(template.ParseFiles("dbFile.tmp", "dbTempl.tmp"))

// ReadSig is a read signal
type ReadSig struct {
	SigName   string
	SerialStr string
	DataType  string
	Key       *datastore.Key
	KeyStr    string
	When      time.Time
}

// WriteSig is a write signal
type WriteSig struct {
	SigName   string
	SerialStr string
	DataType  string
	Key       *datastore.Key
	KeyStr    string
	When      time.Time
}

// Signal contains a bunch of read & write signals
type Signal struct {
	Read  []ReadSig
	Write []WriteSig
}

// iceCubeKey returns the key used for all IceCube entries.
func iceCubeKey(c appengine.Context, username string) *datastore.Key {
	return datastore.NewKey(c, "IceCube", username, 0, nil)
}

func init() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/signedout", signedout)
	http.HandleFunc("/home", home)
	http.HandleFunc("/createSig", createSig)
	http.HandleFunc("/arduino.proto", serveProtoFile)
	http.HandleFunc("/arduino.db", serveDBFile)
	http.HandleFunc("/download.tar", tarit)
}

func serveDBFile(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

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

	if err := dbTempl.Execute(w, sigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveProtoFile(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	// Get ReadSigs from the persistent datastore
	q := datastore.NewQuery("ReadSig").Ancestor(iceCubeKey(c, u.String())).Order("When").Limit(50)
	readSigs := make([]ReadSig, 0, 10)
	if _, err := q.GetAll(c, &readSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get WriteSigs from the persistent datastore
	q = datastore.NewQuery("WriteSig").Ancestor(iceCubeKey(c, u.String())).Order("When").Limit(50)
	writeSigs := make([]WriteSig, 0, 10)
	if _, err := q.GetAll(c, &writeSigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sigs := Signal{
		Read:  readSigs,
		Write: writeSigs,
	}

	if err := protoTempl.Execute(w, sigs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

func signedout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, signedoutForm)
}
