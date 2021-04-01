package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Phone struct {
	Id     int
	Name   string
	Number string
}

var db *sql.DB
var err error

const (
	createTableCommand = `CREATE TABLE IF NOT EXISTS phones (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
 		name TEXT NOT NULL,
		number TEXT NOT NULL)`  // STRING型だと080→80と登録されてしまうためTEXT型にしている
	getPhonesCommand    = `SELECT * FROM phones`
	insertPhonesCommand = `INSERT INTO phones (name, number) VALUES(?, ?)`
)

func init() {
	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("DB Openに失敗しました")
	}
	if _, err := db.Exec(createTableCommand); err != nil {
		log.Fatal("DB作成に失敗しました")
	}
}

func main() {
	for {
		displayAll()
		input()
	}
}

func displayAll() {
	fmt.Printf("\nデータを全件表示します\n\n")
	rows, err := db.Query(getPhonesCommand)
	if err != nil {
		log.Fatal("データ取得に失敗しました")
	}
	for rows.Next() {
		var phone Phone
		if err := rows.Scan(&phone.Id, &phone.Name, &phone.Number); err != nil {
			log.Fatal("データ取得に失敗しました")
		}
		fmt.Printf("名前：%s、電話番号：%s\n", phone.Name, phone.Number)
	}
	rows.Close()
}

func input() {
	fmt.Printf("\n「名前 電話番号」のフォーマットで入力してください\n")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	slicedInput := strings.Split(input, " ")
	if len(slicedInput) != 2 {
		log.Fatal("「名前 電話番号」のフォーマットで入力してください")
	}
	fmt.Println(slicedInput)
	db.Exec(insertPhonesCommand, slicedInput[0], slicedInput[1])
	fmt.Println("登録完了!")
}
