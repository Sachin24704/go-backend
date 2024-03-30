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
	 	http.Error(w , "cannot seperate index from URL", 400)
	 	return
    }

	body,err:=io.ReadAll(r.Body)
	if err!=nil {
		http.Error(w, "error while reading request body ", 400)
		return  
	}
	body_map:=make(map[string]string)
	json.Unmarshal(body,&body_map)

	tweet:=body_map["tweet"]
	if len(tweet) == 0{
		http.Error(w,"tweet invalid",400)
		return
	}

	for i,u:= range user{
		if u.Id == userID{
			user[i].Tweets=append(user[i].Tweets, tweet)
			jsonData,err :=json.Marshal(user[i].Tweets)
			if err!=nil {
				http.Error(w , "cannot convert tweets to json", 400)
		        return

			}
			io.WriteString(w, string(jsonData))
			return
		}
	}
	http.Error(w,"User not found",400)
	
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
	 	http.Error(w , "cannot seperate index from URL", 400)
	 	return
    }

	// when id is passed in body using "id"

	/*body, err := io.ReadAll(r.Body)
	if err!=nil {
		http.Error(w , "error in reading r.body", http.StatusMethodNotAllowed)
		return

	}
	body_map:=make(map[string]string)

	json.Unmarshal(body, &body_map)
    
	index, err := strconv.Atoi(body_map["id"])
	if err!=nil {
		http.Error(w , "could not find id from body", http.StatusMethodNotAllowed)
		return

	} */
	
	for _,u := range user {
		if u.Id ==userID{
			jsonData, err :=json.Marshal(u.Tweets)
			if err!=nil {
				http.Error(w , "cannot convert tweets to json", 400)
		        return

			}
			io.WriteString(w, string(jsonData))
			return

		}

	}
	http.Error(w,"User not found",400)
}


func main(){

  // in body json of key "tweet"- eg{"tweet":"hello world"} 
  http.HandleFunc("/addTweet/{id}",addTweet)

  // in body json of key "id"- eg {"id":"1" }
  //http.HandleFunc("/getTweet",getTweet)

  //in case id is passed as endpoint
  http.HandleFunc("/getTweet/{id}",getTweet)
  
  log.Fatal(http.ListenAndServe(":8000", nil))
}
