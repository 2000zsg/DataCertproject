package models

import (
	"DataCertproject/qkl_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int64	`form:"id"`
	Phone int	`form:"phone"`
	Password string	 `form:"password"`
}

func (u User)SaveUser() (int64,error) {
	mdHash:=md5.New()
	mdHash.Write([]byte(u.Password))
	bytes:=mdHash.Sum(nil)
	u.Password =hex.EncodeToString(bytes)
	result, err := qkl_mysql.Db.Exec("insert into qkl (phone, password) "+"values(?,?)", u.Phone,u.Password )
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (u User)QueryUser()(*User,error){
	row:=qkl_mysql.Db.QueryRow("select phone from qkl where phone = ? and password = ?", u.Phone,u.Password)
	err:=row.Scan(&u.Phone)
	if err !=nil{
		return nil, err
	}
	return &u,nil
}