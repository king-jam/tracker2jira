package backend

import (
	"github.com/king-jam/tracker2jira/rest/models"
)

const (
	// this is the key prefix for storage of all user objects, the format will be
	// /<storage prefix - should be t2j>/<t2j instance ID>/users/<user ID>/<User Object>
	usersPath = "users"
)

// UserBackend interface encapsulates all the implementations of the user peristence
type UserBackend interface {
	GetUsers() ([]*models.User, error)
	GetUserByID(userid string) (*models.User, error)
	PutUser(user *models.User) (*models.User, error)
	DeleteUser(userid string) error
}

// GetUsers returns an array of User objects to the caller
func (b *Backend) GetUsers() ([]*models.User, error) {
	users := []*models.User{}
	key := b.getUserBase()
	values, err := b.store.List(key)
	if len(values) == 0 {
		return users, nil
	}
	if err != nil {
		return users, err
	}
	for _, v := range values {
		user := &models.User{}
		err = user.UnmarshalBinary(v.Value)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByID returns a User object by ID
func (b *Backend) GetUserByID(userid string) (*models.User, error) {
	user := &models.User{}
	key := b.getUserBase() + userid
	pair, err := b.store.Get(key)
	if err != nil {
		return user, err
	}
	err = user.UnmarshalBinary(pair.Value)
	if err != nil {
		return user, err
	}
	return user, nil
}

// PutUser stores a fully formed user model into the DB
func (b *Backend) PutUser(user *models.User) (*models.User, error) {
	key := b.getUserBase() + user.UserID.String()
	value, err := user.MarshalBinary()
	if err != nil {
		return user, err
	}
	err = b.store.Put(key, value, nil)
	if err != nil {
		return user, err
	}
	return user, nil
}

// DeleteUser removes a User entry from the DB by ID
func (b *Backend) DeleteUser(userid string) error {
	key := b.getUserBase() + userid
	err := b.store.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// getUserBase returns the user base path
func (b *Backend) getUserBase() string {
	return b.instanceID + "/" + usersPath + "/"
}
