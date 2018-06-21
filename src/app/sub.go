package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var tpl *template.Template

// func ParseTemplates() *template.Template {
// 	templ := template.New("")
// 	err := filepath.Walk("./src/app", func(path string, info os.FileInfo, err error) error {
// 		if strings.Contains(path, ".js") {
// 			_, err = templ.ParseFiles(path)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}

// 		return err
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	return templ
// }
func init() {

	// repoFrontend := "src/*.js"
	// http.Handle("/static/", http.StripPrefix("/static/",
	// 	http.FileServer(http.Dir(repoFrontend))))
	db, err = sql.Open("mysql", "root:root@/test")
	_, err := template.New("").ParseFiles("index.js", "success.js")
	if err != nil {
		log.Fatal("Error loading templates:" + err.Error())
	}
	//tpl = template.New("").ParseFiles("src/*")
}

// func userForm(w http.ResponseWriter, req *http.Request) {

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
func personalDetails(res http.ResponseWriter, req *http.Request) {
	err = tpl.ExecuteTemplate(res, "index.js", nil)
	if req.Method == http.MethodPost {
		fname := req.FormValue("name")

		_, err = db.Exec(
			"INSERT INTO user (name) VALUES (?)",
			fname,
		)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	http.Error(res, "Method Not Supported", http.StatusMethodNotAllowed)

}

func main() {

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	//http.HandleFunc("/", userForm)
	http.HandleFunc("/index", personalDetails)

	http.ListenAndServe(":8080", nil)
}
