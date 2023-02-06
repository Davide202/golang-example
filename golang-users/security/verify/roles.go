package security

const (
	Admin   Role = "ADMIN"
	User    Role = "USER"
	Viewer  Role = "VIEWER"
	UNKNOWN Role = "UNKNOWN"

	ROLE = "ROLE"
)

type RoleInterf interface {
	fromString(string)
}

type Role string

func (r *Role) fromString(str string) {
	role := getRole(str)
	r = &role
}

type Roles []Role

func GetRoles(str string) map[string]Roles {
	role := getRole(str)
	var roles []Role
	roles = append(roles, role)
	return map[string]Roles{
		ROLE: roles,
	}
}

func getRole(str string) Role {
	switch str {
	case string(Admin):
		return Admin
	case string(User):
		return User
	case string(Viewer):
		return Viewer
	}
	return UNKNOWN
}

func AdminRole() map[string]string {
	return map[string]string{
		ROLE: string(Admin),
	}
}

func UserRole() map[string]string {
	return map[string]string{
		ROLE: string(User),
	}
}

func ViewerRole() map[string]string {
	return map[string]string{
		ROLE: string(Viewer),
	}
}
