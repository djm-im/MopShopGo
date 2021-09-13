package users

import (
    "MopShopGo/repository"
    "log"
)

func getAllUsers() ([]UserDetails, error) {
    result, err := repository.GetMysql().Query("" +
        "SELECT * " +
        "FROM users")
    defer result.Close()

    if err != nil {
        log.Printf("Cannot read data from database. %s", err.Error())

        return []UserDetails{}, err
    }

    var users []UserDetails
    for result.Next() {
        var userDatabase UserDatabase
        err := result.Scan(&userDatabase.Id, &userDatabase.Email, &userDatabase.PasswordHash, &userDatabase.Name, &userDatabase.Address)

        if err != nil {
            log.Printf("Exception %s ", err.Error())

            return []UserDetails{}, err
        }

        users = append(users, mapUserDatabaseToUseridDetails(userDatabase))
    }

    return users, nil
}

func mapUserDatabaseToUseridDetails(userDatabase UserDatabase) UserDetails {
    return UserDetails{
        Id:      userDatabase.Id,
        Email:   userDatabase.Email,
        Name:    userDatabase.Name,
        Address: userDatabase.Address,
    }
}

func signup(userDatabase UserDatabase) (UserDetails, error) {
    result, err := repository.GetMysql().Exec(""+
        "INSERT INTO users "+
        "   (email, passwordHash)"+
        "   VALUE"+
        "   (?, ?)",
        userDatabase.Email, userDatabase.PasswordHash,
    )

    if err != nil {
        return UserDetails{}, err
    }

    id, _ := result.LastInsertId()
    userDatabase.Id = id

    return mapUserDatabaseToUseridDetails(userDatabase), nil
}

func doesExistUser(email string) bool {
    result, err := repository.GetMysql().Query(""+
        "SELECT COUNT(*) "+
        "FROM users "+
        "WHERE email = ?",
        email)
    defer result.Close()

    if err != nil {
        log.Printf("User reading error %s.", err.Error())
        return false
    }

    if result.Next() {
        var count int
        _ = result.Scan(&count)

        return count == 1
    }

    return false
}