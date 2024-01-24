package main

import "fmt"

// Patient представляє клієнта клініки
type Patient struct {
	ID     int
	Name   string
	Doctor *Doctor
}

// AssignDoctor призначає лікаря пацієнту
func (p *Patient) AssignDoctor(d *Doctor) {
	p.Doctor = d
}

// Doctor представляє лікаря
type Doctor struct {
	ID   int
	Name string
}

type PatientReport struct {
	Patient *Patient
}

// PrintPatientInfo виводить інформацію про пацієнта
func (pr *PatientReport) PrintPatientInfo() {
	fmt.Printf("Пацієнт: %s\n", pr.Patient.Name)
	fmt.Printf("Лікар: %s\n", pr.Patient.Doctor.Name)
}
