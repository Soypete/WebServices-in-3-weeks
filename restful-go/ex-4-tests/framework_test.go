package framework

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func setupTestRouter(t *testing.T) *chi.Mux {
	r := chi.NewRouter()

	// setup routes
	r.Route("/register", func(r chi.Router) {
		// subroutes for register
		r.Route("/{username}", func(r chi.Router) {
			r.Get("/get", getUsername)          // GET /register/123/get
			r.Get("/update", updateUsername)    // PUT /register/123/update // TODO: this is a get because I am not providing a body
			r.Delete("/delete", deleteUsername) // DELETE /register/123/delete
		})
	})
	return r
}

func TestUserEndpoints(t *testing.T) {
	t.Run("get username: Pass", testPassGetUserName)
	t.Run("get username: No Username", testFailGetUsernameEmpty)
	t.Run("update username: Pass", testPassUpdateUser)
	t.Run("update username:Fail", testFailUpdateUsernameEmpty)
	t.Run("delete username: Pass", testPassDeleteUser)
	t.Run("delete username:Fail", testFailDeleteUsernameEmpty)
}
func testPassGetUserName(t *testing.T) {
	router := setupTestRouter(t)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/register/captainnobody1/get", nil)
	router.ServeHTTP(w, req)

	// check status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// test no username
func testFailGetUsernameEmpty(t *testing.T) {
	router := setupTestRouter(t)
	w := httptest.NewRecorder()
	reqNoUser := httptest.NewRequest("GET", "/register//get", nil)
	router.ServeHTTP(w, reqNoUser)
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func testPassUpdateUser(t *testing.T) {
	router := setupTestRouter(t)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/register/captainnobody1/update", nil)
	router.ServeHTTP(w, req)

	// check status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// test no username
func testFailUpdateUsernameEmpty(t *testing.T) {
	router := setupTestRouter(t)
	w := httptest.NewRecorder()
	reqNoUser := httptest.NewRequest("GET", "/register//update", nil)
	router.ServeHTTP(w, reqNoUser)
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func testPassDeleteUser(t *testing.T) {
	router := setupTestRouter(t)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/register/captainnobody1/delete", nil)
	router.ServeHTTP(w, req)

	// check status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// test db error
func testFailDeleteUsernameEmpty(t *testing.T) {
	router := setupTestRouter(t)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/register//delete", nil)
	router.ServeHTTP(w, req)

	// check status code
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
