package main

import (
	"9-data-modelling/connection"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Struct
type Project struct {
	Id                     int
	ProjectName            string
	ProjectStartDate       time.Time
	ProjectEndDate         time.Time
	ProjectStartDateString string
	ProjectEndDateString   string
	ProjectDuration        string
	ProjectDescription     string
	ProjectTechnologies    []string
	ProjectImage           string
}

var ProjectList = []Project{}

// MAIN
func main() {
	route := mux.NewRouter()

	//CONNECTION
	connection.DatabaseConnect()

	// ROUTE PUBLIC FOLDER
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// ROUTE RENDER HTML
	route.HandleFunc("/", HomePage).Methods("GET")
	route.HandleFunc("/contact", ContactPage).Methods("GET")
	route.HandleFunc("/project", ProjectPage).Methods("GET")
	route.HandleFunc("/detail-project/{index}", ProjectDetail).Methods("GET")

	// CREATE PROJECT
	route.HandleFunc("/project/create", CreateProject).Methods("POST")
	// UPDATE PROJECT
	route.HandleFunc("/update-project/{index}", UpdateProject).Methods("GET")
	// DELETE PROJECT
	route.HandleFunc("/delete-project/{index}", DeleteProject).Methods("GET")

	// PORT HANDLING
	fmt.Println(("Server running on port 5000"))
	http.ListenAndServe("localhost:5000", route)
}

// RENDER HOME PAGE
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	} else {

		rows, _ := connection.Conn.Query(context.Background(), `SELECT "Id", "ProjectName", "ProjectStartDate", "ProjectEndDate", "ProjectDuration", "ProjectDescription", "ProjectTechnologies", "ProjectImage" FROM tb_project`)
		var renderData []Project
		each := Project{}
		for rows.Next() {

			err := rows.Scan(&each.Id, &each.ProjectName, &each.ProjectStartDate, &each.ProjectEndDate, &each.ProjectDuration, &each.ProjectDescription, &each.ProjectTechnologies, &each.ProjectImage)

			if err != nil {
				fmt.Println(err.Error())
				return
			} else {

				each := Project{
					Id:                  each.Id,
					ProjectName:         each.ProjectName,
					ProjectDuration:     GetDuration(each.ProjectStartDate, each.ProjectEndDate),
					ProjectDescription:  each.ProjectDescription,
					ProjectTechnologies: each.ProjectTechnologies,
					ProjectImage:        each.ProjectImage,
				}

				renderData = append(renderData, each)
			}
		}

		response := map[string]interface{}{
			"renderData": renderData,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}
}

// RENDER CONTACT PAGE
func ContactPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

// RENDER PROJECT PAGE
func ProjectPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

// RENDER PROJECT DETAIL
func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/project-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	} else {
		var renderDetail = Project{}
		index, _ := strconv.Atoi(mux.Vars(r)["index"])

		for i, data := range ProjectList {
			if index == i {
				renderDetail = Project{
					ProjectName:         data.ProjectName,
					ProjectStartDate:    data.ProjectStartDate,
					ProjectEndDate:      data.ProjectEndDate,
					ProjectDuration:     data.ProjectDuration,
					ProjectDescription:  data.ProjectDescription,
					ProjectTechnologies: data.ProjectTechnologies,
				}
			}
		}
		data := map[string]interface{}{
			"renderDetail": renderDetail,
		}
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, data)
	}
}

// CREATE PROJECT
func CreateProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	} else {
		// projectName := r.PostForm.Get("project-name")
		// projectStartDate := r.PostForm.Get("date-start")
		// projectEndDate := r.PostForm.Get("date-end")
		// projectDescription := r.PostForm.Get("project-description")
		// projectUseNodeJS := r.PostForm.Get("nodejs")
		// projectUseReactJS := r.PostForm.Get("reactjs")
		// projectUseNextJS := r.PostForm.Get("nextjs")
		// projectUseTypeScript := r.PostForm.Get("typescript")

		// var newProject = Project{
		// 	ProjectName:         projectName,
		// 	ProjectStartDate:    FormatDate(projectStartDate),
		// 	ProjectEndDate:      FormatDate(projectEndDate),
		// 	ProjectDuration:     GetDuration(projectStartDate, projectEndDate),
		// 	ProjectDescription:  projectDescription,
		// 	ProjectTechnologies: []string{projectUseNodeJS, projectUseReactJS, projectUseNextJS, projectUseTypeScript},
		// }

		// ProjectList = append(ProjectList, newProject)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

// UPDATE PROJECT
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//tmpl, err := template.ParseFiles("views/update-project.html")
	//if err != nil {
	// w.WriteHeader(http.StatusInternalServerError)
	// w.Write([]byte("message : " + err.Error()))
	// return
	//} // else {
	// var updateData = Project{}
	// index, _ := strconv.Atoi(mux.Vars(r)["index"])

	// for i, data := range ProjectList {
	// 	if index == i {
	// 		updateData = Project{
	// 			ProjectName:         data.ProjectName,
	// 			ProjectStartDate:    ReturnDate(data.ProjectStartDate),
	// 			ProjectEndDate:      ReturnDate(data.ProjectEndDate),
	// 			ProjectDescription:  data.ProjectDescription,
	// 			ProjectTechnologies: data.ProjectTechnologies,
	// 		}
	// 		ProjectList = append(ProjectList[:index], ProjectList[index+1:]...)
	// 	}
	// }
	// data := map[string]interface{}{
	// 	"updateData": updateData,
	//}
	//w.WriteHeader(http.StatusOK)
	//tmpl.Execute(w, data)
	//}
}

// DELETE PROJECT
func DeleteProject(w http.ResponseWriter, r *http.Request) {

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	ProjectList = append(ProjectList[:index], ProjectList[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}

// GET DURATION
func GetDuration(startDate time.Time, endDate time.Time) string {

	margin := endDate.Sub(startDate).Hours() / 24
	var duration string

	if margin >= 30 {
		if (margin / 30) == 1 {
			duration = "1 Month"
		} else {
			duration = strconv.Itoa(int(margin/30)) + " Months"
		}
	} else {
		if margin <= 1 {
			duration = "1 Day"
		} else {
			duration = strconv.Itoa(int(margin)) + " Days"
		}
	}

	return duration

}

// CHANGE DATE FORMAT
func FormatDate(InputDate time.Time) string {

	Formated := InputDate.Format("02 January 2006")

	return Formated
}

// RETURN DATE FORMAT
func ReturnDate(InputDate time.Time) string {

	Formated := InputDate.Format("2006-01-02")

	return Formated
}
