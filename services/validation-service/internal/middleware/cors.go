package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS returns a CORS middleware configured for the validation service
func CORS() gin.HandlerFunc {
	config := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			// Allow specific origins for frontend integration
			allowedOrigins := []string{
				"http://localhost:3000",  // Next.js development
				"http://localhost:3001",  // Alternative frontend port
				"http://localhost:8080",  // Transaction service
				"https://localhost:3000", // HTTPS development
			}

			for _, allowed := range allowedOrigins {
				if origin == allowed {
					return true
				}
			}

			// Allow Vercel and Netlify deployments
			return false // For now, be restrictive
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Accept",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return cors.New(config)
}