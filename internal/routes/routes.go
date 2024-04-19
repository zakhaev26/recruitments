package routes

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/zakhaev26/recruitments/internal/constants"
	"github.com/zakhaev26/recruitments/internal/controllers"
	"github.com/zakhaev26/recruitments/internal/database"
	"github.com/zakhaev26/recruitments/internal/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	authController := controllers.NewController(database.Db)
	log.Info(authController)
	r.HandleFunc("/api/v1/signup", authController.SignUpController).Methods("post")
	r.HandleFunc("/api/v1/login", authController.LoginController).Methods("post")
	r.HandleFunc("/api/v1/uploadResume", middleware.AuthorizationMiddleware(func(w http.ResponseWriter, r *http.Request) {}, []string{constants.USER})).Methods("post")
	r.HandleFunc("/api/v1/admin/job", func(w http.ResponseWriter, r *http.Request) {}).Methods("post")
	r.HandleFunc("/api/v1/admin/job/{job_id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("/api/v1/admin/applicants", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("/api/v1/admin/applicant/{applicant_id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("/api/v1/admin/jobs", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("/api/v1/admin/jobs/apply?job_id={job_id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	return r
}
