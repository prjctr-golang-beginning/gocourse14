package main

import (
	"fmt"
	"time"
)

// Patient - тепер має лише базову інформацію
type Patient struct {
	ID   int
	Name string
}

// Doctor - тепер має лише базову інформацію
type Doctor struct {
	ID   int
	Name string
}

// MedicalRecord - тепер зберігає лише рядки, без структур
type MedicalRecord struct {
	History []string
}

// Функції створення об'єктів тепер приймають більше параметрів
func NewPatient(id int, name string) *Patient {
	return &Patient{
		ID:   id,
		Name: name,
	}
}

func NewDoctor(id int, name string) *Doctor {
	return &Doctor{
		ID:   id,
		Name: name,
	}
}

func NewMedicalRecord(history []string) *MedicalRecord {
	return &MedicalRecord{history}
}

// Функції додавання діагнозу та лікування тепер не пов'язані з MedicalRecord
func AddDiagnosis(history []string, diagnosis string) []string {
	return append(history, diagnosis)
}

func AddTreatment(history []string, treatment string) []string {
	return append(history, treatment)
}

func main() {
	// Ініціалізація об'єктів
	patient := NewPatient(1, "Іван Іванов")
	doctor := NewDoctor(1, "Доктор Сміт")
	medRecord := NewMedicalRecord([]string{})

	// Додавання діагнозу та лікування
	medRecord.History = AddDiagnosis(medRecord.History, fmt.Sprintf("Діагноз: Грип, Тяжкість: Середня, Дата: %s", time.Now().Format("2006-01-02")))
	medRecord.History = AddTreatment(medRecord.History, fmt.Sprintf("Лікування: Грип, Дата початку: %s", time.Now().Format("2006-01-02")))

	// Виведення інформації
	fmt.Printf("Пацієнт: %s\n", patient.Name)
	fmt.Printf("Лікар: %s\n", doctor.Name)
	for _, entry := range medRecord.History {
		fmt.Println(entry)
	}
}
