// config/settings.go
// Fine tune your application here

package config

// Set to true in production
const IsProductionEnv = false

// IsSecureHttp
const IsSecureHttp = false

// Tune your bCrypt cost (10 or less is a security risk)
const BcryptCost = 11

// User pagination Page Size
const UserPagePagination = 50

// Login Redirect Path
const LoginRedirectPath = "/login"

// Logout Redirect Path
const LogoutRedirectPath = "/login"
