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

// NewPatient створює нового пацієнта
func NewPatient(id int, name string) *Patient {
	return &Patient{
		ID:   id,
		Name: name,
	}
}

// Doctor представляє лікаря
type Doctor struct {
	ID   int
	Name string
}

// NewDoctor створює нового лікаря
func NewDoctor(id int, name string) *Doctor {
	return &Doctor{
		ID:   id,
		Name: name,
	}
}

// MedicalRecord представляє медичний запис пацієнта
type MedicalRecord struct {
	History    []string
	Diagnoses  []*Diagnosis
	Treatments []*Treatment
}

// NewMedicalRecord створює новий медичний запис
func NewMedicalRecord() *MedicalRecord {
	return &MedicalRecord{}
}

// Diagnosis представляє діагноз пацієнта
type Diagnosis struct {
	Date     time.Time
	Disease  string
	Severity string
}

// NewDiagnosis створює новий діагноз
func NewDiagnosis(date time.Time, disease, severity string) *Diagnosis {
	return &Diagnosis{
		Date:     date,
		Disease:  disease,
		Severity: severity,
	}
}

// Treatment представляє план лікування пацієнта
type Treatment struct {
	DateStarted   time.Time
	TreatmentPlan string
}

// NewTreatment створює нове лікування
func NewTreatment(dateStarted time.Time, plan string) *Treatment {
	return &Treatment{
		DateStarted:   dateStarted,
		TreatmentPlan: plan,
	}
}

// AssignDoctor призначає лікаря пацієнту
func (p *Patient) AssignDoctor(d *Doctor) {
	p.Doctor = d
}

// AddDiagnosis додає діагноз до медичного запису
func (mr *MedicalRecord) AddDiagnosis(d *Diagnosis) {
	mr.Diagnoses = append(mr.Diagnoses, d)
}

// AddTreatment додає лікування до медичного запису
func (mr *MedicalRecord) AddTreatment(t *Treatment) {
	mr.Treatments = append(mr.Treatments, t)
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

// PrintTreatments виводить інформацію про лікування пацієнта
func (pr *PatientReport) PrintTreatments() {
	for _, t := range pr.Patient.MedRecord.Treatments {
		fmt.Printf("Лікування: %s, Дата початку: %s\n", t.TreatmentPlan, t.DateStarted.Format("2006-01-02"))
	}
}
