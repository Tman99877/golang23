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
	Value1 string `json:"key1"`
	Value2 string `json:"key2"`
	jsonNested JsonNested `json: "JsonNested`
}

type JsonNested struct{
	NestedValue1 string `json"nestedKey1"`
	NestedValue2 string `json"nestedKey2"`
}


func main() {

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	nested := JsonNested{
		NestedValue1: "first nested value",
		NestedValue2: "Second nested value",
	}


	jsonResp  := JsonResponse{
		Value1:  "Some Data",
		Value2: "Other Data",
		jsonNested: nested,

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
	http.HandleFunc("/jsonResponse", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(jsonResp)
	})
	
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}



		


	

