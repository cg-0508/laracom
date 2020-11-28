package repo

import (
	"github.com/cg-0508/laracom/user-service/model"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Repository interface {
	Create(user *model.User) error
	Get(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Update(user *model.User) error
}

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) Create(user *model.User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Get(id string) (*model.User, error) {
	var user *model.User
	id1, _ := strconv.ParseUint(id, 10, 64)
	user.ID = uint(id1)
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := repo.Db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}



func (repo *UserRepository) Update(user *model.User) error {
	if err := repo.Db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
