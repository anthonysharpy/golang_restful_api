/* Copyright Anthony Sharp 2021 */

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Change as necessary.
var database_username = "admin"
var database_password = "secretpassword"

var database *sql.DB

type User struct {
	Firstname string
	Lastname string
	Username string
	Timejoined int64
	Darkmodeenabled int64
}

func ListUsers() string {

	var err error
	var results *sql.Rows

	results, err = database.Query("SELECT * FROM users")

	if (err != nil) {
		fmt.Printf("Failed to list users: "+err.Error() + "\n")
		return ""
	}

	auser := User{}
	allusers := []User{}
	
    for (results.Next()) {


		err = results.Scan(&auser.Firstname,
						&auser.Lastname,
						&auser.Username,
						&auser.Timejoined,
						&auser.Darkmodeenabled)

        if err != nil {
			fmt.Printf("Error whilst listing users: "+err.Error() + "\n")
			return ""
		}
		
        allusers = append(allusers, auser)
    }

	// If this were the real world, I would probably not indent the JSON, since whitespace = data = bandwidth = $$$.
	jsonarr, err := json.MarshalIndent(allusers, "", "	")

	if err != nil {
		fmt.Printf("Error whilst building json list: "+err.Error() + "\n")
		return ""
	}

	return string(jsonarr)
}

func SearchUsers(term string) string {

	var err error
	var results *sql.Rows

	term = "%"+term+"%"

	// THIS IS VERY INEFFICIENT FOR LARGE TABLES. BUT IT IS OKAY FOR THIS EXAMPLE.
	results, err = database.Query(`SELECT * FROM users WHERE (username LIKE ? 
															 OR firstname LIKE ? 
															 OR lastname LIKE ?)`, term, term, term)

	if (err != nil) {
		fmt.Printf("Failed to search users: "+err.Error() + "\n")
		return ""
	}

	auser := User{}
	allusers := []User{}
	
    for (results.Next()) {

		err = results.Scan(&auser.Firstname,
						&auser.Lastname,
						&auser.Username,
						&auser.Timejoined,
						&auser.Darkmodeenabled)

        if err != nil {
			fmt.Printf("Error when searching users: "+err.Error() + "\n")
			return ""
		}
		
        allusers = append(allusers, auser)
    }

	jsonarr, err := json.MarshalIndent(allusers, "", "	")

	if err != nil {
		fmt.Printf("Error when building json search: "+err.Error() + "\n")
		return ""
	}

	return string(jsonarr)
}

func CreateUser(firstname string, lastname string, username string, testing bool) bool {

	// Don't allow people to use these usernames because they are reserved for testing purposes.
	if(!testing && (username == "test1" || username == "test2" || username == "test3")) {
		return false
	}

	// Each person must have a unique username.
	if(DoesUserExist(username)) {
		return false
	}

	var err error

    _, err = database.Exec("INSERT INTO users VALUES (?, ?, ?, ?, 1)", firstname, lastname, username, strconv.FormatInt(time.Now().Unix(), 10))

    if (err != nil) {
		fmt.Printf("WARNING: "+err.Error() + "\n")
		return false
    }

	return true
}

func SetUserName(oldusername string, newusername string) bool {

	var err error

	_, err = database.Exec("UPDATE users SET username=? WHERE username=?", newusername, oldusername)

	if (err != nil) {
		fmt.Printf("WARNING: "+err.Error() + "\n")
		return false
	}

	return true
}

func SetDarkModeForUser(username string, on bool) bool {

	var err error

	_, err = database.Exec("UPDATE users SET darkmode=? WHERE username=?", on, username)

	if (err != nil) {
		fmt.Printf("WARNING: "+err.Error() + "\n")
		return false
	}

	return true
}

func RemoveUser(username string) bool {

	var err error

	_, err = database.Exec("DELETE FROM users WHERE username=?", username)

	if (err != nil) {
		fmt.Printf("WARNING: "+err.Error() + "\n")
		return false
	}

	return true
}

func DoesUserExist(username string) bool {
	
	var err error
	var count int

	err = database.QueryRow("SELECT COUNT(*) FROM users WHERE username=?", username).Scan(&count)

	if(err != nil) {
		fmt.Printf("Error seeing if user exists: " + err.Error() + "\n")
		return false
	}

	if(count == 1) {
		return true
	}

	return false
}

func GetDarkModeForUser(username string) bool {
	
	var err error
	var on int

	err = database.QueryRow("SELECT darkmode FROM users WHERE username=?", username).Scan(&on)

	if(err != nil) {
		fmt.Printf("Error seeing if user has dark mode enabled: " + err.Error() + "\n")
		return true
	}

	if(on == 0) {
		return false
	}
	
	return true
}

func GetNumberOfUsers() int {

	var count int = 0
	var err error

	err = database.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	
	if (err != nil) {
		fmt.Printf("WARNING: "+err.Error() + "\n")
	}

	return count
}

func CreateDatabase() bool {

	var err error

	fmt.Printf("Creating database for the first time.\n")

	database, err = sql.Open("mysql", database_username+":"+database_password+"@tcp(127.0.0.1:3306)/")

	_,err = database.Exec("CREATE DATABASE JDIAPIDB")
	if err != nil {
		fmt.Printf("Couldn't create database: " + err.Error() + "\n")
		return false
	}

	_,err = database.Exec("USE JDIAPIDB")
	if err != nil {
		fmt.Printf("Couldn't use JDIAPIDB after creation: "+err.Error() + "\n")
		return false
	}

	_,err = database.Exec(`CREATE TABLE users (firstname varchar(50),
												lastname varchar(50),
												username varchar(20),
												timejoined bigint,
												darkmode boolean)`)

	if err != nil {
		fmt.Printf("Couldn't create users table: " + err.Error() + "\n")
		return false
	}

	fmt.Printf("Database creation went OK. Now adding default users.\n")

	if(!CreateUser("Donna", "Sandford", "D594MxM", false) || 
	!CreateUser("Judy", "Sheindlin", "JudgeJudy41", false) || 
	!CreateUser("Tasha", "Croad", "CCxCTash", false) ||
	!CreateUser("Steven", "Bradley", "BigBrad4", false) || 
	!CreateUser("Phillip", "Mitchell", "ProGamer79", false)) {
		return false
	}

	fmt.Printf("Created new database and added 5 users.\n")

	return true
}

func DoesDatabaseExist() bool {

	var err error
	var count int

	err = database.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

	if(err != nil) {
		if(strings.Contains(err.Error(), "Unknown database")) {
			return false
		} else
		{
			fmt.Printf("Unspecified error: " + err.Error() + "\n")
			return true
		}
	}

	return true
}

func InitialiseDatabase() bool {

	var err error

	database, err = sql.Open("mysql", database_username+":"+database_password+"@tcp(127.0.0.1:3306)/JDIAPIDB")

	if(!DoesDatabaseExist()) {
		CreateDatabase()
	} else {
		fmt.Printf("Database exists and was loaded successfully.\n")
	}

	database, err = sql.Open("mysql", database_username+":"+database_password+"@tcp(127.0.0.1:3306)/JDIAPIDB")

    if (err != nil) {
		fmt.Printf("Couldn't open database: " + err.Error() + "\n")
		return false
	}

	_,err = database.Exec("USE JDIAPIDB")
	if err != nil {
		fmt.Printf("Couldn't use database JDIAPIDB: "+err.Error() + "\n")
		return false
	}

	return true
}

func CloseDatabase() {
    database.Close()
}