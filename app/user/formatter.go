package user

import "golang-blueprint/app/models"

type UserFormatter struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

func FormatUser(user models.User) UserFormatter {
	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return formatter
}
