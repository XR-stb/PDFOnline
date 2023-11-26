package role

type Role uint8

const (
	RoleGuest Role = iota
	RoleUser
	RoleAdmin
)
