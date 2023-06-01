package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"skid_go/controllers"
)

func main() {
	dsn := "user:password@tcp(localhost:3306)/skid_videos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Create an instance of the UserController and pass the database connection
	userController := controllers.NewUserController(db)

	// Create a new Gin router
	router := gin.Default()

	// Define the routes for user authentication and token validation
	router.POST("/api/authenticate", userController.Authenticate)		     //done
	router.POST("/api/validateToken", userController.ValidateToken)              //done

	// Initialize the video controller with the database connection
	videoController := controllers.NewVideoController(db)

	// Define the routes for video management
	router.GET("/api/videos", videoController.GetVideos)                          //done
	router.GET("/api/videos/:videoId", videoController.GetVideoDetails)           //done
	router.POST("/api/videos", videoController.UploadVideoandEncryptURL)          //done
	router.DELETE("/api/videos/:videoId", videoController.DeleteVideo)            //done
	router.GET("/api/videos/:videoId/metadata", videoController.GetVideoMetadata) //done

	// Define the routes for adaptive bitrate streaming
	router.GET("/api/videos/:videoId/stream", videoController.StreamVideo) //done

	// Define the route for video decryption
	router.GET("/api/videos/:videoId/decrypt", videoController.DecryptVideo) //done

	// Define the route for video search
	router.GET("/api/videos/search", videoController.SearchVideos) //done
	router.Run(":8080")
}
