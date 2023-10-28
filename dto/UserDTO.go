package dto

type Register struct {
	Name            string `binding:"required,min=5,max=20"`
	Email           string `binding:"required,min=5,max=50,email"`
	Password        string `binding:"required,min=5,max=10"`
	PasswordConfirm string `binding:"required,min=5,max=10,eqfield=Password"`
}

type Login struct {
	Email    string `binding:"required,min=5,max=50,email"`
	Password string `binding:"required,min=5,max=10"`
}
