package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"praisindo/entity"

	"github.com/redis/go-redis/v9"
)

type IShortUrlService interface {
	CreateShortUrl(p_urllong string) (entity.ShortUrl, error)
	GetShortUrl(p_shortUrl string) (entity.ShortUrl, error)
}

type IShortUrlRepository interface {
	CreateShortUrl(p_urllong string) (entity.ShortUrl, error)
	GetShortUrl(p_shortUrl string) (entity.ShortUrl, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type shortUrlService struct {
	shortUrlRepo IShortUrlRepository
	rdb          *redis.Client
}

func NewShortUrlService(shortUrlRepo IShortUrlRepository, rdb *redis.Client) IShortUrlService {
	return &shortUrlService{shortUrlRepo: shortUrlRepo, rdb: rdb}
}

func (s *shortUrlService) CreateShortUrl(p_urllong string) (entity.ShortUrl, error) {

	createdShortUrl, err := s.shortUrlRepo.CreateShortUrl(p_urllong)
	if err != nil {
		log.Println("gagal create short url:", err)
		return entity.ShortUrl{}, err
	}

	redisKey := fmt.Sprintf("shorturl:%s", createdShortUrl.UrlShort)
	createdShortUrlJSON, err := json.Marshal(createdShortUrl)

	if err != nil {
		log.Println("gagal marshal json created short url")
	}

	if err := s.rdb.Set(context.Background(), redisKey, createdShortUrlJSON, 0).Err(); err != nil {
		log.Println("error when set redis key", redisKey)
	} else {
		log.Println("success set redis key", redisKey)
	}

	return createdShortUrl, nil
}

func (s shortUrlService) GetShortUrl(p_shortUrl string) (entity.ShortUrl, error) {
	id := p_shortUrl

	shortUrl := entity.ShortUrl{}
	redisKey := fmt.Sprintf("shorturl:%s", id)
	val, err := s.rdb.Get(context.Background(), redisKey).Result()
	if err == nil {
		log.Println("data tersedia di redis")
		err = json.Unmarshal([]byte(val), &shortUrl)
		if err != nil {
			log.Println("gagal marshall data di redis, coba ambil ke database")
		}
	}

	if err != nil {
		log.Println("data tidak tersedia di redis, ambil dari database")
		shortUrl, err = s.shortUrlRepo.GetShortUrl(id)
		if err != nil {
			log.Println("gagal ambil data di database")
			return entity.ShortUrl{}, fmt.Errorf("gagal mendapatkan long url berdasarkan shorturl: %v", err)
		}
	}

	return shortUrl, nil
}
