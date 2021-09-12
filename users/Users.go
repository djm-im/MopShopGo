package users

type UserCredentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserDetails struct {
    Id      int64  `json:"id"`
    Email   string `json:"email"`
    Name    string `json:"name"`
    Address string `json:"address"`
}

type UserDatabase struct {
    Id           int64
    Email        string
    PasswordHash string
    Name         string
    Address      string
}
