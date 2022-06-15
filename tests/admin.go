package main

import (
	"easy-apis/easy/bootstrap"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func testSignUp() {
	// 加密
	password := "123456"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encrypt := string(hash)
	fmt.Println(encrypt)

	// 校验
	err = bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(password))
	if err != nil {
		fmt.Println("校验失败")
	} else {
		fmt.Println("校验成功")
	}
}

func testMysql() {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, bootstrap.MysqlUser, bootstrap.MysqlPassword, bootstrap.MysqlHost, bootstrap.MysqlPort, bootstrap.MysqlDb, bootstrap.MysqlCharset)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   bootstrap.MysqlPrefix,
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
		DisableAutomaticPing: false,
	})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	fmt.Println("连接成功")
	fmt.Println(Db)
}

func main() {
	//testSignUp()
	testMysql()
}
