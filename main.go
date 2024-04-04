package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func addTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userIDStr := r.PathValue("id")
	if userIDStr == "" {
		http.Error(w, "User ID not found in URL", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "cannot parse ID", http.StatusInternalServerError)
		return
	}
	var exists bool
	row := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", userID)
	err = row.Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error while reading request body ", http.StatusInternalServerError)
		return
	}
	body_map := make(map[string]string)
	err = json.Unmarshal(body, &body_map)
	if err != nil {
		http.Error(w, "error while unmarshalling data from json ", http.StatusInternalServerError)
		return
	}
	tweet, ok := body_map["tweet"]
	if !ok || len(tweet) == 0 {
		http.Error(w, "Tweet not found in request body", http.StatusBadRequest)
		return
	}
	_, err = db.Exec("INSERT INTO tweets (user_id, tweet) VALUES ($1, $2)", userID, tweet)
	if err != nil {
		http.Error(w, "Failed to add tweet to database", http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Successfully added the Tweet")
}

func getTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userIDStr := r.PathValue("id")
	if userIDStr == "" {
		http.Error(w, "User ID not found in URL", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "cannot seperate index from URL", http.StatusInternalServerError)
		return
	}
	var exists bool
	row := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", userID)
	err = row.Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}
	rows, err := db.Query("SELECT tweet FROM tweets WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	tweets := make([]string, 0)
	for rows.Next() {
		var tweet string
		err := rows.Scan(&tweet)
		if err != nil {
			http.Error(w, "error while scanning tweets", http.StatusInternalServerError)
			return
		}
		tweets = append(tweets, tweet)
	}
	if len(tweets) == 0 {
		io.WriteString(w, "No tweets for this user")
		return
	}
	//convert to json
	jsonData, err := json.Marshal(tweets)
	if err != nil {
		http.Error(w, "cannot convert tweets to json", http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(jsonData))
}

func main() {
	// connect to db
	db, err = sql.Open("postgres", "postgres://postgres:123456@localhost:5432/twitter?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connected")
	// in body json of key "tweet"- eg{"tweet":"hello world"}
	http.HandleFunc("/addTweet/{id}", addTweet)
	//in case id is passed as endpoint
	http.HandleFunc("/getTweet/{id}", getTweet)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
