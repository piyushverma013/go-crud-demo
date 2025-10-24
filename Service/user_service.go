package service

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserService defines CRUD methods
type UserService interface {
	CreateUser(u User) User
	GetAllUsers() []User
	GetUserByID(id int) (User, bool)
	UpdateUser(id int, u User) (User, bool)
	DeleteUser(id int) bool
}

// InMemoryUserService is a simple in-memory implementation
type InMemoryUserService struct {
	users  []User
	nextID int
}

func NewUserService() *InMemoryUserService {
	return &InMemoryUserService{
		users:  []User{},
		nextID: 1,
	}
}

func (s *InMemoryUserService) CreateUser(u User) User {
	u.ID = s.nextID
	s.nextID++
	s.users = append(s.users, u)
	return u
}

func (s *InMemoryUserService) GetAllUsers() []User {
	return s.users
}

func (s *InMemoryUserService) GetUserByID(id int) (User, bool) {
	for _, u := range s.users {
		if u.ID == id {
			return u, true
		}
	}
	return User{}, false
}

func (s *InMemoryUserService) UpdateUser(id int, u User) (User, bool) {
	for i, existing := range s.users {
		if existing.ID == id {
			u.ID = id
			s.users[i] = u
			return u, true
		}
	}
	return User{}, false
}

func (s *InMemoryUserService) DeleteUser(id int) bool {
	for i, u := range s.users {
		if u.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return true
		}
	}
	return false
}
