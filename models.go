package main

// Employee represents an employee record.
type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

// EmployeeStore is an in-memory store for employees.
type EmployeeStore struct {
	employees map[int]Employee
	nextID    int
}

// NewEmployeeStore creates a new EmployeeStore.
func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: make(map[int]Employee),
		nextID:    1,
	}
}
