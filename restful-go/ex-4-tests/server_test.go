package main

import (
	"net/http/httptest"
	"testing"
)

func TestUserendpoints(t *testing.T) {
	t.Run("Get Username: Pass", testPassGetUsername)
	t.Run("Get Username: no Username", testFailGetUsernameempty)
	t.Run("Update Username: Pass", testPassUpdateUser)
	t.Run("Update Username:Fail", testFailUpdateUsernameempty)
	t.Run("Delete Username: Pass", testPassDeleteUser)
	t.Run("Delete Username:Fail", testFailDeleteUsernameempty)
}

func testPassGetUsername(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/getUsername/captainnobody1", nil)
	getUsername(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func testFailGetUsernameempty(t *testing.T) {
	// TODO: implement
}

func testPassUpdateUser(t *testing.T) {
	// TODO: implement
}

func testFailUpdateUsernameempty(t *testing.T) {
	// TODO: implement
}

func testPassDeleteUser(t *testing.T) {
	// ToDO: implement
}

func testFailDeleteUsernameempty(t *testing.T) {
	// TODO: implement
}
