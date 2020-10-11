package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type tour struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"product_name"`
	Details  string `json:"details"`
	Duration string `json:"duration"`
	Cost     string `json:"cost"`
	Img      string `json:"img"`
}

var db *sql.DB

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	var err error
	db, err = sql.Open("mysql", "root:Supersecret1@tcp(Database:3306)/fernweh")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	r := mux.NewRouter()
	// route handler/ endpoint

	r.HandleFunc("/frontend/solo.html", fetchTours).Methods("GET")
	r.HandleFunc("/frontend/family.html", fetchTours).Methods("GET")
	r.HandleFunc("/frontend/yolo.html", fetchTours).Methods("GET")
	r.HandleFunc("/frontend/culture.html", fetchTours).Methods("GET")

	log.Println("listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func fetchTours(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var tours []tour
	stmt := "SELECT * FROM tours"
	rows, err := db.Query(stmt)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var tour tour
		err := rows.Scan(&tour.ID, &tour.Category, &tour.Name, &tour.Details, &tour.Duration, &tour.Cost, &tour.Img)
		if err != nil {
			panic(err.Error())
		}

		tours = append(tours, tour)
	}

	//WHAT IS THIS?
	w.WriteHeader(http.StatusOK)
	//prepares browser for JSON text
	w.Header().Set("Content-Type", "application/json")
	//converts tour slice into json
	json.NewEncoder(w).Encode(tours)
}
