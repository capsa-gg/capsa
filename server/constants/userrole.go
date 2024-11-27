package constants

import "fmt"

// UserRole is used to determine a user's role/type.
type UserRole string

const (
	// UserRoleAdmin indicates a user has admin permissions.
	UserRoleAdmin = "Admin"

	// UserRoleUser indicates that a user has no admin permissions.
	UserRoleUser = "User"
)

// AllUserRoles contains all user roles. This should only be used in cases where any user role should be allowed to access a resource.
var AllUserRoles = []UserRole{UserRoleAdmin, UserRoleUser}

// UserRoleFromString parses a string and validates if it's a valid UserRole.
func UserRoleFromString(s string) (UserRole, error) {
	if s == UserRoleAdmin {
		return UserRoleAdmin, nil
	}

	if s == UserRoleUser {
		return UserRoleUser, nil
	}

	return "", fmt.Errorf("%s is not a valid UserRole value", s)
}
