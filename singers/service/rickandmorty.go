package service

import(
	"net/http"
	"io"
	//"io/ioutil"
	"encoding/json"
	//"fmt"
	"log"
)

type Info struct{
	Characters string //`json:"characters"`
	Locations string //`json:"locations"`
	Episodes string //`json:"episodes"`
}
const baseUrl string = "https://rickandmortyapi.com/api"

//func GetRickandmortyParams(params map[string]string)(*Info,error){


func GetRickandmorty()(*Info,error){
	//fmt.Println("Called RockAndMortyInfo")
	log.Println("Called RockAndMortyInfo")
	response , err := http.Get(baseUrl)
	if err != nil {
		log.Println(err.Error())
		return nil , err
	}
	if response != nil {
		log.Println("Reponse: " + response.Status)
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
			return nil , err
		}
		responseObject := Info{}
		errr := json.Unmarshal(responseData,&responseObject)
		if errr != nil {
			log.Println(errr.Error())
			return nil, errr
		}
		return &responseObject , nil
	}
	log.Println("Response == null")
	return nil , nil
}