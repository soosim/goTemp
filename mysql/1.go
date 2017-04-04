package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// insert()
	// query()
	// query_dict()
	update()
}

func insert() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/sys?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`INSERT INTO test(name) VALUES(?)`)
	checkErr(err)

	res, err := stmt.Exec("xisdae")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func query() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/sys?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string

		rows.Columns()
		err = rows.Scan(&id, &name)
		checkErr(err)

		fmt.Println(id)
		fmt.Println(name)
	}
}

func query_dict() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/sys?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

func update() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/sys?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`UPDATE test SET name=? WHERE id = ?`)
	checkErr(err)
	res, err := stmt.Exec("xiejinlong", 1)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func remove() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/sys?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`DELETE FROM test WHERE id = ?`)
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
