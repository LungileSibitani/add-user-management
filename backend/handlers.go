package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// GetUsers - fetch all users
func GetUsers(c *gin.Context) {
    rows, err := DB.Query(`
        SELECT id, username, first_name, last_name, email
        FROM users
        WHERE deleted_at IS NULL
    `)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        err := rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        users = append(users, u)
    }

    c.JSON(http.StatusOK, users)
}

// GetUserByID - fetch single user
func GetUserByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var u User
    err = DB.QueryRow(`
        SELECT id, username, first_name, last_name, email
        FROM users
        WHERE id=? AND deleted_at IS NULL
    `, id).Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, u)
}

// CreateUser - add new user
func CreateUser(c *gin.Context) {
    var u User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    result, err := DB.Exec(`
        INSERT INTO users (username, first_name, last_name, email, password_hash)
        VALUES (?, ?, ?, ?, ?)
    `, u.Username, u.FirstName, u.LastName, u.Email, u.PasswordHash)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    id, _ := result.LastInsertId()
    u.ID = int(id)

    c.JSON(http.StatusCreated, u)
}

// UpdateUser - modify user
func UpdateUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var u User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    _, err = DB.Exec(`
        UPDATE users
        SET username=?, first_name=?, last_name=?, email=?, updated_at=NOW()
        WHERE id=? AND deleted_at IS NULL
    `, u.Username, u.FirstName, u.LastName, u.Email, id)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    u.ID = id
    c.JSON(http.StatusOK, u)
}

// DeleteUser - soft delete user
func DeleteUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    _, err = DB.Exec(`
        UPDATE users
        SET deleted_at=NOW()
        WHERE id=?
    `, id)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

