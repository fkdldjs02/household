package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/usjeong/household/api/model"
)

// SetUserContext 사용자 생성
func SetUserContext(c *gin.Context) {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		return
	}

	result, err := db.SetUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusConflict)
		log.Println(err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// AuthContext 사용자 인증
func AuthContext(c *gin.Context) {
	name := c.PostForm("name")
	passwd := c.PostForm("passwd")

	user, err := db.Auth(name, passwd)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	c.Set("authUser", name)
	c.JSON(http.StatusOK, user)
}

// AuthMiddleware 세션을 통한 인증 검증 미들웨어
func AuthMiddleware(c *gin.Context) {
	_, exist := c.Get("authUser")
	if exist == false {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Next()
}

// GetCategoryListContext 항목 리스트
func GetCategoryListContext(c *gin.Context) {
	categoryList, err := db.GetCategoryList()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"categoryList": categoryList})
}

// GetCategoryContext 항목 조회
func GetCategoryContext(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	category, err := db.GetCategory(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, category)
}

// SetCategoryContext 항목 생성
func SetCategoryContext(c *gin.Context) {
	category := model.Category{}
	if c.Bind(&category) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := db.SetCategory(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusConflict)
		log.Println(err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.AbortWithStatus(http.StatusConflict)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// PutCategoryContext 항목 수정
func PutCategoryContext(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	category := model.Category{}
	if err := c.Bind(&category); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		return
	}

	category.ID = id
	_, err = db.PutCategory(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.Status(http.StatusOK)
}

// DeleteCategoryContext 항목 제거
func DeleteCategoryContext(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err = db.DeleteCategory(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.Status(http.StatusOK)
}

// GetHouseholdListContext 가계부 리스트
func GetHouseholdListContext(c *gin.Context) {
	name := c.Param("name")
	householdList, err := db.GetHouseholdArray(name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"householdList": householdList})
}

// GetHouseholdContext 가계부 조회
func GetHouseholdContext(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	household, err := db.GetHousehold(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, household)
}

// SetHouseholdConext 가계부 생성
func SetHouseholdConext(c *gin.Context) {
	household := model.Household{}
	if c.Bind(&household) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := db.SetHousehold(&household)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.AbortWithStatus(http.StatusConflict)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// PutHouseholdContext 가계부 수정
func PutHouseholdContext(c *gin.Context) {
	household := model.Household{}
	if c.Bind(&household) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := db.PutHousehold(&household)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.Status(http.StatusOK)
}

// DeleteHouseholdContext 가계부 제거
func DeleteHouseholdContext(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err = db.DeleteHousehold(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.Status(http.StatusOK)
}
