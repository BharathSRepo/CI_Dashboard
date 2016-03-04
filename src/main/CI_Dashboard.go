package main

import (
	"fmt"
	"net/http"
	"controllers/populatetemplates"
)

func main() {
	fmt.Printf("Fisrt code for the dashboard!")
	
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/Page1", Page1)	
	http.ListenAndServe(":9090",nil)
}

func HomePage(w http.ResponseWriter, req *http.Request) {
	templates := populatetemplates.PopulateTemplates() // fetching all the list of html files from templates folder
	template := templates.Lookup("index.html")      // Selecting index.html from list of all html templates
	if template != nil {
		template.Execute(w, nil) // redirecting to HomePage.html page with list of Boxes Information
	} else {
		http.Redirect(w, req, "/errorPage", http.StatusFound) // If template lookup fails to find HomePage.html it will redirect to ErrorPage
	}
}

func Page1(w http.ResponseWriter, req *http.Request) {
	templates := populatetemplates.PopulateTemplates() // fetching all the list of html files from templates folder
	template := templates.Lookup("page1.html")      // Selecting index.html from list of all html templates
	if template != nil {
		template.Execute(w, nil) // redirecting to HomePage.html page with list of Boxes Information
	} else {
		http.Redirect(w, req, "/", http.StatusFound) // If template lookup fails to find HomePage.html it will redirect to ErrorPage
	}
}
