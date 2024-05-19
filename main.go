package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Pragadeesh-C/go-restapi/usecases"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env load error", err)
	}

	log.Println("env file loaded")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("CONNECTION_STRING")))

	if err != nil {
		log.Fatal("Connection failed", err)
	}
	log.Println("Connection succesful")
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping failed", err)
	}
	log.Println("Ping succesful")

}

func main() {
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	empService := usecases.EmployeeService{MongoCollection: coll}

	router := http.NewServeMux()

	router.HandleFunc("GET /health", healthHandler)

	router.HandleFunc("POST /employee", empService.CreateEmployeeHandler)
	router.HandleFunc("GET /employee/{id}", empService.GetEmployeeByID)
	router.HandleFunc("GET /employee", empService.GetAllEmployee)
	router.HandleFunc("PUT /employee/{id}", empService.UpdateEmployeeByID)
	router.HandleFunc("DELETE /employee/{id}", empService.DeleteEmployeeByID)
	router.HandleFunc("DELETE /employee", empService.DeleteAllEmployee)

	log.Println("server is running on port 8080")

	http.ListenAndServe(":8080", router)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running"))
}
