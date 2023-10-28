package web

type MedicalIDRequest struct {
	Birthdate        string `json:"birthdate" form:"birthdate"`
	Gender           string `json:"gender" form:"gender"`
	BloodType        string `json:"blood_type" form:"blood_type"`
	Height           int    `json:"height" form:"height"`
	Weight           int    `json:"weight" form:"weight"`
	MedicalCondition string `json:"medical_condition" form:"medical_condition"`
	EmergencyContact string `json:"emergency_contact" form:"emergency_contact"`
}
