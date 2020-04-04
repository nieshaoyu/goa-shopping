package dao

import (
	"github.com/lib/pq"
)

// 是否存在唯一索引错误
func IsPGUniqueIndexErr(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		return pgErr.Code == "23505"
	}
	return false
}
