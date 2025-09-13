package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    // other imports
)

func main() {
 InitDB()
   r := gin.Default()

    // Add CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))

    // your routes
    r.GET("/api/users", GetUsers)         // fetch all users
r.GET("/api/users/:id", GetUserByID)  // fetch single user
r.POST("/api/users", CreateUser)      // create user
r.PUT("/api/users/:id", UpdateUser)   // update user
r.DELETE("/api/users/:id", DeleteUser) // delete user


    r.Run("0.0.0.0:8080")
}

