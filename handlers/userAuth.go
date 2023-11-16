package handlers

type Login struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type SignUp struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}
