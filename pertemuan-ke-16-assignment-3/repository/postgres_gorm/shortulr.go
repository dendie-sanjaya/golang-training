package postgres_gorm

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"praisindo/entity"
	"praisindo/service"

	"gorm.io/gorm"
)

type shortUrlRepository struct {
	db *gorm.DB
}

// func NewshortUrlRepository(db GormDBIface) service.IUserRepository {
func NewshortUrlRepository(db *gorm.DB) service.IShortUrlRepository {
	return &shortUrlRepository{db: db}
}

// func (r *shortUrl) CreateShorUrl(ctx context.Context, p_urllong string) (entity.ShortUrl, error) {
func (r *shortUrlRepository) CreateShortUrl(p_urllong string) (entity.ShortUrl, error) {

	fmt.Print("masuk ke create short url  gorm \n ")
	v_urlshort := generateShortCode(p_urllong)
	urlShort := entity.ShortUrl{
		UrlLong:  p_urllong,
		UrlShort: v_urlshort,
	}

	if err := r.db.Table("shorturl").Create(&urlShort).Error; err != nil {
		log.Printf("error creating shor url: %v\n", err)
		return entity.ShortUrl{}, err
	}

	return urlShort, nil
}

// func (r *shortUrlRepository) GetUShortUrl(ctx context.Context, id int) (entity.ShortUrl, error) {
func (r *shortUrlRepository) GetShortUrl(p_shortUrl string) (entity.ShortUrl, error) {
	//func (r *shortUrlRepository) GetUShortUrl(id int) (entity.ShortUrl, error) {
	fmt.Print("masuk ke get longsorturl gorm \n ")
	var urlShort entity.ShortUrl
	//if err := r.db.WithContext(ctx).Select("url_long", "url_short").First(&urlShort, id).Error; err != nil {
	if err := r.db.Table("shorturl").Select("url_long", "url_short").Where("url_short = ?", p_shortUrl).First(&urlShort).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ShortUrl{}, nil
		}
		log.Printf("Error getting shorturl ID: %v\n", err)
		return entity.ShortUrl{}, err
	}
	return urlShort, nil
}

func generateShortCode(url string) string {
	b := make([]byte, 6)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
