package web

import (
	"net/http"

	"ecksbee.com/telefacts-taxonomy-package/internal/cache"
	"github.com/gorilla/mux"
)

func Namespaces() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
			return
		}
		vars := mux.Vars(r)
		ns := vars["ns"]
		if len(ns) <= 0 {
			http.Error(w, "Error: invalid ns '"+ns+"'", http.StatusBadRequest)
			return
		}
		data, err := cache.MarshalNamespace(ns)
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func RelationshipSets() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
			return
		}
		vars := mux.Vars(r)
		roleuri := vars["roleuri"]
		if len(roleuri) <= 0 {
			http.Error(w, "Error: invalid roleuri '"+roleuri+"'", http.StatusBadRequest)
			return
		}
		hash := vars["hash"]
		if len(hash) <= 0 {
			http.Error(w, "Error: invalid roote", http.StatusBadRequest)
			return
		}
		data, err := cache.MarshalRelationshipSet(roleuri)
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.Path("/namespaces").HandlerFunc(Namespaces()).Methods("GET")
	r.Path("/relationshipsets").HandlerFunc(RelationshipSets()).Methods("GET")
	return r
}
