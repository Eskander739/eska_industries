package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"path/filepath"
	"time"
)

const adminId = "5f5130c4-74ef-3f13-af4c-0b5137a36fe8"

func mailInfoForReg(toUser string, bodyUUID4 string) {
	from := "ataxxonnext@yandex.com"

	user := "ataxxonnext@yandex.com"
	password := "qpxdchfajfcsmyhj"

	to := []string{
		toUser,
	}
	//"imap": "imap.yandex.ru", "smtp": "smtp.yandex.ru", "imap_port": 993, "smtp_port": 587
	addr := "smtp.yandex.ru:587"
	host := "smtp.yandex.ru"

	msg := []byte("From: ataxxonnext@yandex.com\r\n" +
		"To:" + toUser + "\r\n" +
		"Subject: Регистрация на сайте ESKA\r\n\r\n" + // Заголовок
		"Чтобы завершить регистрацию пройдите по ссылке: http://localhost:8000/new-user/" + bodyUUID4 + "\r\n") // Тело сообщения

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}

func contactMail(User string, bodyUUID4 string, UserEmail string) {
	from := "ataxxonnext@yandex.com"

	user := "ataxxonnext@yandex.com"
	password := "qpxdchfajfcsmyhj"

	to := []string{
		"lastrogue222@gmail.com",
	}
	//"imap": "imap.yandex.ru", "smtp": "smtp.yandex.ru", "imap_port": 993, "smtp_port": 587
	addr := "smtp.yandex.ru:587"
	host := "smtp.yandex.ru"

	msg := []byte("From: ataxxonnext@yandex.com\r\n" +
		"To:" + User + "\r\n" +
		"Subject: Привет, меня зовут " + User + "\r\n\r\n" + // Заголовок
		"Моя почта для связи -" + UserEmail + "\n\n" + bodyUUID4 + "\r\n") // Тело сообщения

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}
func RandomValue() string {

	p, _ := rand.Prime(rand.Reader, 64)
	var value = fmt.Sprintf("%s", p)
	return value[:4]
}

func changeEmail(UserEmail string) string {
	var code = RandomValue()

	go func() {
		from := "ataxxonnext@yandex.com"

		user := "ataxxonnext@yandex.com"
		password := "qpxdchfajfcsmyhj"

		to := []string{
			UserEmail,
		}
		//"imap": "imap.yandex.ru", "smtp": "smtp.yandex.ru", "imap_port": 993, "smtp_port": 587
		addr := "smtp.yandex.ru:587"
		host := "smtp.yandex.ru"

		msg := []byte("From: ataxxonnext@yandex.com\r\n" +
			"To:" + UserEmail + "\r\n" +
			"Subject: Смена почты \r\n\r\n" + // Заголовок
			"Код для смены почты -  " + code) // Тело сообщения

		auth := smtp.PlainAuth("", user, password, host)

		err := smtp.SendMail(addr, auth, from, to, msg)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Email sent successfully")
	}()

	return code
}

func returnS(name string, w http.ResponseWriter) {
	//указываем путь к нужному файлу
	path := filepath.Join("static", name)
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//выводим шаблон клиенту в браузер
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

type pageData struct {
	IdForAdd string
}

func readPostReturn(name string, w http.ResponseWriter, data pageData) {
	//указываем путь к нужному файлу
	path := filepath.Join("static", name)
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//выводим шаблон клиенту в браузер
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func returnSEdit(name string, w http.ResponseWriter, data AddData) {
	//указываем путь к нужному файлу
	path := filepath.Join("static", name)
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//w.Write()
	//выводим шаблон клиенту в браузер
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

type NamePassword struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type IdImage struct {
	Id string `json:"id"`
}

type AddData struct {
	TitleForAdd   string
	ContentForAdd string
	IdForAdd      string
}

//	func ImageData(imagee chan []byte, imagee2 chan []byte) []byte {
//		fmt.Println("Запущена горутина: ", 1)
//		for {
//			imagee <- <-imagee2
//			fmt.Println("Работает горутина номер ", 1)
//
//		}
//
// }
//
//	func verifyUserPass(username, password string) bool {
//		var usersPasswords = map[string][]byte{
//			"joe":  []byte("$2a$12$aMfFQpGSiPiYkekov7LOsu63pZFaWzmlfm1T8lvG6JFj2Bh4SZPWS"),
//			"mary": []byte("$2a$12$l398tX477zeEBP6Se0mAv.ZLR8.LZZehuDgbtw2yoQeMjIyCNCsRW"),
//		}
//
//		wantPass, hasUser := usersPasswords[username]
//		if !hasUser {
//			return false
//		}
//		if cmperr := bcrypt.CompareHashAndPassword(wantPass, []byte(password)); cmperr == nil {
//			return true
//		}
//		return false
//	}
func MD5Hash(text string) string {
	/*
		Генерирует хэш из строки
	*/

	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	//var imagee chan []byte
	//var imagee2 chan []byte

	var editData = func(w http.ResponseWriter, r *http.Request) {

		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == adminId {
				postId := mux.Vars(r)["post-id"]

				if postId != "" {
					var postParams = map[string]string{"Title": "TEXT", "Content": "TEXT"}

					var tables = map[string]map[string]string{"posts": postParams}
					var posts = Db{DbName: "requests", TableName: "posts", FetchInfo: "posts", Tables: tables}

					var data1, _ = posts.fetchInfo()
					var dataToSend PostData

					for _, info := range data1 {
						if info.(PostData).Id == postId {
							fmt.Println("Мы нашли данные по ID: ", info.(PostData))
							dataToSend.Title = info.(PostData).Title
							dataToSend.Content = info.(PostData).Content

						}
					}

					var data = AddData{TitleForAdd: dataToSend.Title, ContentForAdd: dataToSend.Content, IdForAdd: postId}
					fmt.Println("r.URL: ", postId)
					returnSEdit("edit_post.html", w, data)
				}

			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var postList = func(w http.ResponseWriter, r *http.Request) {
		var posts = Db{DbName: "requests", TableName: "posts", FetchInfo: "posts"}
		var data1, _ = posts.fetchInfo()
		fmt.Println("Информация отправленная ИЗ БД: ", data1)
		var data = map[string]any{"posts": data1}
		var dataPosts, jsonError = json.MarshalIndent(data, "", "   ")
		if jsonError != nil {
			panic(jsonError)
		}
		_, err := w.Write(dataPosts)
		if err != nil {
			panic(err)
		}

	}

	var getImageById = func(w http.ResponseWriter, r *http.Request) {
		decoder, errReadAll := ioutil.ReadAll(r.Body)
		if errReadAll != nil {
			panic(errReadAll)
		}
		var data IdImage

		err := json.Unmarshal(decoder, &data)
		if err != nil {
			panic(err)
		}
		fmt.Println("data.Id: ", data.Id)
		var imageId = ImageServ{Id: data.Id}

		var images = Db{DbName: "requests", TableName: "post_images", FetchInfo: "post_images", ImageS: imageId}
		var data1, err2 = images.getImageById()
		if err2 != nil {
			panic(err2)
		}
		_, err = w.Write(data1)
		if err != nil {
			panic(err)
		}

	}

	var getPostById = func(w http.ResponseWriter, r *http.Request) {
		decoder, errReadAll := ioutil.ReadAll(r.Body)
		if errReadAll != nil {
			panic(errReadAll)
		}
		var data PostData

		err := json.Unmarshal(decoder, &data)
		if err != nil {
			panic(err)
		}
		fmt.Println("data.Id: ", data.Id)
		var imageId = ImageServ{Id: data.Id}

		var images = Db{DbName: "requests", TableName: "posts", FetchInfo: "posts", ImageS: imageId, PostD: data}
		var data1, err2 = images.getPostById()
		if err2 != nil {
			panic(err2)
		}
		_, err = w.Write(data1)
		if err != nil {
			panic(err)
		}

	}

	var saveChanges = func(w http.ResponseWriter, r *http.Request) {

		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == adminId {

				decoder, errReadAll := ioutil.ReadAll(r.Body)
				if errReadAll != nil {
					panic(errReadAll)
				}

				var data PostData

				err := json.Unmarshal(decoder, &data)
				if err != nil {
					panic(err)
				}
				fmt.Println(data.Title, data.Content)
				var postParams = map[string]string{"Title": "TEXT", "Content": "TEXT"}

				var tables = map[string]map[string]string{"posts": postParams}
				var post = PostData{Id: data.Id, Title: data.Title, Content: data.Content}

				var posts = Db{DbName: "requests", TableName: "posts", PostD: post, FetchInfo: "posts", Tables: tables}
				err = posts.ChangePost()
				if err != nil {
					panic(err)
				}
				var data1, _ = posts.fetchInfo()
				fmt.Println("Информация ИЗМЕНЕННАЯ в БД: ", data1)
			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var deletePost = func(w http.ResponseWriter, r *http.Request) {

		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == adminId {
				postId := mux.Vars(r)["post-id"]
				fmt.Println("IDDDDDDDDDDDDDD: ", postId)

				var post = PostData{Id: postId}

				var posts = Db{DbName: "requests", TableName: "posts", PostD: post, FetchInfo: "posts"}
				_, err := posts.removeInfo()
				if err != nil {
					panic(err)
				}

				var imageS = ImageServ{Id: postId}
				var images = Db{DbName: "requests", TableName: "post_images", ImageS: imageS}
				_, err = images.removeInfoImage()
				if err != nil {
					panic(err)
				}
				http.Redirect(w, r, "http://localhost:8000/admin/posts", http.StatusSeeOther)
			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var addPostData = func(w http.ResponseWriter, r *http.Request) {

		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == adminId {
				_, data, imageErr := r.FormFile("image")
				if imageErr != nil {
					panic(imageErr)
				}
				fileContent, err := data.Open()
				if err != nil {
					panic(err)
				}
				image, err := ioutil.ReadAll(fileContent)
				if err != nil {
					panic(err)
				}
				fmt.Println(" data.Image: ", r.FormValue("title"))
				fmt.Println(" data.Image: ", r.FormValue("content"))

				var uuidSQL = uuid4SQL()

				date := fmt.Sprintln(time.Now().Date())

				var post = PostData{Id: uuidSQL, Title: r.FormValue("title"), Content: r.FormValue("content"), Date: date}
				fmt.Println("POST DATAAAAA :", post.Id, post.Title, post.Content, post.Date)

				var imageData = ImageServ{Id: uuidSQL, Image: image}
				var posts = Db{DbName: "requests", TableName: "posts", PostD: post, FetchInfo: "posts", ImageS: imageData}
				err = posts.AddPost()
				if err != nil {
					panic(err)
				}

				err = posts.AddImage()
				if err != nil {
					panic(err)
				}

				var data1, _ = posts.fetchInfo()

				for _, dart := range data1 {
					fmt.Println("Информация добавленная в БД: ", dart.(PostData).Id, dart.(PostData).Title, dart.(PostData).Content, dart.(PostData).Date)
				}

				returnS("add_post.html", w)
			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var add_post = func(w http.ResponseWriter, r *http.Request) {

		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == adminId {
				returnS("add_post.html", w)
			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}
	}

	var posts = func(w http.ResponseWriter, r *http.Request) {

		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == adminId {
				returnS("all_posts.html", w)
			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var postPage = func(w http.ResponseWriter, r *http.Request) {
		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == adminId {
				returnS("posts.html", w)
			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var blog = func(w http.ResponseWriter, r *http.Request) {
		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == adminId {
				returnS("blog_management.html", w)
			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var style_guide = func(w http.ResponseWriter, r *http.Request) {
		returnS("style-guide.html", w)

	}

	var index = func(w http.ResponseWriter, r *http.Request) {
		returnS("index.html", w)

	}

	var about = func(w http.ResponseWriter, r *http.Request) {
		returnS("about.html", w)

	}
	var contact = func(w http.ResponseWriter, r *http.Request) {
		returnS("contact.html", w)

	}

	var category = func(w http.ResponseWriter, r *http.Request) {
		returnS("category.html", w)

	}

	var auth = func(w http.ResponseWriter, r *http.Request) {
		returnS("auth.html", w)

	}

	var sendMail = func(w http.ResponseWriter, r *http.Request) {
		var name = r.FormValue("cName")
		var email = r.FormValue("cEmail")
		var message = r.FormValue("cMessage")
		fmt.Println("name: ", name, "email:", email, "message: ", message)
		contactMail(name, message, email)
		returnS("contact.html", w)
	}

	var data = func(w http.ResponseWriter, r *http.Request) {
		decoder, errReadAll := ioutil.ReadAll(r.Body)
		if errReadAll != nil {
			panic(errReadAll)
		}

		var data NamePassword

		err := json.Unmarshal(decoder, &data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data.Email, data.Password, data.ConfirmPassword)
		if data.Password == data.ConfirmPassword {

			userId := uuid4SQL()

			var usersAddForDb = Db{DbName: "requests", TableName: "users", FetchInfo: "users",
				UserD: UserData{Id: userId, Email: data.Email, Password: MD5Hash(data.Password)}}

			errAddUser := usersAddForDb.AddUser()
			if errAddUser != nil {
				panic(errAddUser)
			}

			mailInfoForReg(data.Email, userId)

			timer1 := time.NewTimer(600 * time.Second)

			go func() {
				<-timer1.C

				var users = Db{DbName: "requests", TableName: "users", FetchInfo: "users"}
				usersList, err := users.Users()
				if err != nil {
					panic(err)
				}

				for _, dart := range usersList {
					fmt.Println("Пользователь: ", dart.(UserData).Username, "Id пользователя: ", dart.(UserData).Id)
					if userId == dart.(UserData).Id && dart.(UserData).Username == "unknown" {
						users.UserD = UserData{Id: dart.(UserData).Id}
						_, err2 := users.removeUser()
						if err2 != nil {
							panic(err2)
						}

					}
				}

				fmt.Println("Таймер закончил свою работу")
			}()

		}

	}

	var readPost = func(w http.ResponseWriter, r *http.Request) {
		postId := mux.Vars(r)["post-id"]

		if postId != "" {
			fmt.Println("strPost: ", postId)
			var localData = pageData{IdForAdd: postId}

			readPostReturn("read_post.html", w, localData)
		}

	}

	// verifyUserPass проверяет, что имя пользователя и пароль соответствуют друг другу,
	// сверяясь с нашей "базой данных" userPasswords.

	//var Reg = func(w http.ResponseWriter, req *http.Request) {
	//	user, pass, ok := req.BasicAuth()
	//	if ok && verifyUserPass(user, pass) {
	//		fmt.Fprintf(w, "You get to see the secret\n")
	//	} else {
	//		w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
	//		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	//	}
	//}

	var registration = func(w http.ResponseWriter, r *http.Request) {
		returnS("reg.html", w)

	}
	//var timerForReg = func(w http.ResponseWriter, r *http.Request) {
	//	/*
	//		Если в течение 10 минут пользователь не перейдет по ссылке для регистрации, ссылка станет неактуальной
	//	*/
	//
	//	userId := mux.Vars(r)["user-id"]
	//	timer1 := time.NewTimer(600 * time.Second)
	//
	//	go func() {
	//		<-timer1.C
	//
	//		var users = Db{DbName: "requests", TableName: "users", FetchInfo: "users"}
	//		usersList, err := users.Users()
	//		if err != nil {
	//			panic(err)
	//		}
	//		for _, dart := range usersList {
	//			fmt.Println("Пользователь: ", dart.(UserData).Username, "Id пользователя: ", dart.(UserData).Id)
	//			if userId == dart.(UserData).Id && dart.(UserData).Username == "unknown" {
	//				users.UserD = UserData{Id: dart.(UserData).Id}
	//				_, err2 := users.removeUser()
	//				if err2 != nil {
	//					panic(err2)
	//				}
	//
	//			}
	//		}
	//
	//		fmt.Println("Таймер закончил свою работу")
	//	}()
	//
	//}

	var approveUser = func(w http.ResponseWriter, r *http.Request) {
		/*
			Если в течение 10 минут пользователь не перейдет по ссылке для регистрации, ссылка станет неактуальной
		*/

		userId := mux.Vars(r)["user-id"]

		var usersCheck = Db{DbName: "requests", TableName: "users", FetchInfo: "users"}

		usersListCheck, err2 := usersCheck.Users()
		if err2 != nil {
			panic(err2)
		}

		for _, user := range usersListCheck {
			if user.(UserData).Id == userId && user.(UserData).Username == "unknown" {
				var users = Db{DbName: "requests", TableName: "users", FetchInfo: "users", UserD: UserData{Id: userId, Username: uuid4SQL()}}
				err := users.ChangeUsername()
				if err != nil {
					panic(err)
				}

				usersList, err2 := users.Users()
				if err2 != nil {
					panic(err2)
				}
				fmt.Println("Изменение unknown на имя формата uuid4:", usersList)
				//cookie1 := &http.Cookie{Name: "EskaUser", Value: userId}
				cookie1 := http.Cookie{Name: "EskaUser", Value: userId, Expires: time.Now().Add(time.Hour), HttpOnly: false, MaxAge: 50000, Path: "/"}
				http.SetCookie(w, &cookie1)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var getUser = func(w http.ResponseWriter, r *http.Request) {

		var cookie, err = r.Cookie("EskaUser")
		fmt.Println("EERRRRR: ", err)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
		} else {
			fmt.Println("Cookie - 0 : ", cookie.Value)
			fmt.Println("Cookie: ", cookie)

			if cookie.Value == "5f5130c4-74ef-3f13-af4c-0b5137a36fe8" {
				userId := mux.Vars(r)["user-id"]

				var users = Db{DbName: "requests", TableName: "users", FetchInfo: "users"}

				usersList, err2 := users.Users()
				if err2 != nil {
					panic(err2)
				}

				for _, dart := range usersList {
					fmt.Println("Пользователь: ПРОСТО ПРОВЕРКА", dart.(UserData).Username, "Id пользователя: ", dart.(UserData).Id)
					if dart.(UserData).Id == userId {
						var dataPosts, jsonError = json.MarshalIndent(dart.(UserData), "", "   ")
						if jsonError != nil {
							panic(jsonError)
						}

						_, errWrite := w.Write(dataPosts)

						if errWrite != nil {
							panic(errWrite)
						}

					}

				}
			} else {
				w.WriteHeader(403)
				http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)

			}

		}

	}

	var logIn = func(w http.ResponseWriter, r *http.Request) {
		/*
			Если в течение 10 минут пользователь не перейдет по ссылке для регистрации, ссылка станет неактуальной
		*/

		decoder, errReadAll := ioutil.ReadAll(r.Body)
		if errReadAll != nil {
			panic(errReadAll)
		}

		var data NamePassword

		err := json.Unmarshal(decoder, &data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data.Email, data.Password, data.ConfirmPassword)

		var users = Db{DbName: "requests", TableName: "users", FetchInfo: "users"}

		usersList, err2 := users.Users()
		if err2 != nil {
			panic(err2)
		}
		for _, dart := range usersList {
			fmt.Println("Пользователь LOGIN: ", dart.(UserData).Username, "Id пользователя: ", dart.(UserData).Id)
			if dart.(UserData).Password == MD5Hash(data.Password) && dart.(UserData).Email == data.Email {
				//if dart.(UserData).Password == data.Password && dart.(UserData).Email == data.Email {
				cookie1 := http.Cookie{Name: "EskaUser", Value: dart.(UserData).Id, Expires: time.Now().Add(time.Hour),
					HttpOnly: false, MaxAge: 50000, Path: "/"}
				http.SetCookie(w, &cookie1)

			}
		}

	}

	var account = func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://localhost:8000/account/about-me", http.StatusSeeOther)

	}

	var aboutMe = func(w http.ResponseWriter, r *http.Request) {
		returnS("about-me.html", w)

	}
	var security = func(w http.ResponseWriter, r *http.Request) {
		returnS("security.html", w)

	}
	var support = func(w http.ResponseWriter, r *http.Request) {
		returnS("support.html", w)

	}

	var changeName = func(w http.ResponseWriter, r *http.Request) {
		/*
			Если в течение 10 минут пользователь не перейдет по ссылке для регистрации, ссылка станет неактуальной
		*/

		userId := mux.Vars(r)["user-id"]
		newName := mux.Vars(r)["new-name"]

		var usersCheck = Db{DbName: "requests", TableName: "users", FetchInfo: "users"}

		usersListCheck, err2 := usersCheck.Users()
		if err2 != nil {
			panic(err2)
		}

		for _, user := range usersListCheck {
			if user.(UserData).Id == userId {
				var users = Db{DbName: "requests", TableName: "users", FetchInfo: "users", UserD: UserData{Id: userId, Username: newName}}
				err := users.ChangeUsername()
				if err != nil {
					panic(err)
				}

				usersList, err2 := users.Users()
				if err2 != nil {
					panic(err2)
				}
				fmt.Println("Изменение unknown на имя формата uuid4:", usersList)
				http.Redirect(w, r, "http://localhost:8000/account/about-me", http.StatusSeeOther)

			}

		}

	}

	var changeEmailApprove = func(w http.ResponseWriter, r *http.Request) {
		/*
			Если в течение 10 минут пользователь не перейдет по ссылке для регистрации, ссылка станет неактуальной
		*/

		userId := mux.Vars(r)["user-id"]
		newEmail := mux.Vars(r)["new-email"]

		var usersCheck = Db{DbName: "requests", TableName: "users", FetchInfo: "users"}

		usersListCheck, err2 := usersCheck.Users()
		if err2 != nil {
			panic(err2)
		}

		for _, user := range usersListCheck {
			if user.(UserData).Id == userId {
				var users = Db{DbName: "requests", TableName: "users", FetchInfo: "users", UserD: UserData{Id: userId, Email: newEmail}}
				err := users.ChangeEmail()
				if err != nil {
					panic(err)
				}

				usersList, err2 := users.Users()
				if err2 != nil {
					panic(err2)
				}
				fmt.Println("Изменение ПОЧТЫЫЫЫ: ", usersList)
				http.Redirect(w, r, "http://localhost:8000/account/about-me", http.StatusSeeOther)

			}

		}

	}

	var changeEmaill = func(w http.ResponseWriter, r *http.Request) {
		/*
			Если в течение 10 минут пользователь не перейдет по ссылке для регистрации, ссылка станет неактуальной
		*/

		userEmail := mux.Vars(r)["current-email"]
		fmt.Println("userEmail: ", userEmail)
		var code = changeEmail(userEmail)
		var data1 = map[string]string{"code": code}
		var dataCode, _ = json.Marshal(data1)
		_, err := w.Write(dataCode)
		fmt.Println("dataCode: ", dataCode)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(200)

	}

	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css/"))))          // ПОДКЛЮЧАЕМ CSS ФАЙЛЫ
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./static/img/"))))          // ПОДКЛЮЧАЕМ ИЗОБРАЖЕНИЯ ФАЙЛЫ
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js/"))))             // ПОДКЛЮЧАЕМ JS ФАЙЛЫ
	router.PathPrefix("/vendor/").Handler(http.StripPrefix("/vendor/", http.FileServer(http.Dir("./static/vendor/")))) // ПОДКЛЮЧАЕМ VENDOR ФАЙЛЫ
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./static/fonts/"))))    // ПОДКЛЮЧАЕМ FONTS ФАЙЛЫ

	//router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))           // ПОДКЛЮЧАЕМ CSS ФАЙЛЫ
	//router.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./static/img"))))           // ПОДКЛЮЧАЕМ ИЗОБРАЖЕНИЯ ФАЙЛЫ
	//router.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js/"))))             // ПОДКЛЮЧАЕМ JS ФАЙЛЫ
	//router.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("./static/vendor/")))) // ПОДКЛЮЧАЕМ VENDOR ФАЙЛЫ
	//router.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./static/fonts/"))))    // ПОДКЛЮЧАЕМ FONTS ФАЙЛЫ

	// ***********************************ADMIN SECTION*********************************************
	router.HandleFunc("/admin", blog)                               // защита есть
	router.HandleFunc("/admin/post-page", postPage)                 // защита есть
	router.HandleFunc("/admin/new-post-page", add_post)             // защита есть
	router.HandleFunc("/admin/new-post", addPostData)               // защита есть
	router.HandleFunc("/admin/posts", posts)                        // защита есть
	router.HandleFunc("/post-list", postList)                       // защита есть
	router.HandleFunc("/admin/posts/{post-id}/changed", editData)   // защита есть
	router.HandleFunc("/admin/modified-post", saveChanges)          // защита есть
	router.HandleFunc("/admin/posts/{post-id}/deleted", deletePost) // защита есть

	// ***********************************ADMIN SECTION*********************************************
	router.HandleFunc("/send-mail", sendMail)

	router.HandleFunc("/auth", auth)
	router.HandleFunc("/reg", data)
	router.HandleFunc("/auth-user", logIn)
	router.HandleFunc("/", index)
	router.HandleFunc("/about", about)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/category", category)
	router.HandleFunc("/style-guide", style_guide)
	router.HandleFunc("/reading-page/{post-id}", readPost)

	router.HandleFunc("/get-image-by-id", getImageById)
	router.HandleFunc("/get-post-by-id", getPostById)
	//router.HandleFunc("/start_timer", timerForReg)

	router.HandleFunc("/registration", registration)
	router.HandleFunc("/new-user/{user-id}", approveUser)
	router.HandleFunc("/user/{user-id}", getUser)
	router.HandleFunc("/account", account)
	router.HandleFunc("/account/about-me", aboutMe)
	router.HandleFunc("/account/security", security)
	router.HandleFunc("/account/support", support)
	router.HandleFunc("/account/about-me/change-name/{user-id}/{new-name}", changeName)

	router.HandleFunc("/account/about-me/change-email/{current-email}", changeEmaill)
	router.HandleFunc("/account/about-me/change-email/approve/{user-id}/{new-email}", changeEmailApprove)
	listenError := http.ListenAndServe(":8000", router)

	if listenError != nil {
		panic(listenError)
	}
}
