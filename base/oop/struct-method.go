package main

/*
go struct 方法使用指针与不使用指针的区别?
*/

type User struct {
	username string
}

func (u User) SetUsername1(username string) {
	u.username = username
}

func (u *User) SetUsername2(username string) {
	u.username = username
}

// 上面等价于如下 一个是值传递，一个是指针传递
func SetUsername1(u User, username string) {
	u.username = username
}

func SetUsername2(u *User, username string) {
	u.username = username
}
