package backend

import (
	"log"

	"github.com/king-jam/tracker2jira/rest/models"
	uuid "github.com/satori/go.uuid"
)

const usersPath = "users"

// GetUsers is ...
func (b *Backend) GetUsers() ([]*models.User, error) {
	key := b.GetUserBase()
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
	key := b.GetUserBase() + userid
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

// PutUser ...
func (b *Backend) PutUser(user *models.User) (*models.User, error) {
	uuid := uuid.NewV4()
	key := b.GetUserBase() + uuid.String()
	user.UserID = uuid.String()
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

// DeleteUser ...
func (b *Backend) DeleteUser(userid string) error {
	key := b.GetUserBase() + userid
	err := b.store.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// GetUserBase returns the user base path
func (b *Backend) GetUserBase() string {
	return b.instanceID + "/" + usersPath + "/"
}
