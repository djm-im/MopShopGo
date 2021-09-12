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

func doesExistUser(email string) bool {
    // todo : add implementation
    return false
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
