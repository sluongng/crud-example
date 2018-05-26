package model

type (
	NewUser struct {
		Name    string `json:"name" example:"Son Luong"`
		Email   string `json:"email" example:"sluongng@gmail.com"`
		Website string `json:"website" required:"false" example:"https://sluongng@gmail.com"`
	}

	User struct {
		ID      int32  `json:"id" example:"1"`
		Name    string `json:"name" example:"Son Luong"`
		Email   string `json:"email" example:"sluongng@gmail.com"`
		Website string `json:"website" example:"https://sluongng@gmail.com"`
	}
)
