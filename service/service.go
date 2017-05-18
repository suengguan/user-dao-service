package service

import (
	"dao-service/user-dao-service/dao"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

type Service struct {
}

func (this *Service) Create(user *model.User) error {
	beego.Debug("create user")
	var err error
	userDao := dao.NewUserDao()

	err = userDao.Create(user)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "create user failed")
		return err
	}
	beego.Debug(*user)

	return err
}

func (this *Service) GetByName(name string) (*model.User, error) {
	beego.Debug("GetByName")
	var err error
	var user *model.User
	var userDao = dao.NewUserDao()

	user, err = userDao.GetByName(name)
	if err != nil {
		//err = fmt.Errorf("%s", "user name is not existed!")
		return nil, err
	}

	beego.Debug("result:", *user)

	return user, err
}

func (this *Service) GetById(userId int64) (*model.User, error) {
	var err error
	var user *model.User
	var userDao = dao.NewUserDao()

	user, err = userDao.GetById(userId)
	if err != nil {
		err = fmt.Errorf("%s", "user not existed!")
		return nil, err
	}

	return user, err
}

func (this *Service) GetByRole(userRole int) (*model.User, error) {
	var err error
	var user *model.User
	var userDao = dao.NewUserDao()

	if userRole != model.USER_AUTHORITY_ADMIN &&
		userRole != model.USER_AUTHORITY_USER {
		err = fmt.Errorf("%s", "is invalid user role!")
		return nil, err
	}

	user, err = userDao.GetByRole(userRole)
	if err != nil {
		err = fmt.Errorf("%s", "user not existed!")
		return nil, err
	}

	return user, err
}

func (this *Service) GetAll() ([]*model.User, error) {
	var err error
	var users []*model.User
	var userDao = dao.NewUserDao()

	// get users
	users, err = userDao.GetAll()
	if err != nil {
		return nil, err
	}

	return users, err
}

func (this *Service) GetAllExcludeOneId(userId int64) ([]*model.User, error) {
	var err error
	var users []*model.User
	var userDao = dao.NewUserDao()

	_, err = userDao.GetById(userId)
	if err != nil {
		err = fmt.Errorf("%s", "exclude id is not existed")
		return nil, err
	}

	// get users
	users, err = userDao.GetAllExcludeOneId(userId)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (this *Service) DeleteById(userId int64) error {
	var err error
	userDao := dao.NewUserDao()

	err = userDao.DeleteById(userId)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s%s", "delete user failed! reason:", err.Error())
		return err
	}

	return err
}

func (this *Service) Update(user *model.User) (*model.User, error) {
	var err error
	userDao := dao.NewUserDao()

	beego.Debug("->update")
	err = userDao.Update(user)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s%s", "update user failed! reason:", err.Error())
		return nil, err
	}
	beego.Debug("result:", *user)

	return user, err
}
