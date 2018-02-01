package backend

import (
	log "github.com/sirupsen/logrus"

	"github.com/king-jam/tracker2jira/rest/models"
)

// TODO: IDs should be returned from the backing store or "appear" to be

const usersPath = "users"

// UserBackend interface encapsulates all the implementations of the task peristence
type UserBackend interface {
	GetUsers() ([]*models.User, error)
	GetUserByID(userid string) (*models.User, error)
	PutUser(user *models.User) (*models.User, error)
	DeleteUser(userid string) error
}

// GetUsers is ...
func (b *Backend) GetUsers() ([]*models.User, error) {
	key := b.getUserBase()
	values, err := b.store.List(key)
	if len(values) == 0 {
		return []*models.User{}, nil
	}
	if err != nil {
		return []*models.User{}, err
	}
	users := []*models.User{}
	for _, v := range values {
		user := &models.User{}
		err = user.UnmarshalBinary(v.Value)
		if err != nil {
			return []*models.User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByID ...
func (b *Backend) GetUserByID(userid string) (*models.User, error) {
	key := b.getUserBase() + userid
	pair, err := b.store.Get(key)
	if err != nil {
		log.Printf("no version")
	}
	user := &models.User{}
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
		return nil, err
	}
	err = b.store.Put(key, value, nil)
	if err != nil {
		return nil, err
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
