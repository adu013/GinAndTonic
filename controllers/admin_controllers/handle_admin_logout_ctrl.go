package admincontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleLogout clears out the active login cookies
func HandleAdminLogout(c *gin.Context) {
	c.SetCookie("admin_session", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin/login")
}
