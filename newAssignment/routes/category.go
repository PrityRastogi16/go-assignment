package routes

import (
	"fmt"
	"net/http"
	"newAssignment/db"
	"newAssignment/models"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// func createCategory(c *gin.Context) {
// 	var category models.Category
// 	if err := c.ShouldBindJSON(&category); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
// 		return
// 	}
// 	if result := db.DB.Create(&category); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "category": category})
// }

func createCategory(c *gin.Context) {
	// Parse form data including files
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse form"})
		return
	}

	// Get file from form
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to get file"})
		return
	}
	defer file.Close()

	// Create a directory to save the image if it doesn't exist
	imageDir := "./uploads/"
	if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create image directory"})
		return
	}

	// Save the file
	filePath := filepath.Join(imageDir, filepath.Base(fileHeader.Filename))

	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save file"})
		return
	}
	defer out.Close()

	if _, err := out.ReadFrom(file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save file"})
		return
	}

	// Construct image URL or path
	imageURL := fmt.Sprintf("image%s", filepath.Base(filePath))

	// Bind category data
	var category models.Category
	fmt.Println(category)
	if err := c.ShouldBind(&category); err != nil {
		fmt.Println("___", category)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Assign image URL to category
	category.Image = imageURL

	// Save category to database
	if result := db.DB.Create(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "category": category})
}

// func createCategory(c *gin.Context) {
// 	// Parse form data (including file)
// 	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB limit
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid form data"})
// 		return
// 	}

// 	// Extract category name
// 	name := c.Request.FormValue("name")
// 	if name == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Category name is required"})
// 		return
// 	}

// 	// Handle file upload
// 	file, _, err := c.Request.FormFile("image")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "No file is received"})
// 		return
// 	}

// 	// Save the file
// 	fileName := time.Now().Format("20060102150405") + filepath.Ext(c.Request.MultipartForm.File["image"][0].Filename)
// 	filePath := filepath.Join("uploads", fileName)

// 	out, err := os.Create(filePath)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save the file"})
// 		return
// 	}
// 	defer out.Close()

// 	_, err = file.Copy(out)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save the file"})
// 		return
// 	}

// 	// Create category
// 	category := Category{Name: name, Image: filePath}
// 	if result := db.DB.Create(&category); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "category": category})
// }

func listCategories(c *gin.Context) {
	var categories []models.Category
	if result := db.DB.Find(&categories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// Delete Category
func deleteCategory(c *gin.Context) {
	id := c.Param("id")
	if result := db.DB.Delete(&models.Category{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
