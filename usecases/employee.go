package usecases

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pragadeesh-C/go-restapi/models"
	"github.com/Pragadeesh-C/go-restapi/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp models.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}
	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("INsert error", err)
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)

	log.Println("Employee created1 successful", insertID, emp)
}

func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	empID := r.PathValue("id")

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployeeByID(empID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Fetch id error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindAllEmployee()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	empID := r.PathValue("id")

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid emp ID")
		res.Error = "Invalid emp ID"
		return
	}
	var emp *models.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}
	emp.EmployeeID = empID

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.UpdateEmployeeByID(empID, emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)

}
func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	empID := r.PathValue("id")

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.DeleteEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Fetch id error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)

}
func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error deleting emps", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}
