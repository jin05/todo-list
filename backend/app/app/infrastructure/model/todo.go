package model

import (
	"github.com/kayac/ddl-maker/dialect"
	"github.com/kayac/ddl-maker/dialect/mysql"
	"time"
)

type Todo struct {
	ID        int64 `ddl:"auto"`
	UserID    int64
	Title     string    `ddl:"size=255"`
	Content   string    `ddl:"type=longtext"`
	Check     bool      `ddl:"default=false"`
	CreatedAt time.Time `ddl:"default=CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `ddl:"default=CURRENT_TIMESTAMP"`
}

func (m Todo) Table() string {
	return "todos"
}

func (m Todo) PrimaryKey() dialect.PrimaryKey {
	return mysql.AddPrimaryKey("id")
}

func (m Todo) Indexes() dialect.Indexes {
	return dialect.Indexes{
		mysql.AddIndex("user_id_idx", "user_id"),
	}
}
