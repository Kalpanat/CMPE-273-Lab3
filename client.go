package main

import (
	"strconv"
	"sort"
	"hash/fnv"
	"net/http"
	"log"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type KeyValue []struct {
	Key int `json:"key"`
	Value string `json:"value"`
}
type KeyValueSingle struct {
	Key int `json:"key"`
	Value string `json:"value"`
}

var hasingFuncMap= map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e",6:"f",7:"g",8:"h",9:"i",10:"j"}
var serverHashMap= make(map[int]string)
var sortedServerHashMap= make(map[int]string)
var hostname=[]string{"0","1","2"}
var serverNameshash []int

func main() {
	for i:=0;i<len(hostname);i++{
		serverHashMap[hash(hostname[i])]=hostname[i]
	}

 
    for s := range serverHashMap {
        serverNameshash = append(serverNameshash, s)
        //fmt.Println("Server ",s)
    }
    sort.Ints(serverNameshash)

	for _, s := range serverNameshash {
		sortedServerHashMap[s]=serverHashMap[s]
	}


	var keys []int
    for k := range hasingFuncMap {
        keys = append(keys, k)
    }
    sort.Ints(keys)

	for _, k := range keys {
		keyStr:=strconv.Itoa(k)
		serverName:=addKey(hash(keyStr))
		url :="http://localhost:300"+serverName+"/keys/"+keyStr+"/"+hasingFuncMap[k]
		client := &http.Client{}
		request, err := http.NewRequest("PUT", url, strings.NewReader("<golang>Consistent Hash map</golang>"))
		response,err:= client.Do(request)
		if err != nil {
			    log.Fatal(err)
			}
		fmt.Printf("key=%s value=%s host=%s \n",keyStr,hasingFuncMap[k],"300"+serverName)
		fmt.Println(response)
    }
    var keyval KeyValueSingle
    url :="http://localhost:3000/keys/8"
    responseValue, err := http.Get(url)
	defer responseValue.Body.Close()
	reply, err := ioutil.ReadAll(responseValue.Body)
	json.Unmarshal([]byte(reply), &keyval)
	if err != nil {
	                fmt.Printf("%s", err)
	        	}
	fmt.Println("GET Request:",url)      	
	fmt.Println("GET Response:",keyval)
}

func hash(s string) int {
        h := fnv.New32a()
        h.Write([]byte(s))
        return int(h.Sum32())
}
func addKey(object int)(serverName string){
	for i:=0;i<len(serverNameshash);i++{
		 if object<=serverNameshash[i]{
		 	return sortedServerHashMap[serverNameshash[i]]
		 }
	}
	return sortedServerHashMap[serverNameshash[1]]
}






