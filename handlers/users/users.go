package users

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/server/operations/users"
)

// GetUser ...
func GetUser(db *backend.Backend, params users.GetUserByIDParams) middleware.Responder {
	value, err := db.GetUserByID(params.UserID)
	if err != nil {
		return &users.GetUserByIDNotFound{}
	}
	return &users.GetUserByIDOK{
		Payload: value,
	}
}

// GetUsers ...
func GetUsers(db *backend.Backend, params users.GetUsersParams) middleware.Responder {
	values, err := db.GetUsers()
	if err != nil {
		return &users.GetUsersBadRequest{}
	}
	return &users.GetUsersOK{
		Payload: values,
	}
}

// PostUser ...
func PostUser(db *backend.Backend, params users.PostUserParams) middleware.Responder {
	value, err := db.PutUser(params.Body)
	if err != nil {
		return &users.PostUserBadRequest{}
	}
	return &users.PostUserCreated{
		Payload: value,
	}
}
