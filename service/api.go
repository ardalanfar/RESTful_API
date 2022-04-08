package service

import (
	"log"
	"net/http"

	"RESTful_API/app/handler"
	"RESTful_API/database"
	"RESTful_API/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.DB = database.Connect()

	models.UsersMigrate(app.DB)
	models.CourseMigrate(app.DB)
	models.TeacherMigrate(app.DB)

	app.SetRouters()
}

func (app *App) SetRouters() {
	app.Get("/users", app.GetUsers)
	app.Post("/user/{id}", app.GetUser)
	app.Post("/user", app.CreateUser)

	app.Get("/courses", app.GetCourses)
	app.Post("/course/{id}", app.GetCourse)
	app.Post("/course", app.CreateCourse)

	app.Post("/student/{name_c}", app.Get_UserCourse)

	app.Get("/report", app.Get_Report)
}

/*----------------------------------------------------------*/

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

/*----------------------------------------------------------*/
//users

func (app *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	handler.Getusers(app.DB, w, r)
}

func (app *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.Getuser(app.DB, w, r)
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.Createuser(app.DB, w, r)
}

/*----------------------------------------------------------*/
//course

func (app *App) GetCourses(w http.ResponseWriter, r *http.Request) {
	handler.Getcourses(app.DB, w, r)
}

func (app *App) GetCourse(w http.ResponseWriter, r *http.Request) {
	handler.Getcourse(app.DB, w, r)
}

func (app *App) CreateCourse(w http.ResponseWriter, r *http.Request) {
	handler.Createcourse(app.DB, w, r)
}

/*----------------------------------------------------------*/
//controllers

//Show students in course
func (app *App) Get_UserCourse(w http.ResponseWriter, r *http.Request) {
	handler.Get_usercourse(app.DB, w, r)
}

//Show students and teachers and courses
func (app *App) Get_Report(w http.ResponseWriter, r *http.Request) {
	handler.Get_report(app.DB, w, r)
}

/*----------------------------------------------------------*/

//Run the app on it's router
func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}

/*----------------------------------------------------------*/
