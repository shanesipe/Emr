package models

import "golang.org/x/crypto/bcrypt"

// User represents an application user
type User struct {
  ID       string `json:"id"`
  Name     string `json:"name"` 
  Email    string `json:"email"`
  Password string `json:"-"`
}

// HashPassword hashes the password before storing 
func (u *User) HashPassword(password string) error {
  hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
  if err != nil {
    return err
  }
  
  u.Password = string(hashed)
  return nil
}

// CheckPassword compares the hashed password 
func (u *User) CheckPassword(password string) error {
  return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))  
}

// NewUser creates a new user
func NewUser(name, email, password string) (*User, error) {
  user := &User{
    Name: name,
    Email: email,
  }
  
  err := user.HashPassword(password)
  if err != nil {
    return nil, err
  }
  
  return user, nil
}

