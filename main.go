package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id     int  `json:"id"`
	Name   string  `json:"name"`
	Tweets []string  `json:"tweets"`
}

var user = []User{
	{1, "sachin", []string{"my name is sachin id 1"}}, 
	{2, "dev", []string{"my name is dev id 2"}}, 
	{3, "rahul", []string{"my name is rahul id 3"}}, 
	{4, "vinay", []string{"my name is vinay id 4"}}, 
	{5, "ankit", []string{"my name is ankit id 5"}},
}


func addTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST"{
		http.Error(w , "Method not allowed", http.StatusMethodNotAllowed)
		return 
	}
    userIDStr := r.PathValue("id")
	if userIDStr == "" {
		http.Error(w, "User ID not found in URL", http.StatusBadRequest)
		return
	}
    userID,err:=strconv.Atoi(userIDStr)
    if  err!= nil {
	 	http.Error(w , "cannot seperate index from URL", http.StatusInternalServerError)
	 	return
    }
    body,err:=io.ReadAll(r.Body)
	if err!=nil {
		http.Error(w, "error while reading request body ", http.StatusInternalServerError)
		return  
	}
	body_map:=make(map[string]string)
	json.Unmarshal(body,&body_map)
    tweet:=body_map["tweet"]
	if len(tweet) == 0{
		http.Error(w,"tweet invalid",http.StatusBadRequest)
		return
	}
    for i,u:= range user{
		if u.Id == userID{
			user[i].Tweets=append(user[i].Tweets, tweet)
			jsonData,err :=json.Marshal(user[i].Tweets)
			if err!=nil {
				http.Error(w , "cannot convert tweets to json", http.StatusInternalServerError)
		        return
            }
			io.WriteString(w, string(jsonData))
			return
		}
	}
	http.Error(w,"User not found",http.StatusNotFound)
}

func getTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET"{
		http.Error(w , "Method not allowed", http.StatusMethodNotAllowed)
		return 
	}
	userIDStr := r.PathValue("id")
	if userIDStr == "" {
		http.Error(w, "User ID not found in URL", http.StatusBadRequest)
		return
	}
    userID,err:=strconv.Atoi(userIDStr)
    if  err!= nil {
	 	http.Error(w , "cannot seperate index from URL", http.StatusInternalServerError)
	 	return
    }
	for _,u := range user {
		if u.Id ==userID{
			jsonData, err :=json.Marshal(u.Tweets)
			if err!=nil {
				http.Error(w , "cannot convert tweets to json", http.StatusInternalServerError)
		        return
			}
			io.WriteString(w, string(jsonData))
			return
		}
	}
	http.Error(w,"User not found",http.StatusNotFound)
}


func main(){
  // in body json of key "tweet"- eg{"tweet":"hello world"} 
  http.HandleFunc("/addTweet/{id}",addTweet)
 //in case id is passed as endpoint
  http.HandleFunc("/getTweet/{id}",getTweet)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
