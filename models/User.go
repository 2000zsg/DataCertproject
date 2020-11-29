package models

import (
	"DataCertproject/qkl_mysql"
	"DataCertproject/util"
	"hl/db_mysql"
)

type User struct {
	Id       int64  `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Card     string `form:"card"`
	Sex      string `form:"sex"`
}

func (u User) Update() (int64, error) {
	rs, err := db_mysql.DB.Exec("update qkl set name=?, card=?, sex=? where phone=?", u.Name, u.Card, u.Sex, &u.Phone)
	if err != nil {
		return -1, err
	}
	return rs.RowsAffected()
}
func (u User) SaveUser() (int64, error) {
	//mdHash:=md5.New()
	//mdHash.Write([]byte(u.Password))
	//bytes:=mdHash.Sum(nil)
	//u.Password = hex.EncodeToString(bytes)
	u.Password = util.MD5HashString(u.Password)
	result, err := qkl_mysql.Db.Exec("insert into qkl (phone, password) "+"values(?,?)", u.Phone, u.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (u User) QueryUser() (*User, error) {
	//mdHash:=md5.New()
	//mdHash.Write([]byte(u.Password))
	//bytes:=mdHash.Sum(nil)
	//u.Password = hex.EncodeToString(bytes)
	u.Password = util.MD5HashString(u.Password)
	row := qkl_mysql.Db.QueryRow("select phone,name,card,sex from qkl where phone = ? and password = ?", u.Phone, u.Password)
	err := row.Scan(&u.Phone, &u.Sex, &u.Name, &u.Card)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func QueryUserByPhone(phone string) (*User, error) {
	row := db_mysql.DB.QueryRow("select phone, name, card, sex from qkl where phone = ?", phone)
	var user User
	err := row.Scan(&user.Phone, &user.Name, &user.Card, &user.Sex)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
