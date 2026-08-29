package admincontrollers

import (
	"gin-and-tonic/config"
	"gin-and-tonic/database"
	"gin-and-tonic/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RenderAdminDashboard pulls all user records from SQLite and loads the dashboard grid view
func RenderAdminDashboard(c *gin.Context) {
	// Get query params
	searchQuery := c.Query("q")
	pageStr := c.Query("page")

	page := 1
	if p, err := strconv.Atoi(pageStr); err != nil && p > 0 {
		page = p
	}

	pageSize := config.UserPagePagination
	offset := (page - 1) * config.UserPagePagination

	// Build the query builder
	query := database.DB.Model(&models.User{})

	// If search query is provided, filter by email (case-insensitive via LIKE)
	if searchQuery != "" {
		query = query.Where("email LIKE ?", "%"+searchQuery+"%")
	}

	// Count total matching rows (needed to calculate total pages)
	var totalRows int64
	query.Count(&totalRows)

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	if totalPages == 0 {
		totalPages = 1
	}

	// Fetch the paginated records
	var users []models.User
	query.Limit(pageSize).Offset(offset).Order("id asc").Find(&users)

	// Helper arrays for template rendering
	var pageNumbers []int
	for i := 1; i <= totalPages; i++ {
		pageNumbers = append(pageNumbers, i)
	}

	c.HTML(http.StatusOK, "admin_dashboard.html", gin.H{
		"Title":       "Gin & Tonic",
		"Users":       users,
		"SearchQuery": searchQuery,
		"CurrentPage": page,
		"TotalPages":  totalPages,
		"PageNumbers": pageNumbers,
		"HasPrev":     page > 1,
		"HasNext":     page < totalPages,
		"PrevPage":    page - 1,
		"NextPage":    page + 1,
	})
}
