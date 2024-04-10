package models

const (
	RoleAdmin  UserRole = "admin"
	RoleClient UserRole = "client"
)

type UserRole string

func (r UserRole) IsAdmin() bool {
	return r == RoleAdmin
}
