package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)


type Welcome struct {
	Name string
	Time string
}


type JsonResponse struct{
	Fname string `json:"First Name"`
	Lname string `json:"Last Name"`
	JsonInfo JsonInfo `json: "Client  info`
}

type JsonInfo struct{
	Maddress string `json"Mailing Address"`
	Cinfo string `json"Contact Info"`
}


func main() {

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	nested := JsonInfo{
		Maddress: "200 49th St Apt 50 Columbus,Ga 31906",
		Cinfo: "706-999-0156, UCantCMe99@gmail.com",
	}


	jsonResp  := JsonResponse{
		Fname:  "John",
		Lname: "Cena",
		JsonInfo: nested,

	}


	http.Handle("/static/", 
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) 

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if name := r.FormValue("name"); name != "" {
			welcome.Name=name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html",welcome); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/contactinfo", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(jsonResp)
	})
	
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}



		


	

