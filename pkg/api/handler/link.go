package handler

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"shotenLinker/pkg/domain"
	"time"

	"github.com/gin-gonic/gin"
	services "shotenLinker/pkg/service/interface"
)

type LinksHandler struct {
	linkUseCase services.LinkUseCase
}

func NewLinkHandler(useCase services.LinkUseCase) *LinksHandler {
	return &LinksHandler{
		linkUseCase: useCase,
	}
}

// Could be moved to special class Encoder
// We could be saved map to keys to decrease checks in DB
func encode(url string) string {
	hash := sha256.Sum256([]byte(url))
	short := base64.URLEncoding.EncodeToString(hash[:])[:8]

	return short
}

func (cr *LinksHandler) Encode(c *gin.Context) {
	var json struct {
		URL string `json:"url"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}

	_, err := url.ParseRequestURI(json.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL format"})
		return
	}

	appPort, exists := c.Get("AppPort")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AppPort not found"})
		return
	}

	appUrl, exists := c.Get("AppUrl")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AppUrl not found"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	existLink, err := cr.linkUseCase.FindByLink(ctx, json.URL)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			shortURL := encode(json.URL)
			address := fmt.Sprintf("%s:%s/%s", appUrl, appPort, shortURL)
			link := domain.Links{
				Link:      json.URL,
				ShortLink: address,
			}

			err = cr.linkUseCase.Save(c, link)
			c.JSON(200, gin.H{"shortened_url": link.ShortLink})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't find record"})
			return
		}
	} else {
		c.JSON(200, gin.H{"shortened_url": existLink.ShortLink})
	}
}

func (cr *LinksHandler) Decode(c *gin.Context) {
	shortURL := c.Param("shortURL")
	retrieveURL, err := cr.linkUseCase.FindByShortLink(c, shortURL)
	if err == nil {
		c.Redirect(302, retrieveURL.Link)
	} else {
		c.JSON(404, gin.H{"error": "URL not found"})
	}
}
