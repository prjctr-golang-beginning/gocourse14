package doctor

type DoctorID string
type Specialization string

var doctorsSpecializations map[DoctorID]Specialization

func Init() {
	doctorsSpecializations = make(map[DoctorID]Specialization)
}

func RegisterDoctor(id DoctorID, specialization Specialization) {
	doctorsSpecializations[id] = specialization
}

func GetSpecialization(id DoctorID) Specialization {
	return doctorsSpecializations[id]
}
