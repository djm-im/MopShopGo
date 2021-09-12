package users

import (
    "MopShopGo/exceptions"
    "encoding/json"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

func GetAllUsers() ([]UserDetails, error) {
    return getAllUsers()
}

func Signup(response http.ResponseWriter, request *http.Request) {
    decoder := json.NewDecoder(request.Body)

    var userCredentials UserCredentials
    err := decoder.Decode(&userCredentials)

    if err != nil {
        exceptions.BuildErrorResponse(&response, http.StatusNotAcceptable, err.Error())
        return
    }

    email := userCredentials.Email
    password := userCredentials.Password

    exist := doesExistUser(email)
    if exist {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusConflict,
            fmt.Sprintf("User with email %s address already exists.", email))
        return
    }

    var userDb UserDatabase
    passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
    userDb = UserDatabase{
        Email:        email,
        PasswordHash: string(passwordHash),
        Name:         "",
        Address:      "",
    }

    newUser, err := signup(userDb)

    if err != nil {
        exceptions.BuildErrorResponse(&response, http.StatusNotAcceptable, err.Error())
        return
    }

    response.Header().Set("Content-Type", "application/json")
    response.WriteHeader(http.StatusCreated)

    _ = json.NewEncoder(response).Encode(newUser)
}

func Login(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}

func Logout(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}
