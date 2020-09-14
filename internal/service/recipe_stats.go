package service

import (
	"encoding/json"
	"fmt"
	"hello-fresh-app/internal/model"
	"hello-fresh-app/internal/service/worker"
	"io"
	"sync"
)

type Recipe struct {}

const (
	deliveryStartTime = "10AM"
	deliveryEndTime   = "03PM"
	defautPostCode    = "10120"
)

type RecipeStats interface {
	CalculateStats(file io.Reader, postCode string, recipeName string) model.Response
	FindStatsByRecipeName(file io.Reader,recipeName string) model.RecipeType
	FindStatsByPostCodeAndTime(file io.Reader, postCode string, startTime string, endTime string)
}

func (r Recipe) CalculateStats(file io.Reader) model.Response {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan model.RecipeStats)
	result := make(chan model.Response)
	go readRecipesData(file, &wg, ch)
	go worker.Start(ch, result, defautPostCode, deliveryStartTime, deliveryEndTime, "")
	go func() {
		wg.Wait()
		close(ch)
	}()
	return <- result
}


func (r Recipe) FindStatsByRecipeName(file io.Reader, recipeName string) model.RecipeType {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan model.RecipeStats)
	go readRecipesData(file, &wg, ch)
	response := worker.StatByRecipeName(ch, recipeName)
	return response
}

func (r Recipe) FindStatsByPostCodeAndTime(file io.Reader, postCode string, startTime string, endTime string) model.Response {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan model.RecipeStats)
	result := make(chan model.Response)
	go readRecipesData(file, &wg, ch)
	go worker.Start(ch, result, postCode, startTime, endTime, "")
	return <- result
}

func readRecipesData(file io.Reader, wg *sync.WaitGroup, ch chan model.RecipeStats )  {
	defer wg.Done()
	decoder := json.NewDecoder(file)
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Empty")
	}
	fmt.Println("Token: ", token)
	for decoder.More() {
		var recipeStats model.RecipeStats
		err := decoder.Decode(&recipeStats)
		if err != nil {
			fmt.Println(err)
		}
		ch <- recipeStats
	}
}