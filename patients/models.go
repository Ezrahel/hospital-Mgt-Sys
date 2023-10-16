package patients

type patientsDetails struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phoneno"`
	Sickness string `json:"sickness"`
}
