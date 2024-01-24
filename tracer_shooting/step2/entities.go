package main

import (
	"fmt"
	"time"
)

// Patient представляє клієнта клініки
type Patient struct {
	ID        int
	Name      string
	Doctor    *Doctor
	MedRecord *MedicalRecord
}

// Doctor представляє лікаря
type Doctor struct {
	ID   int
	Name string
}

// MedicalRecord представляє медичний запис пацієнта
type MedicalRecord struct {
	History   []string
	Diagnoses []*Diagnosis
}

// Diagnosis представляє діагноз пацієнта
type Diagnosis struct {
	Date     time.Time
	Disease  string
	Severity string
}

// AssignDoctor призначає лікаря пацієнту
func (p *Patient) AssignDoctor(d *Doctor) {
	p.Doctor = d
}

// AddDiagnosis додає діагноз до медичного запису
func (mr *MedicalRecord) AddDiagnosis(d *Diagnosis) {
	mr.Diagnoses = append(mr.Diagnoses, d)
}

type PatientReport struct {
	Patient *Patient
}

// PrintPatientInfo виводить інформацію про пацієнта
func (pr *PatientReport) PrintPatientInfo() {
	fmt.Printf("Пацієнт: %s\n", pr.Patient.Name)
	fmt.Printf("Лікар: %s\n", pr.Patient.Doctor.Name)
}

// PrintDiagnoses виводить інформацію про діагнози пацієнта
func (pr *PatientReport) PrintDiagnoses() {
	for _, d := range pr.Patient.MedRecord.Diagnoses {
		fmt.Printf("Діагноз: %s, Тяжкість: %s, Дата: %s\n", d.Disease, d.Severity, d.Date.Format("2006-01-02"))
	}
}
