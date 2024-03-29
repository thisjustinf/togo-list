// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ResponseData interface {
	IsResponseData()
}

type CreateTodoDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type ResponseStatus struct {
	Status  State        `json:"status"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data,omitempty"`
}

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	UserID      uuid.UUID `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (Todo) IsResponseData() {}

type UpdateTodoDto struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   *string `json:"completed,omitempty"`
}

type UpdateUserDto struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (User) IsResponseData() {}

type State string

const (
	StateSuccess State = "SUCCESS"
	StateError   State = "ERROR"
)

var AllState = []State{
	StateSuccess,
	StateError,
}

func (e State) IsValid() bool {
	switch e {
	case StateSuccess, StateError:
		return true
	}
	return false
}

func (e State) String() string {
	return string(e)
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = State(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid State", str)
	}
	return nil
}

func (e State) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
