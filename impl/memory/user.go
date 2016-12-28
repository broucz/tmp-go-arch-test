package memory

import (
	"errors"
	"sync"

	"github.com/broucz/tmp-go-arch-test/domain/user"
)

type userRepository struct {
	mtx       sync.RWMutex
	container map[user.ID]*user.User
}

func (r *userRepository) Store(u *user.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.container[u.ID] = u
	return nil
}

func (r *userRepository) Find(id user.ID) (*user.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if result, exists := r.container[id]; exists {
		return result, nil
	}
	return nil, errors.New("User not found")
}

func (r *userRepository) FindAll() []*user.User {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	result := make([]*user.User, 0, len(r.container))
	for _, val := range r.container {
		result = append(result, val)
	}
	return result
}

// NewUserRepository returns a new instance of a in-memory user repository.
func NewUserRepository() user.Repository {
	return &userRepository{
		container: make(map[user.ID]*user.User),
	}
}
