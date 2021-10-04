package models

import "strings"

const (
	// ErrNotFound is returned when the resource is not found in the database
	ErrNotFound modelError = "models: resource not found"
	// ErrIDInvalid is returned when an invalid ID is provided to a method like delete.
	ErrIDInvalid modelError = "models: invalid ID was provided"
	// ErrInvalidEmail is returned when incorrect email is provided
	ErrInvalidEmail modelError = "models: invalid email provided"
	// ErrPasswordInCorrect is returned when incorrect password is provided during authentication
	ErrPasswordInCorrect modelError = "models: incorrect password provided"
	// ErrEmailRequired is returened when an email address is not provided
	// when creating a user
	ErrEmailRequired modelError = "models: email address is required"
	// ErrEmailInvalid is returned when an email address provided
	// does not match any of the requirements
	ErrEmailInvalid modelError = "models: email address is invalid"
	// ErrEmailTaken is returned when update or create is attempted with email that is already in use
	ErrEmailTaken modelError = "models: email address is already taken "
	// ErrPasswordTooShort is returned when update or create is attempted with a user
	// password that is less than 8 characters
	ErrPasswordTooShort modelError = "models: password must be atleast 8 characters long"
	// ErrPasswordRequired is returned when create is attempted without a user password
	ErrPasswordRequired modelError = "models: password is required"
	// ErrRememberTooShort is returned when a remember token is not atleast 32 bytes
	ErrRememberTooShort modelError = "model: remember token must be atleast 32 bytes"
	// ErrRememberRequired is returned when create or update is attempted without
	// a valid user remember token hash
	ErrRememberRequired modelError = "models: remember token is required"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = strings.Title(split[0])
	return strings.Join(split, " ")
}
