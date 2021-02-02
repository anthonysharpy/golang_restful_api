/* Copyright Anthony Sharp 2021 */

package main

import (
	"fmt"
	"net/http"
)

func RequestHandlerFunction(w http.ResponseWriter, r *http.Request) {

	switch(r.URL.Path) {
		case "/createuser":
			fmt.Printf("Got createuser request.\n")
			HandleCreateUserRequest(w, r)
			break
		case "/updatename":
			fmt.Printf("Got updatename request.\n")
			HandleUpdateNameRequest(w, r)
			break
		case "/setdarkmode":
			fmt.Printf("Got setdarkmode request.\n")
			HandleSetDarkModeRequest(w, r)
			break
		case "/deleteuser":
			fmt.Printf("Got deleteuser request.\n")
			HandleDeleteUserRequest(w, r)
			break
		case "/listusers":
			fmt.Printf("Got listusers request.\n")
			HandleListUsersRequest(w, r)
			break
		case "/searchusers":
			fmt.Printf("Got searchusers request.\n")
			HandleSearchUsersRequest(w, r)
			break
		default:
			http.Error(w, "404", http.StatusNotFound)
			break
	}
}

func IsFormValid(r *http.Request) bool {
	
	if (r.ParseForm() != nil) {
		return false
	}

	return true
}

/* Only allows POST requests for security reasons (GET requests leak information through the URL, which is sometimes a helpful thing, but not in this case). */
func IsPOSTRequest(r *http.Request) bool {

	if(r.Method == "POST") {
		return true
	}
	return false
}

func CorrectAuthcode(authcode string) bool {

	// IN THE REAL WORLD, THIS WOULD BE RANDOMLY GENERATED EVERY TIME THE ADMINISTRATOR LOGS IN AND STORED IN A DATABASE.
	if(authcode == "a764bcjd") {
		return true
	}
	return false
}

func HandleCreateUserRequest(w http.ResponseWriter, r *http.Request) {

	if(IsPOSTRequest(r) && IsFormValid(r)) {

		var firstname = r.FormValue("firstname")
		var lastname = r.FormValue("lastname")
		var username = r.FormValue("username")
		var authcode = r.FormValue("authcode")

		if(CorrectAuthcode(authcode)) {
			if(CreateUser(firstname, lastname, username, false)) {
				fmt.Fprintf(w, "1")
			} else {
				fmt.Fprintf(w, "0")
			}
		}
	}
}

func HandleUpdateNameRequest(w http.ResponseWriter, r *http.Request) {

	if(IsPOSTRequest(r) && IsFormValid(r)) {

		var oldusername = r.FormValue("oldusername")
		var newusername = r.FormValue("newusername")
		var authcode = r.FormValue("authcode")

		if(CorrectAuthcode(authcode)) {
			if(SetUserName(oldusername, newusername)) {
				fmt.Fprintf(w, "1")
			} else {
				fmt.Fprintf(w, "0")
			}
		}
	}
}

func HandleSetDarkModeRequest(w http.ResponseWriter, r *http.Request) {

	if(IsPOSTRequest(r) && IsFormValid(r)) {

		var username = r.FormValue("username")
		var on = r.FormValue("on")
		var authcode = r.FormValue("authcode")

		if(CorrectAuthcode(authcode)) {
			if(SetDarkModeForUser(username, !(on == "0"))) {
				fmt.Fprintf(w, "1")
			} else {
				fmt.Fprintf(w, "0")
			}
		}
	}
}

func HandleDeleteUserRequest(w http.ResponseWriter, r *http.Request) {
	
	if(IsPOSTRequest(r) && IsFormValid(r)) {

		var username = r.FormValue("username")
		var authcode = r.FormValue("authcode")

		if(CorrectAuthcode(authcode)) {
			if(RemoveUser(username)) {
				fmt.Fprintf(w, "1")
			} else {
				fmt.Fprintf(w, "0")
			}
		}
	}
}

func HandleListUsersRequest(w http.ResponseWriter, r *http.Request) {
	
	if(IsPOSTRequest(r) && IsFormValid(r)) {

		var authcode = r.FormValue("authcode")

		if(CorrectAuthcode(authcode)) {
			fmt.Fprintf(w, ListUsers())
		}
	}
}

func HandleSearchUsersRequest(w http.ResponseWriter, r *http.Request) {

	if(IsPOSTRequest(r) && IsFormValid(r)) {

		var term = r.FormValue("term")
		var authcode = r.FormValue("authcode")

		if(CorrectAuthcode(authcode)) {
			fmt.Fprintf(w, SearchUsers(term))
		}
	}
}