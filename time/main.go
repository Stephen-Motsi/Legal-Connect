package main

import(
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
)

type Client struct{
	ID int `json:"id"`
	Name string `json:"name"`
}

type Bill struct{
	ClientID int  `json:client_id"`
	Hours float64  `json:"hours"`
	Rate  float64  `json:"rate"`
	Total  float64 `json:"total"`
	
}

func main() {
	db, err:= sql.Open("sqlite3", "./database.db")
	if err != nil{
		log.fatal(err)
	}
	defer db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/api/clients", func(w http.ResponseWriter,r *http.Request){
		rows, err := db.Query("SELECT id, name FROM clients")
		if err!=nil{
			log.Fatal()
		}
		defer rows.Close()
		 
		clients :=[]client{}
		for rows.Next(){
			var client Client
			err :=rows.Scan(&client.ID, &client.Name)
			if err!=nil{
				log.Fatal(err)
			}
			clients = append(clients, client)
		}
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(clients)

	}).Methods("GET")
	router..HandleFunc("/api/bill", func(w http.ResponseWriter, r *http.Request){
		decoder := json.NewDecoder(r.Body)
		var bill Bill
		err :=decoder.Decode(&bill)
		if err !=nil{
			log.Fatal(err)
		}

		_, err = db.Exec("INSERT INTO bills (client_id, hours, rate, total) VALUES(?, ?, ?, ?)",bill.ClientID, bill.Hours,bill.Rate,bill.Total)
		if err !=nil{
			log.Fatal(err)
		}
		
		w.WriterHeader(http.StatusOK)
	})
	.Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}