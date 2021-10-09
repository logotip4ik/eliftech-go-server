package controllers

import (
	"fmt"
	"home/work/models"
	"home/work/utils/token"

	"github.com/gin-gonic/gin"
)

func UpdateCurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	var input models.UpdateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
    return
  }

	var user models.User
	if err := models.DB.Table("users").Where("id = ?", userId).Take(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": fmt.Sprintf("user with id: \"%d\" was not found", userId)})
		return
	}

  if err := models.DB.Model(&user).Updates(models.User{Username: input.Username, Email: input.Email}).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var posts []models.Post
	
	if err := models.DB.Table("posts").Where("user_id = ?", userId).Order("created_at desc").Select("id", "title", "content", "created_at", "updated_at").Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return 
	}

	user.Username = input.Username
	user.Email = input.Email
	user.Posts = posts
	user.Password = ""
	
  c.JSON(200, gin.H{"data": user})

}

func DeleteCurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
  
	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
    c.JSON(404, gin.H{"error": err.Error()})
    return
  }

  if err := models.DB.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()});
		return
	}

  c.JSON(200, gin.H{"data": true})
}

func FindManyUsers(c *gin.Context) {
	var users []models.User

	if err := models.DB.Select("username", "email", "id").Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < len(users); i++ {
		users[i].Password = ""
	}

	c.JSON(200, gin.H{ "data": users })
}

func FindUserByUsername(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": fmt.Sprintf("user with username: \"%s\" not found", c.Param("username"))})
		return 
	}

	var posts []models.Post
	
	if err := models.DB.Table("posts").Where("user_id = ?", user.ID).Order("created_at desc").Select("title", "content", "created_at", "updated_at").Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return 
	}

	user.Posts = posts
	user.Password = ""

	c.JSON(200, gin.H{"data": user})
}

func GetCurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByID(userId)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func LoginUser(c *gin.Context) {
	var input models.LoginUserInput
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u := models.User{ Email: input.Email, Password: input.Password }
	token, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": "username or password is incorrect"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func RegisterUser(c *gin.Context) {
	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u := models.User{Username: input.Username, Email: input.Email, Password: input.Password}

	if _, err := u.SaveUser(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token, err := token.Generate(u.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": "something went wrong D_D"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
