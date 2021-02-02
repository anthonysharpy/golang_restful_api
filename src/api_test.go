/* Copyright Anthony Sharp 2021 */

package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {

	fmt.Printf("Performing TestCreateUser\n")

	if(!InitialiseDatabase()) {
		t.Errorf("ERROR INITIALISING DATABASE. STOPPING.\n")
	}

	defer CloseDatabase()

	if(DoesUserExist("test1")) {
		RemoveUser("test1")
	}

	CreateUser("Test", "Testington", "test1", true)

	if(!DoesUserExist("test1")) {
		t.Errorf("TestCreateUser FAILED: failed to create user with username test1!\n")
	}

	RemoveUser("test1")
	
	fmt.Printf("Success.\n")
}

func TestDeleteUser(t *testing.T) {

	fmt.Printf("Performing TestDeleteUser\n")

	if(!InitialiseDatabase()) {
		t.Errorf("ERROR INITIALISING DATABASE. STOPPING\n")
	}
	
	defer CloseDatabase()

	CreateUser("Test", "Testington", "test1", true)
	RemoveUser("test1")

	if(DoesUserExist("test1")) {
		t.Errorf("TestDeleteUser FAILED: failed to delete user with username test1!\n")
	}
	
	fmt.Printf("Success.\n")
}

func TestNameChange(t *testing.T) {

	fmt.Printf("Performing TestNameChange\n")

	if(!InitialiseDatabase()) {
		t.Errorf("ERROR INITIALISING DATABASE. STOPPING.\n")
	}
	
	defer CloseDatabase()

	if(DoesUserExist("test1")) {
		RemoveUser("test1")
	}
	if(DoesUserExist("test2")) {
		RemoveUser("test2")
	}

	CreateUser("Test", "Testington", "test1", true)
	SetUserName("test1", "test2")

	if(DoesUserExist("test1")) {
		t.Errorf("TestNameChange FAILED: failed to rename user test1!\n")
	}
	if(!DoesUserExist("test2")) {
		t.Errorf("TestNameChange FAILED: failed to rename user test1 to test2!\n")
	} else {
		RemoveUser("test2")
	}
	
	fmt.Printf("Success.\n")
}

func TestSetDarkMode(t *testing.T) {

	fmt.Printf("Performing TestSetDarkMode\n")

	if(!InitialiseDatabase()) {
		t.Errorf("ERROR INITIALISING DATABASE. STOPPING.\n")
	}
	
	defer CloseDatabase()

	CreateUser("Test", "Testington", "test1", true)

	SetDarkModeForUser("test1", true)

	if(!GetDarkModeForUser("test1")) {
		t.Errorf("TestSetDarkMode FAILED: failed to turn dark mode on for user test1!\n")
	}

	SetDarkModeForUser("test1", false)

	if(GetDarkModeForUser("test1")) {
		t.Errorf("TestSetDarkMode FAILED: failed to turn dark mode off for user test1!\n")
	}

	RemoveUser("test1")

	fmt.Printf("Success.\n")
}

func TestListUsers(t *testing.T) {
	
	fmt.Printf("Performing TestListUsers\n")
	
	if(!InitialiseDatabase()) {
		t.Errorf("ERROR INITIALISING DATABASE. STOPPING.\n")
	}
	
	defer CloseDatabase()

	var numberofusers int = GetNumberOfUsers()
	var userlist = ListUsers()

	if(strings.Count(userlist, `"Darkmodeenabled": `) < numberofusers) {
		t.Errorf("TestListUsers FAILED: didn't return all users!\n")
	}

	fmt.Printf("Success.\n")
}

func TestSearchUsers(t *testing.T) {
	
	fmt.Printf("Performing TestSearchUsers\n")

	if(!InitialiseDatabase()) {
		t.Errorf("ERROR INITIALISING DATABASE. STOPPING.\n")
	}
	
	defer CloseDatabase()

	CreateUser("Test", "Testington", "test1", true)
	CreateUser("Tester", "Testing", "test2", true)
	CreateUser("Tested", "Testem", "test3", true)

	var usersearch = SearchUsers("test")

	if(strings.Count(usersearch, `"test`) < 3)	{
		t.Errorf("TestSearchUsers FAILED: expected at least three accounts!\n")
	}

	usersearch = SearchUsers("swdfrv435i25")

	if(strings.Count(usersearch, `"test`) > 0)	{
		t.Errorf("TestSearchUsers FAILED: didn't expect to find an account called swdfrv435i25!\n")
	}

	RemoveUser("test1")
	RemoveUser("test2")
	RemoveUser("test3")

	fmt.Printf("Success.\n")
}

