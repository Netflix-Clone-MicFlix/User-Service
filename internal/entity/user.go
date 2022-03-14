// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type User struct {
	Username string `json:"username"                   example:"poeky"`
	Password string `json:"password"             example:"dsafsd111fa"`
	Email    string `json:"email"     example:"example@hotmail.com"`
}
