package main

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const path = "C:\\Users\\eskander.baisiev\\GolandProjects\\awesomeProject\\eska_holdings\\eska.db"

type PostData struct {
	Id      string
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type UserData struct {
	Id       string
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type ImageServ struct {
	Id    string
	Image []byte `json:"image"`
}

type Db struct {
	Tables    map[string]map[string]string
	DbName    string
	DbPath    string
	PostD     PostData
	TableName string
	FetchInfo string
	ImageS    ImageServ
	UserD     UserData
}

//type Post struct {
//	Id      string
//	Title   string `json:"title"`
//	Content string `json:"content"`
//}

func uuid4SQL() string {
	/*
		Генератор уникальных id
	*/

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",

		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func (c Db) AddPost() error {
	/*
		Добавляет пост в БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}

	var ErrorAddInfo error

	records := `INSERT INTO posts(Id, Title, Content, Date) VALUES (?, ?, ?, ?)`
	query, prepareError := db.Prepare(records)
	if prepareError != nil {
		ErrorAddInfo = prepareError
	}

	_, execError := query.Exec(c.PostD.Id, c.PostD.Title, c.PostD.Content, c.PostD.Date)
	if execError != nil {
		ErrorAddInfo = execError
	}
	return ErrorAddInfo
}

func (c Db) AddUser() error {
	/*
		Добавляет пост в БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}

	var ErrorAddInfo error

	records := `INSERT INTO users(Id, Email, Password, Username) VALUES (?, ?, ?, ?)`
	query, prepareError := db.Prepare(records)
	if prepareError != nil {
		ErrorAddInfo = prepareError
	}

	_, execError := query.Exec(c.UserD.Id, c.UserD.Email, c.UserD.Password, "unknown") //c.UserD.Username
	if execError != nil {
		ErrorAddInfo = execError
	}

	return ErrorAddInfo
}

func (c Db) Users() ([]any, error) {
	/*
		Выкачивает всю инфу из БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}
	var results []any
	record, queryError := db.Query("SELECT * FROM " + c.TableName)

	if queryError != nil {
		return nil, queryError
	}

	defer func(record *sql.Rows) {
		err := record.Close()
		if err != nil {
			panic(err)
		}

	}(record)

	if c.FetchInfo == "users" {
		for record.Next() {
			var Id string
			var Email string
			var Password string
			var Username string
			scanError := record.Scan(&Id, &Email, &Password, &Username)

			if scanError != nil {
				return results, scanError
			}

			var user = UserData{Id: Id, Email: Email, Password: Password, Username: Username}

			results = append(results, user)
		}

		return results, nil

	}
	return nil, nil

}

func (c Db) removeUser() (bool, error) {
	/*
		Удаляет данные из БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		return false, nil
	}

	var IdDb = c.UserD.Id
	fmt.Println("АЙДЫШНИК: ", IdDb)
	var deleteReq = fmt.Sprintf("DELETE FROM " + c.TableName + " WHERE Id = '" + IdDb + "'")
	fmt.Println("ЗАПРОС : ", deleteReq)
	_, execError := db.Exec(deleteReq)

	if execError != nil {
		return false, nil
	}

	return true, nil

}

func (c Db) ChangeUsername() error {
	/*
		Добавляет пост в БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}

	var ErrorAddInfo error

	records := `UPDATE users SET Username = ?WHERE Id = ?`
	_, execError := db.Exec(records, c.UserD.Username, c.UserD.Id)

	if execError != nil {
		ErrorAddInfo = execError
	}
	return ErrorAddInfo
}

func (c Db) AddImage() error {
	/*
		Добавляет пост в БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}

	var ErrorAddInfo error

	records := `INSERT INTO post_images(Id, Image) VALUES (?, ?)`
	query, prepareError := db.Prepare(records)
	if prepareError != nil {
		ErrorAddInfo = prepareError
	}

	_, execError := query.Exec(c.ImageS.Id, c.ImageS.Image)
	if execError != nil {
		ErrorAddInfo = execError
	}
	return ErrorAddInfo
}

func (c Db) ChangePost() error {
	/*
		Добавляет пост в БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}

	var ErrorAddInfo error

	records := `UPDATE posts SET Title = ?, Content = ? WHERE Id = ?`
	_, execError := db.Exec(records, c.PostD.Title, c.PostD.Content, c.PostD.Id)

	if execError != nil {
		ErrorAddInfo = execError
	}
	return ErrorAddInfo
}

func (c Db) AddPostRequest() error {
	/*
		Добавляет запрос в БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}

	var ErrorAddInfo error

	records := `INSERT INTO posts(Id, Number, Code) VALUES (?, ?, ?)`
	query, prepareError := db.Prepare(records)
	if prepareError != nil {
		ErrorAddInfo = prepareError
	}

	_, execError := query.Exec(c.PostD.Id, c.PostD.Title, c.PostD.Content)

	if execError != nil {
		ErrorAddInfo = execError
	}
	return ErrorAddInfo
}

func (c Db) removeInfo() (bool, error) {
	/*
		Удаляет данные из БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		return false, nil
	}

	var IdDb = c.PostD.Id
	fmt.Println("АЙДЫШНИК: ", IdDb)
	var deleteReq = fmt.Sprintf("DELETE FROM " + c.TableName + " WHERE Id = '" + IdDb + "'")
	fmt.Println("ЗАПРОС : ", deleteReq)
	_, execError := db.Exec(deleteReq)

	if execError != nil {
		return false, nil
	}

	return true, nil

}

func (c Db) removeInfoImage() (bool, error) {
	/*
		Удаляет данные из БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		return false, nil
	}

	fmt.Println("АЙДЫШНИК IMAGE: ", c.ImageS.Id)
	var deleteReq = fmt.Sprintf("DELETE FROM " + c.TableName + " WHERE Id = '" + c.ImageS.Id + "'")
	fmt.Println("ЗАПРОС : ", deleteReq)
	_, execError := db.Exec(deleteReq)

	if execError != nil {
		return false, nil
	}

	return true, nil

}

func (c Db) removeRegInfo() (bool, error) {
	/*
		Удаляет данные из БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		return false, nil
	}

	var values, Err = c.fetchInfo()

	if Err != nil {
		return false, nil
	}

	for _, userReg := range values {
		fmt.Println(userReg.(PostData).Id == c.PostD.Id)
		if userReg.(PostData).Id == c.PostD.Id {
			var deleteReq = fmt.Sprintf("delete from " + c.TableName + " where Id = '" + c.PostD.Id + "'")
			fmt.Println(deleteReq)
			_, execError := db.Exec(deleteReq)

			if execError != nil {
				return false, nil
			}

			return true, nil

		}
	}

	return false, nil

}

func GetMD5Hash1(text string) string {
	/*
		Генерирует хэш из строки
	*/

	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (c Db) fetchInfo() ([]any, error) {
	/*
		Выкачивает всю инфу из БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}
	var results []any
	record, queryError := db.Query("SELECT * FROM " + c.TableName)

	if queryError != nil {
		return nil, queryError
	}

	defer func(record *sql.Rows) {
		err := record.Close()
		if err != nil {
			panic(err)
		}

	}(record)

	if c.FetchInfo == "posts" {
		for record.Next() {
			var Id string
			var Title string
			var Content string
			var Date string
			scanError := record.Scan(&Id, &Title, &Content, &Date)

			if scanError != nil {
				return results, scanError
			}

			var user = PostData{Id: Id, Title: Title, Content: Content, Date: Date}

			results = append(results, user)
		}

		return results, nil

	}
	return nil, nil

}

func (c Db) getImageById() ([]byte, error) {
	/*
		Выкачивает всю инфу из БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}
	fmt.Println("SELECT Image FROM " + c.TableName + "WHERE Id= '" + c.ImageS.Id + "'")
	record, queryError := db.Query("SELECT Image FROM " + c.TableName + " WHERE Id= '" + c.ImageS.Id + "'")

	if queryError != nil {
		return nil, queryError
	}

	defer func(record *sql.Rows) {
		err := record.Close()
		if err != nil {
			panic(err)
		}

	}(record)

	if c.FetchInfo == "post_images" {
		for record.Next() {
			var Image []byte
			scanError := record.Scan(&Image)

			return Image, scanError
		}

	}
	return nil, nil

}

func (c Db) getPostById() ([]byte, error) {
	/*
		Выкачивает всю инфу из БД
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}
	fmt.Println("SELECT * FROM " + c.TableName + "WHERE Id= '" + c.PostD.Id + "'")
	record, queryError := db.Query("SELECT * FROM " + c.TableName + " WHERE Id= '" + c.PostD.Id + "'")

	if queryError != nil {
		return nil, queryError
	}

	defer func(record *sql.Rows) {
		err := record.Close()
		if err != nil {
			panic(err)
		}

	}(record)

	if c.FetchInfo == "posts" {
		for record.Next() {
			var Id string
			var Title string
			var Content string
			var Date string
			scanError := record.Scan(&Id, &Title, &Content, &Date)

			var data = map[string]any{"title": Title, "content": Content, "id": Id, "date": Date}
			var dataPosts, jsonError = json.MarshalIndent(data, "", "   ")
			if jsonError != nil {
				panic(jsonError)
			}
			return dataPosts, scanError
		}

	}
	return nil, nil

}

func (c Db) createTable() error {
	/*
		Создает таблицу posts
	*/

	db, dateBaseError := sql.Open("sqlite3", c.DbName)
	if dateBaseError != nil {
		panic(dateBaseError)
	}

	var table string

	for tableName, Params := range c.Tables {
		var tableType = "CREATE TABLE IF NOT EXISTS "
		table = tableType + tableName + " (\"Id\" CHAR(50) NOT NULL,\n"
		for name, param := range Params {
			table = table + " \"" + name + "\" " + param + ",\n"
		}
		table = table[:len(table)-2]
		table = table + ");"
	}
	query, prepareError := db.Prepare(table)
	if prepareError != nil {
		return prepareError
	}
	_, execError := query.Exec()

	if execError != nil {
		return execError
	}

	return nil
}

//func main() {
//	//	//gh := uuid4()
//	//
//
//	var imageTable = map[string]string{"Image": "BLOB"}
//	var tables = map[string]map[string]string{"post_images": imageTable}
//	var ff1 = Db{DbName: "requests", TableName: "post_images", FetchInfo: "post_images", Tables: tables}
//	Err := ff1.createTable()
//	//
//	if Err != nil {
//		panic(Err)
//	}
//
//	var postsTable = map[string]string{"Title": "TEXT", "Content": "TEXT", "Date": "TEXT"}
//	var tables2 = map[string]map[string]string{"posts": postsTable}
//	var ff2 = Db{DbName: "requests", TableName: "posts", FetchInfo: "posts", Tables: tables2}
//	Err2 := ff2.createTable()
//	//
//	if Err2 != nil {
//		panic(Err2)
//	}
//
//	var usersTable = map[string]string{"Email": "TEXT", "Password": "TEXT", "Username": "TEXT"}
//	var tables3 = map[string]map[string]string{"users": usersTable}
//	var ff3 = Db{DbName: "requests", TableName: "users", FetchInfo: "users", Tables: tables3}
//	Err3 := ff3.createTable()
//	//
//	if Err3 != nil {
//		panic(Err3)
//	}
//
//	//Err := ff1.AddPost()
//	//if Err != nil {
//	//	panic(Err)
//	//}removeRegInfo
//	//ff1.removeRegInfo()
//	value, _ := ff1.fetchInfo()
//	fmt.Println(value)
//}

//var Err = strcct.createTable()
//
//if Err != nil {
//	panic(Err)
//}
//
//var res, Err1 = strcct.fetchRequests()
//
//if Err1 != nil {
//	panic(Err1)
//}
//fmt.Println("RES: ", res)
