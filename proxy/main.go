package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	pb "github.com/zerobase-xyz/go-grpc-sample/pb"
	"google.golang.org/grpc"
)

var (
	empty = new(pb.Empty)
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {

	json := simplejson.New()
	json.Set("status", "OK")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func GetBackendHostname(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("grpc-server:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewHostnamePodServiceClient(conn)
	res, err := client.GetPodHostname(context.TODO(), empty)
	fmt.Printf("result:%#v \n", res.Name)
	fmt.Printf("error::%#v \n", err)

	json := simplejson.New()
	json.Set("hostname", res.Name)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func main() {
	allowedOrigins := handlers.AllowedOrigins([]string{":8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", Healthcheck)
	r.HandleFunc("/hostname", GetBackendHostname)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
