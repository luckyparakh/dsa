package main

import "fmt"

// User struct
type User struct {
	Username string
	Password string
}

// UserRepository struct for user management
type UserRepository struct {
	Users []User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{Users: make([]User, 0)}
}

func (u *UserRepository) AddUser(user User) {
	u.Users = append(u.Users, user)
}

type Authenticator interface {
	Authenticate(user User, credential any) bool
}
type PasswordAuthenticator struct {
}

func (p PasswordAuthenticator) Authenticate(user User, credential any) bool {
	password, ok := credential.(string)
	if !ok {
		return false
	}
	return user.Password == password
}

// AuthenticationService represents a service for authenticating users.
type AuthenticationService struct {
	userRepository *UserRepository
	authenticator  Authenticator
}

// NewAuthenticationService creates a new instance of AuthenticationService.
// It takes a userRepository as a parameter and returns a pointer to AuthenticationService.
func NewAuthenticationService(userRepository *UserRepository, auth Authenticator) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
		authenticator:  auth,
	}
}

func (as *AuthenticationService) SetAuthenticator(auth Authenticator) {
	as.authenticator = auth
}

func (as *AuthenticationService) getUserByName(userName string) (User, error) {
	for _, user := range as.userRepository.Users {
		if user.Username == userName {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User Not Found")
}

func (as *AuthenticationService) Authenticate(userName, password string) (bool, error) {
	user, err := as.getUserByName(userName)
	if err != nil {
		return false, err
	}
	return as.authenticator.Authenticate(user, password), nil
}

func main() {
	userRepository := NewUserRepository()
	userRepository.AddUser(User{Username: "user1", Password: "password1"})
	userRepository.AddUser(User{Username: "user2", Password: "password2"})
	p:=PasswordAuthenticator{}
	authService := NewAuthenticationService(userRepository, PasswordAuthenticator{})
	isAuthenticated, err := authService.Authenticate("user1", "password11")
	if err != nil {
		fmt.Println(err)
		return
	}
	if isAuthenticated {
		fmt.Println("User is authenticated")
	} else {
		fmt.Println("User is not authenticated")
	}

}
