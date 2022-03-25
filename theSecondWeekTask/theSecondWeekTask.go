package theSecondWeekTask

import (
	"database/sql"
	"errors"
	"log"
)

type Data struct {
	Something string
}

//错误不用抛给上层代码，这一层直接处理错误，返回空数据
func Dao() *Data {
	var (
		err  = errors.New("something is unknown")
		data = Data{}
	)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Printf("错误不用抛给上层代码，这一层直接处理错误，返回空数据,\n%s", err)
	}
	return &data
}
