package main

import (
	"time"
)

func main() {
	patient := NewPatient(1, "Іван Іванов")
	doctor := NewDoctor(1, "Доктор Сміт")

	// Призначення лікаря пацієнту
	patient.AssignDoctor(doctor)

	// Створення медичного запису
	patient.MedRecord = NewMedicalRecord()

	// Проведення базового обстеження та встановлення діагнозу
	diagnosis := NewDiagnosis(time.Now(), "Грип", "Середня тяжкість")
	patient.MedRecord.AddDiagnosis(diagnosis)

	// Призначення лікування
	treatment := NewTreatment(time.Now(), "Лікування грипу: відпочинок, пиття багато рідини, прийом антивірусних препаратів")
	patient.MedRecord.AddTreatment(treatment)

	report := &PatientReport{Patient: patient}

	// Виведення звіту
	report.PrintPatientInfo()
	report.PrintDiagnoses()
	report.PrintTreatments()
}
