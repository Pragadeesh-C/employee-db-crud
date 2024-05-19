package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Pragadeesh-C/go-restapi/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(os.Getenv("CONNECTION_STRING")))

	if err != nil {
		log.Fatal("Error while connecting to Mongo", err)
	}

	log.Println("Mongodb connected succesfully")

	err = mongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Ping failed", err)
	}

	log.Println("Ping success")

	return mongoClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	emp1 := uuid.New().String()

	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}

	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := models.Employee{
			Name:       "TestUser",
			Department: "CSE",
			EmployeeID: emp1,
		}
		result, err := empRepo.InsertEmployee((&emp))

		if err != nil {
			t.Fatal("Insert 1 operation failed", err)
		}

		t.Log("Insert 1 successful", result)
	})

	t.Run("Get Employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			t.Fatal("Error", err)
		}

		t.Log("Emp 1", result.Name)
	})

	t.Run("Get All Employees", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()

		if err != nil {
			t.Fatal("Error get operation failed", err)
		}

		t.Log("Employees", results)
	})

	t.Run("Update Employee 1 Name", func(t *testing.T) {
		emp := models.Employee{
			Name:       "Test",
			Department: "CSE",
			EmployeeID: emp1,
		}

		result, err := empRepo.UpdateEmployeeByID(emp1, &emp)

		if err != nil {
			log.Fatal("Update operation faield", err)
		}

		t.Log("emp 1", result)
	})

	t.Run("Get Employee 1 after update", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			t.Fatal("Get operation failed", err)
		}

		t.Log("Emp 1 ", result.Name)
	})

	t.Run("Delete Employee 1", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)

		if err != nil {
			t.Fatal("Delete operation failed", err)
		}

		t.Log("Deleted successfully", result)
	})

	t.Run("Get All Employees after delete", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()

		if err != nil {
			t.Fatal("Error get operation failed", err)
		}

		t.Log("Employees", results)
	})

	t.Run("Delete all Employees", func(t *testing.T) {
		results, err := empRepo.DeleteAllEmployee()

		if err != nil {
			t.Fatal("Delete operation failed", err)
		}

		t.Log("Delete oepration success", results)
	})

}
