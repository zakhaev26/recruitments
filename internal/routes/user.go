package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("api/v1/signup", func(w http.ResponseWriter, r *http.Request) {}).Methods("post")
	r.HandleFunc("api/v1/login", func(w http.ResponseWriter, r *http.Request) {}).Methods("post")
	r.HandleFunc("api/v1/uploadResume", func(w http.ResponseWriter, r *http.Request) {}).Methods("post")
	r.HandleFunc("api/v1/admin/job", func(w http.ResponseWriter, r *http.Request) {}).Methods("post")
	r.HandleFunc("api/v1/admin/job/{job_id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("api/v1/admin/applicants", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("api/v1/admin/applicant/{applicant_id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("api/v1/admin/jobs", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("api/v1/admin/jobs/apply?job_id={job_id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	return r
}
