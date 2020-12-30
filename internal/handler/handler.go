package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"scb-recipe-app/internal/service"
)

func RecipeStatsHandler(res http.ResponseWriter, req *http.Request)  {
	file, _, err := req.FormFile("recipesData")
	postCode := req.FormValue("postCode")
	recipeName := req.FormValue("recipeName")
	startTime := req.FormValue("deliveryStartTime")
	endTime := req.FormValue("deliveryEndTime")
	recipeName = recipeName
	postCode = postCode
	if err != nil {
		log.Print("error retrieving the file")
		return
	}
	defer file.Close()

	var service service.Recipe
	var data interface{}
	if recipeName != "" {
		data = service.FindStatsByRecipeName(file, recipeName);
	} else if postCode != ""  && startTime != "" && endTime != "" {
		data = service.FindStatsByPostCodeAndTime(file, postCode, startTime, endTime)
	} else {
		data = service.CalculateStats(file)
	}

	res.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(data)
	res.Write(response)
	return
}

func RecipeTemplateHandler(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			file := filepath.Join("./internal/template", "file.gtpl")
			recipeTemplate, err := template.ParseFiles(file)
			if err != nil {
				log.Fatal("error parsing recipe template file", err)
			}
			recipeTemplate.Execute(res, nil)
		}
}
