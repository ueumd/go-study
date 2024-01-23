package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

//type Model struct {
//	ID        uint `gorm:"primarykey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt DeletedAt `gorm:"index"`
//}

type userinfo struct {
	// gorm.Model
	ID   uint   `gorm:"primary_key:id,size:10;not null"`
	Name string `gorm:"type:varchar(50)" json:"name"`
	Age  uint   `gorm:"size:3"`
	// 使用指针为了传值时可以存空值
	Email     *string `gorm:"type:varchar(50);comment:邮箱"`
	Address   *string `gorm:"type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // soft delete
}

// 创建时的Hook
func (user *userinfo) BeforeCreate(tx *gorm.DB) (err error) {
	email := fmt.Sprintf("%s@qq.com", user.Name)
	user.Email = &email
	return nil
}

var db *gorm.DB

func initMySQLConn() error {

	username := "root"          //账号
	password := "Abcdef@123456" //密码
	host := "127.0.0.1"         //数据库地址，可以是Ip或者域名
	port := 3306                //数据库端口
	Dbname := "gin_vr"          //数据库名
	timeout := "10s"            //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 使用彩色打印
		},
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 默认日志只打印错误和慢sql
		Logger: newLogger,

		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,

		// 禁用默认事务（提高运行速度）
		// 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。
		// 如果没有这方面的要求，您可以在初始化时禁用它，这样可以获得60%的性能提升
		SkipDefaultTransaction: true,

		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			// 是否单表  使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,

			NoLowerCase: false, // 是否小写转换
		},
	})

	db.Debug() // logger

	return err

}

func add() {
	//email := "xx@qq.com"
	//user := userinfo{
	//	Name:    "haha",
	//	Age:     18,
	//	Email:   &email, // 指针类型
	//	Address: nil,
	//}
	// db.Create(&user)

	// 批量插入
	var list []userinfo

	for i := 1; i <= 10; i++ {
		list = append(list, userinfo{
			Name:    fmt.Sprintf("haha-%d", i),
			Age:     uint(18 + i),
			Email:   nil,
			Address: nil,
		})
	}

	db.Create(&list)
}

func query() {
	// 查询单调记录
	var userinfo1 userinfo

	// SELECT * FROM `userinfo` LIMIT 1
	db.Take(&userinfo1)
	fmt.Println(userinfo1)

	userinfo1 = userinfo{}
	//SELECT * FROM `userinfo` ORDER BY `userinfo`.`id` LIMIT 1
	db.First(&userinfo1)
	fmt.Println(userinfo1)

	userinfo1 = userinfo{}
	// SELECT * FROM `userinfo` ORDER BY `userinfo`.`id` DESC LIMIT 1
	db.Last(&userinfo1)
	fmt.Println(userinfo1)

	userinfo1 = userinfo{}
	//  SELECT * FROM `userinfo` WHERE name = 'haha-10' LIMIT 1
	db.Take(&userinfo1, "name = ?", "haha-10")
	fmt.Println(userinfo1)

	//userinfo1 = userinfo{}
	//err2 := db.Where("name=?", "haha-1").First(&userinfo1).Error
	//if err2 != nil {
	//	panic(err2)
	//}

	// 查询多条记录
	var userListCount []userinfo
	count := db.Find(&userListCount).RowsAffected
	fmt.Println("count: ", count)

	//for _, user := range userListCount {
	//	fmt.Println(user)
	//}

	// 转JSON
	data, _ := json.Marshal(userListCount)
	fmt.Println(string(data))

	/******主健ID查*******/
	//  SELECT * FROM `userinfo` WHERE `userinfo`.`id` IN (1,3,5)
	userinfo1 = userinfo{}
	db.Find(&userinfo1, []int{1, 3, 5})
	fmt.Println(userinfo1)
	// {1 haha 18 0xc0001f82d0 <nil>}

	// SELECT * FROM `userinfo` WHERE `userinfo`.`id` IN (1,3,5)
	var list2 []userinfo
	db.Find(&list2, []int{1, 3, 5})
	fmt.Println(list2)
	// [{1 haha 18 0xc0001f8370 <nil>} {3 haha-1 19 <nil> <nil>} {5 haha-3 21 <nil> <nil>}
}

func update() {
	// save 保存所有字段

	// 单个字段更新

	var user userinfo

	// SELECT * FROM `userinfo` WHERE `userinfo`.`id` = 11 LIMIT 1

	db.Take(&user, 11)
	user.Name = "名字改了"

	address := "上海市"
	user.Address = &address
	// UPDATE `userinfo` SET `name`='名字改了',`age`=27,`email`=NULL,`address`=NULL WHERE `id` = 1
	db.Save(&user)

	db.Save(&user)

	var list []userinfo

	// UPDATE `userinfo` SET `address`='上海市' WHERE `userinfo`.`id` IN (1,11,12) AND `id` IN (1,11,12)
	//db.Find(&list, []int{1, 11, 12}).Update("Address", "上海市")
	//db.Find(&list, []int{1, 11, 12}).Updates(userinfo{
	//	Name: "名字改了",
	//})
	db.Find(&list, []int{1, 11, 12}).Updates(map[string]any{
		"Name": "名字改了哦",
	})
}

func del() {
	var user userinfo

	// 清空所有记录
	// DELETE FROM `users` WHERE 1=1
	// db.Where("1 = 1").Delete(&user)

	// 有了 DeletedAt gorm.DeletedAt 则为软删除，没有硬删除
	// UPDATE `userinfo` SET `deleted_at`='2023-12-04 14:26:11.563' WHERE `userinfo`.`id` IN (3,5) AND `userinfo`.`deleted_at` IS NULL
	db.Delete(&user, []int{3, 5})

	var list []userinfo

	// 软删除普通不会再查询到
	affected := db.Find(&list).RowsAffected
	fmt.Println(affected)
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func update2() {
	var user userinfo

	user.ID = 4
	// UPDATE `userinfo` SET `name`='haha',`updated_at`='2023-12-04 14:53:11.435' WHERE `userinfo`.`deleted_at` IS NULL
	res := db.Model(&user).Update("name", "haha")
	fmt.Println(res.RowsAffected)

	user = userinfo{}
	db.Where(&userinfo{Name: "hello7", Age: 25}).First(&user)
	fmt.Println(user)

	user = userinfo{}
	db.Where("name = ?", "hello8").First(&user)
	fmt.Println(user)

}

func query2() {
	var user userinfo
	db.Raw("select * from userinfo where name = ?", "hello7").Scan(&user)

	fmt.Println(user)
}

func main() {
	err := initMySQLConn()
	if err != nil {
		return
	}

	// AutoMigrate 只新增，不修改，大小会修改
	// db.AutoMigrate(&userinfo{})

	// add()

	// update()

	// del()

	// update2()

	query2()

}

// [单条]数据查询
func GetOne[T any](data T, query string, args ...any) T {
	err := db.Where(query, args...).First(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return data
}
