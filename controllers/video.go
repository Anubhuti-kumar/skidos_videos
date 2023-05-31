// controllers/video_controller.go

package controllers

import (
	"github.com/dgrijalva/jwt-go/v4"
	"net/http"
	"strconv"
	"time"

	"skid_go/models"

	"skid_go/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (uc *UserController) ValidateToken(ctx *gin.Context) {
	var tokenData struct {
		Token string `json:"token" binding:"required"`
	}

	// Bind the request body to the tokenData struct
	if err := ctx.ShouldBindJSON(&tokenData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	token, err := jwt.Parse(tokenData.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisismysecretkey"), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Token is valid, continue with the subsequent requests
	ctx.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}

type VideoController struct {
	db *gorm.DB
}

func NewVideoController(db *gorm.DB) *VideoController {
	return &VideoController{
		db: db,
	}
}

func (vc *VideoController) StreamVideo(ctx *gin.Context) {
	tokenData := struct {
		Token string `json:"token" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&tokenData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	token, err := jwt.Parse(tokenData.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisismysecretkey"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	// Retrieve the video ID from the request path parameters
	videoID := ctx.Param("videoId")

	// Parse the video ID to an integer
	id, err := strconv.Atoi(videoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	// Fetch the video with the given ID from the database
	var video models.Video
	err = vc.db.First(&video, id).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	// Perform adaptive bitrate streaming logic here
	// Generate different quality levels of the video and provide streaming URLs for each quality level
	streamURLs := generateStreamingURLs(video)

	ctx.JSON(http.StatusOK, streamURLs)
}

func generateStreamingURLs(video models.Video) map[string]string {
	// Generate different quality levels of the video
	streamURLs := make(map[string]string)

	// Example: Generate streaming URLs for different quality levels
	streamURLs["360p"] = generateStreamingURL(video.URL, "360p")
	streamURLs["720p"] = generateStreamingURL(video.URL, "720p")
	streamURLs["1080p"] = generateStreamingURL(video.URL, "1080p")

	return streamURLs
}

func generateStreamingURL(videoURL, quality string) string {
	// Perform logic to generate streaming URL for the given video URL and quality level
	// Example: Concatenate the video URL and quality level
	streamingURL := videoURL + "?quality=" + quality

	return streamingURL
}

func (vc *VideoController) GetVideos(ctx *gin.Context) {
	tokenData := struct {
		Token string `json:"token" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&tokenData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	token, err := jwt.Parse(tokenData.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisismysecretkey"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Token is valid, continue with retrieving videos
	var videos []models.Video
	err = vc.db.Find(&videos).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
		return
	}

	ctx.JSON(http.StatusOK, videos)
}

func (vc *VideoController) GetVideoDetails(ctx *gin.Context) {
	tokenData := struct {
		Token string `json:"token" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&tokenData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	token, err := jwt.Parse(tokenData.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisismysecretkey"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	// Retrieve the video ID from the request path parameters
	videoID := ctx.Param("videoId")

	// Parse the video ID to an integer
	id, err := strconv.Atoi(videoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	var video models.Video

	// Fetch the video with the given ID from the database
	err = vc.db.First(&video, id).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	ctx.JSON(http.StatusOK, video)
}

func (vc *VideoController) UploadVideoandEncryptURL(ctx *gin.Context) {

	var video models.Video

	// Retrieve the video details from the request body
	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video data"})
		return
	}

	// Set the upload date to the current time
	video.UploadDate = time.Now()

	encryptedURL := utils.EncryptDatas(video.URL)

	video.URL = encryptedURL

	// Save the video to the database
	err := vc.db.Create(&video).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload video"})
		return
	}

	ctx.JSON(http.StatusOK, video)
}
func (vc *VideoController) DeleteVideo(ctx *gin.Context) {
	tokenData := struct {
		Token string `json:"token" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&tokenData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	token, err := jwt.Parse(tokenData.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisismysecretkey"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	// Retrieve the video ID from the request path parameters
	videoID := ctx.Param("videoId")

	// Parse the video ID to an integer
	id, err := strconv.Atoi(videoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	// Delete the video with the given ID from the database
	err = vc.db.Delete(&models.Video{}, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video deleted"})
}

func (vc *VideoController) GetVideoMetadata(ctx *gin.Context) {
	tokenData := struct {
		Token string `json:"token" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&tokenData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	token, err := jwt.Parse(tokenData.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisismysecretkey"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	// Retrieve the video ID from the request path parameters
	videoID := ctx.Param("videoId")

	// Parse the video ID to an integer
	videoIDInt, err := strconv.Atoi(videoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	var video models.Video

	// Fetch the video with the given ID from the database
	err = vc.db.Where("id = ?", videoIDInt).First(&video).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	metadata := gin.H{
		"title":       video.Title,
		"resolution":  video.Resolution,
		"upload_date": video.UploadDate,
		"description": video.Description,
		"url":         video.URL,
	}

	ctx.JSON(http.StatusOK, metadata)
}

func (vc *VideoController) DecryptVideo(ctx *gin.Context) {
	// Retrieve the video ID from the request path parameters
	videoID := ctx.Param("videoId")

	// Parse the video ID to an integer
	id, err := strconv.Atoi(videoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	var video models.Video

	// Fetch the video with the given ID from the database
	err = vc.db.First(&video, id).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	// Decrypt the video URL using the DecryptDatas function from utils package
	decryptedURL, err := utils.DecryptDatas(video.URL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decrypt video URL"})
		return
	}

	response := gin.H{
		"id":          video.ID,
		"title":       video.Title,
		"description": video.Description,
		"duration":    video.Duration,
		"resolution":  video.Resolution,
		"upload_date": video.UploadDate,
		"url":         decryptedURL,
	}

	ctx.JSON(http.StatusOK, response)
}

func (vc *VideoController) SearchVideos(ctx *gin.Context) {
	tokenData := struct {
		Token string `json:"token" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&tokenData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	token, err := jwt.Parse(tokenData.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("thisismysecretkey"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	// Retrieve the search query from the request query parameter
	searchQuery := ctx.Query("query")

	// Fetch videos that match the search query from the database
	var videos []models.Video
	err = vc.db.Where("title LIKE ?", "%"+searchQuery+"%").Find(&videos).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search videos"})
		return
	}

	ctx.JSON(http.StatusOK, videos)
}
