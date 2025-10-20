package auth

// User represents a system user
type User struct {
	Name     string
	IsAdmin  bool
	LoggedIn bool
}

// NewUser creates a new user with the given parameters
func NewUser(name string, isAdmin bool, loggedIn bool) *User {
	return &User{
		Name:     name,
		IsAdmin:  isAdmin,
		LoggedIn: loggedIn,
	}
}

// CanAccessGrades checks if the user can access grade information
func (u *User) CanAccessGrades() bool {
	return u.LoggedIn
}

// CanAccessAdminFeatures checks if user can access admin features
func (u *User) CanAccessAdminFeatures() bool {
	return u.LoggedIn && u.IsAdmin
}

// GetGreeting returns a personalized greeting for the user
func (u *User) GetGreeting() string {
	return "Hello, " + u.Name
}

// GetStatusMessage returns the user's login status
func (u *User) GetStatusMessage() string {
	if !u.LoggedIn {
		return "The Request is Denied you should log in first"
	}

	if u.IsAdmin {
		return "User is logged in as an Admin"
	}

	return "User is logged in as a Regular User"
}
