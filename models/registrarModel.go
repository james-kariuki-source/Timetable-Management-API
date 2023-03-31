package models

type Registrar struct{
	EmployeeID uint `json:"id"`
	Fname string `json:"first_name"`
	Lname string `json:"last_name"`
	Password []byte `json:"password"`
}