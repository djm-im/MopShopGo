package users

import (
    "MopShopGo/exceptions"
    "MopShopGo/sessions"
    "encoding/json"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

func GetAllUsers() ([]UserDetails, error) {
    return getAllUsers()
}

func GetUser(userId int) (UserDetails, error) {
    return getUser(userId)
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
    decoder := json.NewDecoder(request.Body)

    var userCredentials UserCredentials
    err1 := decoder.Decode(&userCredentials)

    if err1 != nil {
        exceptions.BuildErrorResponse(&response, http.StatusNotAcceptable, err1.Error())
        return
    }

    email := userCredentials.Email
    password := userCredentials.Password

    exist := doesExistUser(email)
    if !exist {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusConflict,
            fmt.Sprintf("Cannot find a user with email %s address.", email))
        return
    }

    passwordHash, err2 := getUserHash(email)
    if err2 != nil {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusConflict,
            fmt.Sprintf("Cannot fetch password hash for %s address.", email))
        return
    }

    err3 := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
    if err3 != nil {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusUnauthorized,
            fmt.Sprintf("Bad combination of email (%s) address and password.", email))
        return
    }

    sessionToken := sessions.CreateSessionToken(email)

    response.Header().Set("Content-Type", "application/json")
    http.SetCookie(response, &http.Cookie{
        Name:    "session_token",
        Value:   sessionToken.Token,
        Expires: sessionToken.Expire,
    })
    response.WriteHeader(http.StatusNoContent)
}

func Logout(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}
