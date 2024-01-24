package main

import (
	"time"
)

func main() {
	patient := &Patient{ID: 1, Name: "Іван Іванов"}
	doctor := &Doctor{1, "Доктор Сміт"}

	// Призначення лікаря пацієнту
	patient.AssignDoctor(doctor)

	// Створення медичного запису
	patient.MedRecord = &MedicalRecord{}

	// Проведення базового обстеження та встановлення діагнозу
	diagnosis := &Diagnosis{time.Now(), "Грип", "Середня тяжкість"}
	patient.MedRecord.AddDiagnosis(diagnosis)

	report := &PatientReport{Patient: patient}

	// Виведення звіту
	report.PrintPatientInfo()
	report.PrintDiagnoses()
}
