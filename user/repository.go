package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	FindByID(id int) (User, error)
	Create(input UserRegister) (User, error)
	Update(id int, input UserUpdate) (User, error)
	Delete(id int) (User, error)
	DeletePermanent(id int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindByID(id int) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *repository) Create(input UserRegister) (User, error) {
	var createdUser User
	err := r.db.Create(&User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Addres:   input.Addres,
		Gender:   input.Gender,
	}).Error

	return createdUser, err
}

func (r *repository) Update(id int, input UserUpdate) (User, error) {
	var user User
	err := r.db.Model(&user).Where("id = ?", id).Updates(User{
		Name:   input.Name,
		Email:  input.Email,
		Phone:  input.Phone,
		Addres: input.Addres,
		Gender: input.Gender,
	}).Error
	return user, err
}

func (r *repository) Delete(id int) (User, error) {
	var user User
	err := r.db.Delete(&user, id).Error
	return user, err
}

func (r *repository) DeletePermanent(id int) (User, error) {
	var user User
	err := r.db.Unscoped().Delete(&user, id).Error
	return user, err
}
