package model

type (
	NewUser struct {
		Name    string `json:"name" db:"username"`
		Email   string `json:"email" db:"email"`
		Website string `json:"website" db:"website" required:"false"`
	}

	User struct {
		ID int32 `json:"id" db:"id"`
		NewUser
	}
)
