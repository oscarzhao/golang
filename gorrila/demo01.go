package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func init() {
	flag.Set("alsologtostderr", "true") //  print log to stdout
	flag.Parse()
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func MyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, found := vars["name"]
	if found == false {
		glog.Errorf("invalid request name: %v\n", name)
		w.WriteHeader(406)
		w.Write([]byte("invalid request"))
		return
	}

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := store.Get(r, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if session.IsNew {
		glog.Infof("session %s is new\n", name)
		// Set some session values.
		session.Values["name"] = name
		session.ID = name
		session.Save(r, w)
		w.Write([]byte("session " + name + " saved\n"))
	} else {
		glog.Infof("session %s already exist, details:%v\n", session.ID, session.Values)
		w.Write([]byte("session " + name + " already exists\n"))
	}
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/mine/{name}", MyHandler)

	// Bind to a port and pass our router in
	http.ListenAndServe(":8000", r)

	defer store.Close()
}
