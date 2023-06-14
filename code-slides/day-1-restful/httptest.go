package main

// from https://github.com/Soypete/golang-cli-game/blob/main/server/api_test.go
func TestPassGetUsernameEmpty(t *testing.T) {
	...
	sPass.Router = setupTestRouter(sPass, t)
	w := httptest.NewRecorder()
	reqNoUser := httptest.NewRequest("GET", "/register//get", nil)
	reqNoUser.Header.Set("Authorization", "Basic Y2FwdGFpbm5vYm9keTE6cGFzc3dvcmQK")
	sPass.Router.ServeHTTP(w, reqNoUser)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}ackage main
