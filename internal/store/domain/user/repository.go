package user

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

//interface
type IUserRepository interface {
	Migration()
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]User, error)
	FindByID(id uuid.UUID) (*User, error)
	FindByName(name string) (*User, error)
	Search(s string) ([]User, error)
	GetWithRoles(id uuid.UUID) (*User, error)
	Seed()
}

var UserRepository IUserRepository = &userRepository{}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
func (r *userRepository) Migration() {
	r.db.AutoMigrate(&User{})
}

//Create a new user
func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

//Update a user
func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

//Delete a user
func (r *userRepository) Delete(user *User) error {
	return r.db.Delete(user).Error
}

//Delete a user by id
func (r *userRepository) DeleteByID(id uuid.UUID) error {
	user := User{}
	user.ID = id
	return r.db.Delete(&user).Error
}

//Find all users
func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

//Find a user by id
func (r *userRepository) FindByID(id uuid.UUID) (*User, error) {
	user := User{}
	err := r.db.First(&user, id).Error
	return &user, err
}

//Find a user by name
func (r *userRepository) FindByName(name string) (*User, error) {
	user := User{}
	err := r.db.Where("name = ?", name).First(&user).Error
	return &user, err
}

//Search users by name or email
func (r *userRepository) Search(s string) ([]User, error) {
	var users []User
	err := r.db.Where("name LIKE ? OR email LIKE ?", "%"+s+"%", "%"+s+"%").Find(&users).Error
	return users, err
}

//Get with roles relationship //TODO: Check this
func (r *userRepository) GetWithRoles(id uuid.UUID) (*User, error) {
	user := User{}
	err := r.db.Preload("Roles").First(&user, id).Error
	return &user, err
}

//Seed a user
func (r *userRepository) Seed() {
	users := []User{
		{Name: "John Doe", Email: "johnDoe@gmail.com", Password: "123456", IsActive: true}, //TODO: encrypt password
	}
	for _, user := range users {
		r.Create(&user)
	}
}
