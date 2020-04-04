package main

import (
	"GatewayMsg/Server"
	"Server"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	StartServer()
}

func StartServer() {

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/gateway", createGateway)
	http.HandleFunc("/getGateway/{id}", getGateway)
	http.HandleFunc("/route", Route)
	http.HandleFunc("/search/route/{number,}", routeSearch)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func createGateway(w http.ResponseWriter, r *http.Request) {

	var gateway Server.CreateGatewayStruct

	err := json.NewDecoder(r.Body).Decode(&gateway)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in Request")
		log.Println(err)
	}

	log.Println(gateway.Name, gateway.IP_addresses)

	if creatingGateway(gateway) == false {
		fmt.Fprint(w, "Gateway adding not successfull")
		return
	} else {
		fmt.Fprint(w, "Gateway added")
		return
	}
}

func creatingGateway(gateway Server.CreateGatewayStruct) bool {

	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, host_port, username, password, database_name)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Println(err)
		return false
	} else {
		log.Println("Database connected")
	}
	defer db.Close()

	var insertGateway Server.GatewayStruct

	if gateway.Name == "Airtel" || gateway.Name == "Vodafone" || gateway.Name == "Tata" {

		if gateway.Name == "Airtel" {
			insertGateway.Prefix = "123"
		}
		if gateway.Name == "Vodafone" {
			insertGateway.Prefix = "1234"
		}
		if gateway.Name == "Tata" {
			insertGateway.Prefix = "9194"
		}
		insertGateway.Ip_addresses = gateway.IP_addresses
		insertGateway.Name = gateway.Name

		statement := `INSERT into public.gateway_table (name, ip_addresses, prefix) VALUES($1, $2, $3) RETURNING name`
		name := ""
		// size := len(insertGateway.ip_addresses)
		// var ipaddress [size]string
		err = db.QueryRow(statement, insertGateway.Name, strings.Join(insertGateway.Ip_addresses, ","), insertGateway.Prefix).Scan(&name)
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println("New record is : ", name)

	} else {
		return false
	}
	return true
}

const (
	hostname      = "localhost"
	host_port     = 5432
	username      = "postgres"
	password      = "root"
	database_name = "hackerrank_gateway"
)

func getGateway(w http.ResponseWriter, r *http.Request) {

	id, ok := r.URL.Query()["id"]
	if ok == false {
		log.Println(ok)
	}

	if gettinggateway(strings.Join(id, ",")) == false {
		w.WriteHeader(404)
		return
	} else {
		w.WriteHeader(200)
		return
	}

}
func gettinggateway(id string) bool {

	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, host_port, username, password, database_name)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Println(err)
		return false
	} else {
		log.Println("Database connected")
	}
	defer db.Close()

	statement := `SELECT id FROM public.gateway_table WHERE id = $1`
	_, err = db.Exec(statement, id)
	if err != nil {
		return false
	}

	return true
}
func Route(w http.ResponseWriter, r *http.Request) {

	var G Server.GatewayRouteStruct

	err := json.NewDecoder(r.Body).Decode(&G)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in Request")
		log.Println(err)
	}

	log.Println(G.Gateway_id, G.Prefix)

}

func routeSearch(w http.ResponseWriter, r *http.Request) {

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Hello")
}
