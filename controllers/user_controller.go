package controllers

import (
	"net/http"
	"os"
	"time"
    "fmt"
	"github.com/vishnu/instabe/models"
	"github.com/vishnu/instabe/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user);
	err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err := services.CreateUser(&user);
	err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context){
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input);
	err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user , err := services.FindUserByEmail(input.Email)
	if err!=nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))!=nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "user_id": user.ID.String(),  // UUID should be string
    "exp":     time.Now().Add(time.Hour * 24).Unix(),
})

	
	
jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
if err != nil {
    fmt.Println("Error signing token:", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
    return
}

c.JSON(http.StatusOK, gin.H{"token": jwtToken})

}







func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}





func FollowUser(c *gin.Context) {
	userID := c.Param("id")
	followerID, exists := c.Get("userID") // From JWT
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := services.FollowUserService(userID, followerID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Followed successfully"})
}

func UnfollowUser(c *gin.Context) {
	userID := c.Param("id")
	followerID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	err := services.UnfollowUserService(userID, followerID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unfollowed successfully"})
}




func GetFollowers(c *gin.Context) {
	userID := c.Param("id")
	followers, err := services.GetFollowersService(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, followers)
}

func GetFollowing(c *gin.Context) {
	userID := c.Param("id")
	following, err := services.GetFollowingService(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, following)
}