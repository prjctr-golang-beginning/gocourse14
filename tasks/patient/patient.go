package patient

import "gocourse14/tasks/doctor"

type PatientID string

var doctorsPatients map[doctor.DoctorID][]PatientID

func Init() {
	doctorsPatients = make(map[doctor.DoctorID][]PatientID)
}

func AssignPatient(dID doctor.DoctorID, pId PatientID) {
	if doctorsPatients[dID] == nil {
		doctorsPatients[dID] = []PatientID{pId}
	} else {
		doctorsPatients[dID] = append(doctorsPatients[dID], pId)
	}
}

func GetPatients(dID doctor.DoctorID) []PatientID {
	return doctorsPatients[dID]
}
