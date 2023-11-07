package db

import (
	"bee-go-demo/config"
	"bee-go-demo/model"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func Init(cfg *config.Configuration) {
	// set default database
	orm.RegisterDriver("postgres", orm.DRPostgres)
	fmt.Println("DbDsn", cfg.DbDsn)
	// set default database
	orm.RegisterDataBase("default", "postgres", cfg.DbDsn)

	// register model
	orm.RegisterModel(new(model.User))

	// create table
	orm.RunSyncdb("default", false, true)
}
