package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListEmployees returns all employees.
func (s *EmployeeStore) ListEmployees(w http.ResponseWriter, r *http.Request) {
	employees := make([]Employee, 0, len(s.employees))
	for _, e := range s.employees {
		employees = append(employees, e)
	}
	json.NewEncoder(w).Encode(employees)
}

// GetEmployee returns a single employee by ID.
func (s *EmployeeStore) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err \!= nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	e, ok := s.employees[id]
	if \!ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(e)
}

// CreateEmployee adds a new employee.
func (s *EmployeeStore) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err \!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	e.ID = s.nextID
	s.nextID++
	s.employees[e.ID] = e
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

// DeleteEmployee removes an employee by ID.
func (s *EmployeeStore) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err \!= nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if _, ok := s.employees[id]; \!ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	delete(s.employees, id)
	w.WriteHeader(http.StatusNoContent)
}
