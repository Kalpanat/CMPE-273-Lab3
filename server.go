package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
  	"strconv"
)
type KeyValue struct {
	Key int `json:"key"`
	Value string `json:"value"`
}
type AllKeyValue struct{
	allKeys []KeyValue
}
var server3000map= make(map[int]string)
var server3001map= make(map[int]string)
var server3002map= make(map[int]string)

func main() {

		finish := make(chan bool)
		router1 := mux.NewRouter()
		router1.HandleFunc("/keys", handleGetAllKey1).Methods("GET")
		router1.HandleFunc("/keys/{key}", handleGetKey1).Methods("GET")
		router1.HandleFunc("/keys/{key}/{value}", handlePutKey1).Methods("PUT")
		go func() {
			http.ListenAndServe(":3000", router1)
		}()

		
		router2 := mux.NewRouter()
		router2.HandleFunc("/keys", handleGetAllKey2).Methods("GET")
		router2.HandleFunc("/keys/{key}", handleGetKey2).Methods("GET")
		router2.HandleFunc("/keys/{key}/{value}", handlePutKey2).Methods("PUT")
		go func() {
			http.ListenAndServe(":3001", router2)
		}()


		router3 := mux.NewRouter()
		router3.HandleFunc("/keys", handleGetAllKey3).Methods("GET")
		router3.HandleFunc("/keys/{key}", handleGetKey3).Methods("GET")
		router3.HandleFunc("/keys/{key}/{value}", handlePutKey3).Methods("PUT")
		go func() {
		http.ListenAndServe(":3002", router3)
		}()

		<-finish
}
func handleGetAllKey1(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		var jsonStr string
		jsonStr+=`[`
		for key, value := range server3000map {
			jsonStr+= `{"key":`+strconv.Itoa(key)+`,"value:":"`+value+`"},`
    	}
    	jsonStr=jsonStr[0:len(jsonStr)-1]
    	jsonStr+=`]`
    	res.WriteHeader(http.StatusOK)
		fmt.Fprint(res,jsonStr)
	}

func handleGetKey1(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		key := vars["key"]
		keyNumberVal, _ := strconv.Atoi(key)
		res2D := &KeyValue{
			Key: keyNumberVal,
			Value: server3000map[keyNumberVal]}
		res2B, _ := json.Marshal(res2D)
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res,string(res2B))
	}

func handlePutKey1(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		key := vars["key"]
		value := vars["value"]
		keyNumberVal, _ := strconv.Atoi(key)
		server3000map[keyNumberVal]=value
		res.WriteHeader(http.StatusCreated)	
	}


func handleGetAllKey2(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		var jsonStr string
		jsonStr+=`[`
		for key, value := range server3001map {
			jsonStr+= `{"key":`+strconv.Itoa(key)+`,"value:":"`+value+`"},`
    	}
    	jsonStr=jsonStr[0:len(jsonStr)-1]
    	jsonStr+=`]`
    	res.WriteHeader(http.StatusOK)
		fmt.Fprint(res,jsonStr)
	}

func handleGetKey2(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		key := vars["key"]
		keyNumberVal, _ := strconv.Atoi(key)
		res2D := &KeyValue{
			Key: keyNumberVal,
			Value: server3001map[keyNumberVal]}
		res2B, _ := json.Marshal(res2D)
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res,string(res2B))
	}

func handlePutKey2(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		key := vars["key"]
		value := vars["value"]
		keyNumberVal, _ := strconv.Atoi(key)
		server3001map[keyNumberVal]=value
		res.WriteHeader(http.StatusCreated)	
	}

func handleGetAllKey3(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		var jsonStr string
		jsonStr+=`[`
		for key, value := range server3002map {
			jsonStr+= `{"key":`+strconv.Itoa(key)+`,"value:":"`+value+`"},`
    	}
    	jsonStr=jsonStr[0:len(jsonStr)-1]
    	jsonStr+=`]`
    	res.WriteHeader(http.StatusOK)
		fmt.Fprint(res,jsonStr)
	}

func handleGetKey3(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		key := vars["key"]
		keyNumberVal, _ := strconv.Atoi(key)
		res2D := &KeyValue{
			Key: keyNumberVal,
			Value: server3002map[keyNumberVal]}
		res2B, _ := json.Marshal(res2D)
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res,string(res2B))
	}
func handlePutKey3(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		key := vars["key"]
		value := vars["value"]
		keyNumberVal, _ := strconv.Atoi(key)
		server3002map[keyNumberVal]=value
		res.WriteHeader(http.StatusCreated)	
	}
