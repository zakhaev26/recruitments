package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zakhaev26/recruitments/internal/constants"
	"github.com/zakhaev26/recruitments/internal/controllers"
	"github.com/zakhaev26/recruitments/internal/database"
	"github.com/zakhaev26/recruitments/internal/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	c := controllers.NewController(database.Db)
	r.HandleFunc("/api/v1/signup", c.SignUpController).Methods("post")
	r.HandleFunc("/api/v1/login", c.LoginController).Methods("post")
	r.HandleFunc("/api/v1/uploadResume", middleware.AuthorizationMiddleware(c.ResumeUploadController, []string{constants.APPLICANT})).Methods("post")
	r.HandleFunc("/api/v1/admin/job", middleware.AuthorizationMiddleware(c.CreateJobController, []string{constants.ADMIN})).Methods("post")
	r.HandleFunc("/api/v1/admin/job/{job_id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("/api/v1/admin/applicants", middleware.AuthorizationMiddleware(c.ApplicantsController, []string{constants.ADMIN})).Methods("get")
	r.HandleFunc("/api/v1/admin/applicant/{applicant_id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("get")
	r.HandleFunc("/api/v1/jobs", middleware.AuthorizationMiddleware(c.GetJobsCotroller, []string{"*"})).Methods("get")
	r.HandleFunc("/api/v1/jobs/apply", middleware.AuthorizationMiddleware(c.ApplyJobController, []string{constants.APPLICANT})).Methods("get")
	return r
}
