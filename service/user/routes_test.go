package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/spani73/go-ecommerce-api/service/user"
	"github.com/spani73/go-ecommerce-api/types"
)

func TestUserServiceHandlers(t *testing.T){
	userStore := &mockUserStore{}
	handler:= user.NewHandler(userStore)
	t.Run("should fail f the user payload is invalid" , func (t *testing.T)  {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName: "abc",
			Email: "asdadasd",
			Password: "asdasd",
		}
		marshalled, _ := json.Marshal(payload)
		req,err := http.NewRequest(http.MethodPost , "/register" , bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register" , handler.HandleRegister)
		router.ServeHTTP(rr,req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d , got %d", http.StatusBadRequest , rr.Code)
		}
	})
}

type mockUserStore struct {

}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User ,error){
	return nil,fmt.Errorf("user not found")
} 

func (m *mockUserStore) GetUserByID(id int ) (*types.User , error){
	return nil,nil
}

func (m *mockUserStore) CreateUser(types.User) (error){
	return nil
}