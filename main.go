package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lib/pq"
)

var db  *sql.DB

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

var userIDs= make([]int,0)

func getUserID(){
	rows,err:=db.Query("SELECT user_id FROM users")
	if err !=nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var uid int
		err:=rows.Scan( &uid )
		if err !=nil {
			log.Fatal(err)
		}
		userIDs = append(userIDs, uid)
	}
}

// to check if userid is in table

func  IsUserExist(id int) bool {
	getUserID()
	for _,v := range userIDs {
		if v==id{
			return true
		}
	}
	return false
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

	// to check if user exist in DB

	userID_bool:=IsUserExist(userID)
	if !userID_bool {
		http.Error(w , "userID doesnot exist in table", http.StatusInternalServerError)
	 	return
	}

	body,err:=io.ReadAll(r.Body)
	if err!=nil {
		http.Error(w, "error while reading request body ", http.StatusInternalServerError)
		return  
	}
	body_map:=make(map[string]string)
	json.Unmarshal(body,&body_map)

	tweet, ok := body_map["tweet"]
	if !ok || len(tweet) ==0 {
		http.Error(w, "Tweet not found in request body", http.StatusInternalServerError)
		return
	}
	

	_, err = db.Exec("INSERT INTO tweets (user_id, tweet) VALUES ($1, $2)", userID, tweet)
	if err != nil {
		http.Error(w, "Failed to add tweet to database", http.StatusInternalServerError)
		return
	}
	io.WriteString(w,"Successfully added the Tweet")

	/*  in case db is not defined

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
	http.Error(w,"User not found",http.StatusInternalServerError)
	*/
	
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
// check if user exist in db
	userID_bool:=IsUserExist(userID)
	if !userID_bool {
		http.Error(w , "userID doesnot exist in table", http.StatusInternalServerError)
	 	return
	}

	rows,err:=db.Query("SELECT tweet FROM tweets WHERE user_id = $1",userID)
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	defer rows.Close()
	tweets := make([]string,0)
	for rows.Next() {
		var tweet string
		err:=rows.Scan( &tweet )
		if err!=nil {
			http.Error(w , "error while scanning tweets", http.StatusInternalServerError)
	 	    return

		}
		tweets=append(tweets,tweet)
	}
	//convert to json
	jsonData,err := json.Marshal(tweets)
	if err!=nil {
		http.Error(w , "cannot convert tweets to json", http.StatusInternalServerError)
		return

	}
	io.WriteString(w,string(jsonData))
	

	

	
	/*for _,u := range user {
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
	
	http.Error(w,"User not found",http.StatusInternalServerError)
	*/
}


func main(){
	// connect to db
	connection_url,err:=pq.ParseURL("postgres://postgres:123456@localhost:5432/twitter?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", connection_url)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	// status will check if db is connected
	status:=true
	err=db.Ping()
	if err != nil{
		log.Println(err)
		status=false
	}
	log.Println("DB connected status:",status)
  

  // in body json of key "tweet"- eg{"tweet":"hello world"} 
    http.HandleFunc("/addTweet/{id}",addTweet)

 

  //in case id is passed as endpoint
    http.HandleFunc("/getTweet/{id}",getTweet)
  
    log.Fatal(http.ListenAndServe(":8000", nil))
}
