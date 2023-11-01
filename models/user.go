package models

import (
	"github.com/Shakezidin/configs"
)
func RegisterUser(email string) error {
	return configs.Client.Set("user",email,0).Err()
}
func AuthenticateUser(email string) error {
	return configs.Client.Get("user").Err()
}