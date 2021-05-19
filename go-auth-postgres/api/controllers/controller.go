package controllers

import (
	"fmt"
	"go-authentication/api/db"
	"go-authentication/api/model"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

// Login is gonna create new template
func Login(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Login", nil)
}

// Success is the representation of success page
func Success(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Success", nil)
}

// LoginProcess is defining the login functionality
func LoginProcess(w http.ResponseWriter, r *http.Request) {
	db := db.OpenConnection()
	if r.Method == "POST" {
		emailID := r.FormValue("emailId")
		password := r.FormValue("password")

		// Validate form input
		if strings.Trim(emailID, " ") == "" || strings.Trim(password, " ") == "" {
			fmt.Println("Parameter's can't be empty")
			http.Redirect(w, r, "/", 301)
			return
		}

		checkUser, err := db.Query("SELECT id,password_hash,first_name,last_name,email_id FROM users WHERE email_id=$1", emailID)

		if err != nil {
			panic(err.Error())
		}
		user := model.User{}
		for checkUser.Next() {
			var id int
			var password, firstName, lastName, emailID string
			err = checkUser.Scan(&id, &password, &firstName, &lastName, &emailID)
			if err != nil {
				panic(err.Error())
			}
			user.ID = id
			user.FirstName = firstName
			user.LastName = lastName
			user.EmailID = emailID
			user.Password = password
		}

		fmt.Println(user)

		errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			fmt.Println(errf)
			http.Redirect(w, r, "/", 301)
		} else {
			fmt.Println("Success")

			fmt.Println(user)
			tmpl.ExecuteTemplate(w, "Success", user)
		}

	}
	defer db.Close()

}

// Register is the represntation of list page
func Register(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Register", nil)
}

// RegisterUser is the representation of user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	db := db.OpenConnection()
	if r.Method == "POST" {

		firstName := r.FormValue("firstname")
		lastName := r.FormValue("lastname")
		EmailID := r.FormValue("emailId")
		mobileNumber := r.FormValue("mobilenumber")
		fmt.Println(EmailID)
		password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			tmpl.ExecuteTemplate(w, "Register", err)
		}
		fmt.Println(len(password))
		dt := time.Now()
		//Format YYYY-MM-DD
		createdDate := dt.Format("2006-01-02 15:04:05")
		//confirmPassword := r.FormValue("confirmpassword")
		insForm, err := db.Prepare("INSERT INTO users(first_name, last_name,email_id,mobile_number,password_hash,created_date) VALUES($1,$2,$3,$4,$5,$6)")
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(insForm)
		_, err = insForm.Exec(firstName, lastName, EmailID, mobileNumber, password, createdDate)
		log.Println("INSERT: First Name: " + firstName + " | Last Name: " + lastName)
		if err != nil {
			log.Fatalf("Something wrong", err)
		}

	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
