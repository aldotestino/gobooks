package gobooks

import "github.com/uptrace/bun"

type Book struct {
	bun.BaseModel `bun:"table:books,alias:b"`

	ID     int64  `bun:"id,pk,autoincrement" json:"id"`
	Title  string `bun:"title,notnull" json:"title"`
	Author string `bun:"author,notnull" json:"author"`
}
