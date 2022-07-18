package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test_sequence(t *testing.T) {
	// DSN:Data Source Name
	dsn := "dksl:123456@tcp(127.0.0.1:13306)/employees"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面

	sqlStr := "INSERT INTO student (uid, name, score, nickname, gender, birth_year, created_at, modified_at) VALUES (?, ?, 95, ?, 0, 16, NOW(), NOW());"
	for i := 41; i < 1000; i++ {
		ret, err := db.Exec(sqlStr, i, fmt.Sprintf("%s_%d", "scott", i), fmt.Sprintf("%s_%d", "nc_scott", i))
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
		theID, err := ret.LastInsertId() // 新插入数据的id
		if err != nil {
			fmt.Printf("get lastinsert ID failed, err:%v\n", err)
			return
		}
		fmt.Printf("insert success, the id is %d.\n", theID)
	}
}
