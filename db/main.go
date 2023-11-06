package db

import (
	"bee-go-demo/model"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func Init() {
	// set default database
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// set default database
	orm.RegisterDataBase("default", "postgres", "postgres://user:user123@localhost:5443/myuser?sslmode=disable&connect_timeout=5")

	// register model
	orm.RegisterModel(new(model.User))

	// create table
	orm.RunSyncdb("default", false, true)
}
