package services

import (
	"bytes"
	"io"
	"net/http"

	//"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

type Info struct {
	Characters string //`json:"characters"`
	Locations  string //`json:"locations"`
	Episodes   string //`json:"episodes"`
}

const baseUrl string = "https://rickandmortyapi.com/api"

//func GetRickandmortyParams(params map[string]string)(*Info,error){

func GetRickandmorty() (*Info, error) {
	//fmt.Println("Called RockAndMortyInfo")
	log.Println("Called RockAndMortyInfo")
	response, err := http.Get(baseUrl)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if response != nil {
		log.Println("Reponse: " + response.Status)
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		responseObject := Info{}
		errr := json.Unmarshal(responseData, &responseObject) // UN-MARSHAL
		if errr != nil {
			log.Println(errr.Error())
			return nil, errr
		}
		return &responseObject, nil
	}
	log.Println("Response == null")
	return nil, nil
}

func PostToMock(array *[]string) (*map[string]any, error) {
	url := "http://localhost:8081/mock/post"
	p := *array
	values := map[string]string{"key1": p[0], "key2": p[1]}
	json_data, err := json.Marshal(values) // MARSHAL

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["key1"])
	return &res, nil
}
