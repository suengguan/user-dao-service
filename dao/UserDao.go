package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"model"
)

type UserDao struct {
	m_Orm        orm.Ormer
	m_QuerySeter orm.QuerySeter
	m_QueryTable *model.User
}

func NewUserDao() *UserDao {
	d := new(UserDao)
	cfg := beego.AppConfig

	d.m_Orm = orm.NewOrm()
	d.m_Orm.Using(cfg.String("dbname"))

	d.m_QuerySeter = d.m_Orm.QueryTable(d.m_QueryTable)
	d.m_QuerySeter.Limit(-1)

	return d
}

//add
func (this *UserDao) Create(user *model.User) error {
	_, err := this.m_Orm.Insert(user)
	// if err != nil {
	// 	beego.Debug(num, err)
	// }

	return err
}

// delete
func (this *UserDao) DeleteById(id int64) error {
	num, err := this.m_QuerySeter.Filter("ID", id).Delete()

	if err != nil {
		return err
	}

	if num < 1 {
		err = fmt.Errorf("%s", "there is no user to delete")
		return err
	}

	return err
}

func (this *UserDao) Update(user *model.User) error {
	num, err := this.m_Orm.Update(user)

	if err != nil {
		return err
	}

	if num < 1 {
		beego.Debug("there is no user to update")
		//err = fmt.Errorf("%s", "there is no user to update")
		//return err
	}

	return err
}

// find
func (this *UserDao) GetById(id int64) (*model.User, error) {
	var user model.User

	err := this.m_QuerySeter.Filter("ID", id).One(&user)
	// if err != nil {
	// 	beego.Warning(err)
	// 	return nil, err
	// }

	return &user, err
}

func (this *UserDao) GetByName(name string) (*model.User, error) {
	var user model.User
	var count int64
	var err error

	qs := this.m_QuerySeter.Filter("NAME", name)

	count, err = qs.Count()
	if err != nil {
		return nil, err
	}

	if count > 1 {
		err = fmt.Errorf("%s%s", "there is one more user ", name)
		return nil, err
	}

	if count < 1 {
		err = fmt.Errorf("%s", "user is not existed")
		return nil, err
	}

	err = qs.One(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (this *UserDao) GetByEmail(email string) (*model.User, error) {
	var user model.User

	err := this.m_QuerySeter.Filter("EMAIL", email).One(&user)
	// if err != nil {
	// 	beego.Warning(err)
	// 	return nil, err
	// }

	return &user, err
}

func (this *UserDao) GetByPhone(phone string) (*model.User, error) {
	var user model.User

	err := this.m_QuerySeter.Filter("PHONE", phone).One(&user)
	// if err != nil {
	// 	beego.Warning(err)
	// 	return nil, err
	// }

	return &user, err
}

func (this *UserDao) GetAllExcludeByName(name string) ([]*model.User, error) {
	var users []*model.User

	_, err := this.m_QuerySeter.Exclude("NAME", name).All(&users)
	// if err != nil {
	// 	beego.Warning(num, err)
	// 	return nil, err
	// }

	return users, err
}

func (this *UserDao) GetAllExcludeOneId(id int64) ([]*model.User, error) {
	var users []*model.User

	_, err := this.m_QuerySeter.Exclude("ID", id).All(&users)

	return users, err
}

func (this *UserDao) GetByRole(role int) (*model.User, error) {
	var user model.User

	err := this.m_QuerySeter.Filter("ROLE", role).One(&user)
	// if err != nil {
	// 	beego.Warning(err)
	// 	return nil, err
	// }

	return &user, err
}

func (this *UserDao) GetAll() ([]*model.User, error) {
	var users []*model.User

	_, err := this.m_QuerySeter.All(&users)

	// if err != nil {
	// 	beego.Debug(num, err)
	// }

	return users, err
}
