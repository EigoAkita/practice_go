package practiceSqlite

import (
	//コード不使用 => コンパイル時に一緒にビルド🙅‍♂️ => Sqliteアクセス不可
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func PracticeSqlite() {
	//【ドライバー名】 sqlite3 【開くsqlのファイル指定】 ./example.sql
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	//sqlのファイル開後、閉
	defer DbConnection.Close()
	//データベース作成(データベース作成コマンドの変数)
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	//cmd(コマンド)実行(コマンド実行後、データベースの情報非表示の為、エラーの有無限定の為、_ , err)
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		//データベースのエラーハンドリング
		log.Fatalln(err)
	}

	//personテーブルに対してデータをインサート(?, ?　<= データを後に渡可能)

	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// //上記(?, ?)した事で、"Nancy", 20 渡可能
	// _, err = DbConnection.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	/*
		上記コード記述後、"Nancy"と２０がテーブル内に代入確認
		P2020119noMacBook-Air:practice_go 3202$ sqlite3 example.sql
		SQLite version 3.36.0 2021-06-18 18:58:49
		Enter ".help" for usage hints.
		sqlite> select * from person;
		Nancy|20　　👈 コレ
	*/

	//テーブルアップデート

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 30, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//データ取得 ~マルチセレクト編~

	tableName := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	fmt.Println(cmd)
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		//Scanが Personのstructにデータ代入
		//1個1個エラー確認
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	// //纏めてエラー確認
	// err = rows.Err()
	// if err != nil {
	// //エラーの場合、プログラムをEXIT
	// 	log.Fatalln(err)
	// }
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	//データ取得 ~シングルセレクト編~

	// cmd = "SELECT * FROM person where name = ?"
	// row := DbConnection.QueryRow(cmd, "Nancy")
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		//テーブル内データ検索後、データ未登録時下記エラー発動
	// 		log.Println("Now row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	//データ削除

	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
