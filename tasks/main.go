package main

import (
	"fmt"
	"gocourse14/tasks/doctor"
	"gocourse14/tasks/patient"
)

func main() {
	doctor.Init()
	patient.Init()

	// Реєстрація нового лікаря
	newDoctorID := doctor.DoctorID("Doctor1")
	newDoctorSpecialization := doctor.Specialization("Neurologist")
	doctor.RegisterDoctor(newDoctorID, newDoctorSpecialization)

	// Прив'язка пацієнтів до нового лікаря
	patient.AssignPatient(newDoctorID, `Alex`)
	patient.AssignPatient(newDoctorID, `Eugen`)

	// Виведення інформації
	fmt.Println("Doctor's Specialization:", doctor.GetSpecialization(newDoctorID))
	fmt.Println("Assigned Patients:", patient.GetPatients(newDoctorID))
}
