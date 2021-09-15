package users

import (
    "MopShopGo/repository"
    "errors"
    "fmt"
    "log"
)

func getAllUsers() ([]UserDetails, error) {
    result, err1 := repository.GetMysql().Query("" +
        "SELECT * " +
        "FROM users")
    defer result.Close()

    if err1 != nil {
        log.Printf("Cannot read data from database. %s", err1.Error())

        return []UserDetails{}, err1
    }

    var users []UserDetails
    for result.Next() {
        var userDatabase UserDatabase
        err2 := result.Scan(&userDatabase.Id, &userDatabase.Email, &userDatabase.PasswordHash, &userDatabase.Name, &userDatabase.Address)

        if err2 != nil {
            log.Printf("Exception %s ", err2.Error())

            return []UserDetails{}, err2
        }

        users = append(users, mapUserDatabaseToUseridDetails(userDatabase))
    }

    return users, nil
}

func getUser(userId int) (UserDetails, error) {
    result, err1 := repository.GetMysql().Query(""+
        "SELECT * "+
        "FROM users "+
        "WHERE id = ?", userId)
    defer result.Close()

    if err1 != nil {
        log.Printf("Cannot read data from database. %s", err1.Error())

        return UserDetails{}, err1
    }

    if result.Next() {
        var userDatabase UserDatabase
        err2 := result.Scan(&userDatabase.Id, &userDatabase.Email, &userDatabase.PasswordHash, &userDatabase.Name, &userDatabase.Address)

        if err2 != nil {
            log.Printf("Exception %s ", err2.Error())

            return UserDetails{}, err2
        }

        return mapUserDatabaseToUseridDetails(userDatabase), nil
    }

    return UserDetails{}, errors.New(fmt.Sprintf("Cannot find a user with id %s.", userId))
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

func mapUserDatabaseToUseridDetails(userDatabase UserDatabase) UserDetails {
    return UserDetails{
        Id:      userDatabase.Id,
        Email:   userDatabase.Email,
        Name:    userDatabase.Name,
        Address: userDatabase.Address,
    }
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

func getUserHash(email string) (string, error) {
    result, err := repository.
        GetMysql().
        Query(""+
            "SELECT passwordHash "+
            "FROM users "+
            "WHERE email = ?",
            email)

    if err != nil {
        log.Printf("User reading error %s.", err.Error())
        return "", err
    }

    if result.Next() {
        var passwordHash string
        _ = result.Scan(&passwordHash)

        return passwordHash, nil
    }

    return "", errors.New("cannot find hash for email " + email)
}