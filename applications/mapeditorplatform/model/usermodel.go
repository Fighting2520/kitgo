package model

import (
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	User struct {
		Id         int       `db:"id,omitempty"`
		Username   string    `db:"username,omitempty"`
		Password   string    `db:"password,omitempty"`
		RoleId     int       `db:"role_id,omitempty"`
		CreatedAt  time.Time `db:"created_at,omitempty"`
		ModifiedAt time.Time `db:"modified_at,omitempty"`
	}

	UserModel struct {
		conn  sqlx.SqlConn
		table string
	}
)

func NewUserModel(dsn, table string) *UserModel {
	return &UserModel{
		table: table,
		conn:  sqlx.NewSqlConn("mysql", dsn),
	}
}

func (um *UserModel) Insert(user *User) error {
	querySql := fmt.Sprintf("INSERT INTO %s (username, password, role_id) VALUE (?,?,?)", um.table)
	_, err := um.conn.Exec(querySql, user.Username, user.Password, user.RoleId)
	return err
}

func (um *UserModel) FindByUsername(username string) (*User, error) {
	var user User
	querySql := fmt.Sprintf("SELECT id, username, password, role_id, created_at, modified_at FROM %s where username=?", um.table)
	if err := um.conn.QueryRow(&user, querySql, username); err != nil {
		if err == sqlc.ErrNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}
