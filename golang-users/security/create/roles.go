package create

const (
	Admin   Role = "ADMIN"
	User    Role = "USER"
	Viewer  Role = "VIEWER"
	UNKNOWN Role = "UNKNOWN"

	//ROLE = "ROLE"
)

type RoleInterf interface {
	fromString(string) *Role
}

type Role string

func (r Role) String() string {
	switch r {
	case Admin:
		return "ADMIN"
	case User:
		return "USER"
	case Viewer:
		return "VIEWER"
	}
	return "UNKNOWN"
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
