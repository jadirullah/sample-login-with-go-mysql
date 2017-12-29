package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	// "os"
)

// deklarasi variable db end errors
var db *sql.DB
var err error

//deklarasi variabel untuk componen users
type user struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
}

//membuat function koneksi ke database mysql
func connect_db() {
	db, err = sql.Open("mysql", "root:jadir123@tcp(127.0.0.1:3306)/go_db")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

//membuat routes untuk alamat URL pada browser
func routes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
}

//fungsi utama yang akan di eksekusi
func main() {
	// conect db
	connect_db()
	// connect router untuk url
	routes()

	defer db.Close()
	//server aktif pada port :8000
	fmt.Println("Server is Actived in port :8000")
	http.ListenAndServe(":8000", nil)
}

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {

		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host+r.URL.Path, 301)
		return false
	}

	return true
}

// function query user untuk mengambil user berdasarkan userame nya
func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		first_name, 
		last_name, 
		password 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Password,
		)
	return users
}

//function untuk mengarahkan ke halaman home
func home(w http.ResponseWriter, r *http.Request) {
	//meng-chek ketersedian session
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	//properti untuk di set di html
	var data = map[string]string{
		"username": session.GetString("username"),
		"message":  "Selamat datang di menu utama",
	}

	var t, err = template.ParseFiles("views/home.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, data)
	return

}

//fungsi registers
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	users := QueryUser(username)

	fmt.Printf("%+v\n", (user{}))

	fmt.Printf("%+v\n", users)

	//perbandingan user yang di post (users{}) dengan user yang ada di database users
	if (user{}) == users {

		// user belum tersedia, boleh registers
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			stmt, err := db.Prepare("INSERT users SET username=?, password=?, first_name=?, last_name=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &first_name, &last_name)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}

	} else {
		//users sudah tersedia, gagak save
		http.Redirect(w, r, "/register", 302)

	}
}

//fungsi eksekusi login
func login(w http.ResponseWriter, r *http.Request) {
	//jika users sudah melakukan login dan masih aktif session nya
	//maka users tsb tidak usah login kembali
	session := sessions.Start(w, r)
	if len(session.GetString("username")) != 0 && checkErr(w, r, err) {
		http.Redirect(w, r, "/", 302)
	}

	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login.html")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)

	//deskripsi dan compare password
	var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if password_tes == nil {
		//login success
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		http.Redirect(w, r, "/", 302)
	} else {
		//login failed
		fmt.Println("Gagal Login")
		http.Redirect(w, r, "/login", 302)
	}
}

//fungsi logout
func logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	http.Redirect(w, r, "/", 302)
}
