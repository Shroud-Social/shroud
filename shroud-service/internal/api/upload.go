package api

import (
	"fmt"
	"net/http"
	"services/internal/storage"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type config struct {
	UserUploadSecret []byte
	UploadUri        string
}

var conf config

func LoadUploadConf(envUploadSecret, envUploadUri string) {
	conf.UserUploadSecret = []byte(envUploadSecret)
	conf.UploadUri = envUploadUri
}

type UploadRequest struct {
	UserId     string             `json:"user_id" binding:"required"`
	ChannelId  string             `json:"channel_id"`
	GuildId    string             `json:"guild_id"`
	FileName   string             `json:"file_name" binding:"required"`
	Size       uint32             `json:"size" binding:"required"`
	Hash       string             `json:"hash" binding:"required"`
	UploadType storage.UploadType `json:"upload_type" binding:"required"`
}

type UploadReceipt struct {
	UploadRequest
	UploadId string `json:"upload_id" binding:"required"`
	Path     string `json:"path" binding:"required"`
}

func GetUploadToken(c *gin.Context) {
	var req UploadRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Binding Error": err.Error()})
		return
	}

	// TODO: User Auth/Permission Checks
	// TODO: Deduplication
	// TODO: Size Check/Fetch User Tier Upload Size Limit

	uploadId, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to generate UUID": err.Error()})
		return
	}

	replacer := strings.NewReplacer(
		"{user_id}", req.UserId,
		"{channel_id}", req.ChannelId,
		"{file_name}", fmt.Sprintf("%s-%s", uploadId, req.FileName),
	)

	uploadConfig, ok := storage.UploadConfigs[req.UploadType]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Upload type does not exist": req.UploadType})
	}
	path := replacer.Replace(uploadConfig.Path)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":         time.Now().Add(10 * time.Minute).Unix(),
		"path":        path,
		"hash":        req.Hash,
		"size":        req.Size,
		"upload_id":   uploadId.String(),
		"user_id":     req.UserId,
		"channel_id":  req.ChannelId,
		"guild_id":    req.GuildId,
		"file_name":   req.FileName,
		"upload_type": req.UploadType,
	})
	tokenString, err := token.SignedString(conf.UserUploadSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Couldn't generate JWT": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Authorization": tokenString,
		"uri":           fmt.Sprintf("%s%s", conf.UploadUri, path),
	})

}

func ProcessUpload(c *gin.Context) {
	var receipt UploadReceipt
	err := c.ShouldBindJSON(&receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Binding Error": err.Error()})
		return
	}

	// TODO: Commit Metadata to DB
}
