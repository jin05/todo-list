package model

import (
	"github.com/kayac/ddl-maker/dialect"
	"github.com/kayac/ddl-maker/dialect/mysql"
	"time"
)

type User struct {
	ID        int64     `ddl:"auto"`
	UserName  string    `ddl:"size=255"`
	AuthID    string    `ddl:"size=191"`
	Email     string    `ddl:"size=255"`
	CreatedAt time.Time `ddl:"default=CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `ddl:"default=CURRENT_TIMESTAMP"`
}

func (m User) Table() string {
	return "users"
}

func (m User) PrimaryKey() dialect.PrimaryKey {
	return mysql.AddPrimaryKey("id")
}

func (m User) Indexes() dialect.Indexes {
	return dialect.Indexes{
		mysql.AddUniqueIndex("auth_id_idx", "auth_id"),
		mysql.AddUniqueIndex("email_idx", "email"),
	}
}
