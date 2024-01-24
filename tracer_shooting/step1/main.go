package main

func main() {
	patient := &Patient{ID: 1, Name: "Іван Іванов"}
	doctor := &Doctor{1, "Доктор Сміт"}

	patient.AssignDoctor(doctor)

	report := &PatientReport{Patient: patient}

	// Виведення звіту
	report.PrintPatientInfo()
}
