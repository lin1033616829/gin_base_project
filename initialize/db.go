package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mmfile/global"
	"mmfile/store/sqlstore"
	"time"
)

func InitDb() {
	fmt.Println("初始化数据库")

	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v",
		global.ServerConf.Mysql.Username,
		global.ServerConf.Mysql.Password,
		global.ServerConf.Mysql.HostName,
		global.ServerConf.Mysql.Port,
		global.ServerConf.Mysql.Dbname,
		global.ServerConf.Mysql.Config,
	)

	fmt.Println(fmt.Sprintf("mysql dataSource %v", dataSource))

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dataSource,
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(fmt.Sprintf("mysql connect err1 %v", err.Error()))
		panic("failed to connect database")
	}

	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(fmt.Sprintf("mysql connect err2 %v", err.Error()))
		panic("failed to connect database")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db = db.Debug()
	_ = sqlstore.NewSqlSupplier(db, global.ServerLog)

	fmt.Println("数据库连接成功")

}
