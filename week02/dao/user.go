package dao

import (
	"database/sql"
	"errors"
	xerrors "github.com/pkg/errors"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"title,"`
}

func QueryUser(id int64) (*User, error) {
	//执行查询语句
	row := db.QueryRow("select * from user where id = ?", id)

	user := new(User)
	err := row.Scan(&user.ID, &user.Name)

	switch {
	case err == nil:
		return user, nil
	case errors.Is(err, sql.ErrNoRows):
		return nil, err
	default:
		return nil, xerrors.Wrap(err, "数据库操作失败")
	}

}
