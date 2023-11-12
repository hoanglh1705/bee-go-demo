package user

import (
	"bee-go-demo/model"
	"context"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/client/orm"
)

type UserRepository interface {
	Insert(ctx context.Context, user *model.User) (*model.User, error)
}

type userRepository struct {
	o orm.Ormer
}

func NewUserRepository(o orm.Ormer) UserRepository {
	return &userRepository{o: o}
}

func (r *userRepository) Insert(ctx context.Context, user *model.User) (*model.User, error) {
	to, err := r.o.Begin()
	if err != nil {
		logs.Error("start the transaction failed")
		return nil, err
	}
	var id int64
	id, err = to.Insert(user)
	if err != nil {
		logs.Error("execute transaction's sql fail, rollback.", err)
		err = to.Rollback()
		if err != nil {
			logs.Error("roll back transaction failed", err)
			return nil, err
		}
		return nil, err
	} else {
		err = to.Commit()
		if err != nil {
			logs.Error("commit transaction failed.", err)
			return nil, err
		}
	}

	newUser := &model.User{Id: int(id)}
	if err := r.o.Read(newUser); err != nil {
		logs.Error("read user failed.", err)
		return nil, err
	}
	return newUser, nil
}
