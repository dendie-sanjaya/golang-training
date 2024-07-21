package service

import (
	"praisindo/entity"

	"github.com/redis/go-redis/v9"
)

type IShortUrlService interface {
	CreateShortUrl(p_urllong string) (entity.ShortUrl, error)
	GetShortUrl(p_shortUrl string) (entity.ShortUrl, error)
}

// IUserRepository mendefinisikan interface untuk repository pengguna

type IShortUrlRepository interface {
	CreateShortUrl(p_urllong string) (entity.ShortUrl, error)
	GetShortUrl(p_shortUrl string) (entity.ShortUrl, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type shortUrlService struct {
	shortUrlRepo IShortUrlRepository
	rdb          *redis.Client
}

// NewUserService membuat instance baru dari userService
func NewShortUrlService(shortUrlRepo IShortUrlRepository, rdb *redis.Client) IShortUrlService {
	return &shortUrlService{shortUrlRepo: shortUrlRepo, rdb: rdb}
}

// CreateUser membuat pengguna baru
func (s *shortUrlService) CreateShortUrl(p_urllong string) (entity.ShortUrl, error) {
	// Memanggil CreateUser dari repository untuk membuat pengguna baru
	// createdUser, err := s.userRepo.CreateUser(ctx, user)
	// if err != nil {
	// 	return entity.User{}, fmt.Errorf("gagal membuat pengguna: %v", err)
	// }

	// redisKey := fmt.Sprintf("createdUser:%d", createdUser.ID)
	// createdUserJSON, err := json.Marshal(createdUser)
	// if err != nil {
	// 	log.Println("gagal marshal json")
	// }
	// if err := s.rdb.Set(ctx, redisKey, createdUserJSON, 60*time.Second).Err(); err != nil {
	// 	log.Println("error when set redis key", redisKey)
	// }
	createdUser := entity.ShortUrl{}

	return createdUser, nil
}

// GetUserByID mendapatkan pengguna berdasarkan ID
func (s shortUrlService) GetShortUrl(p_shortUrl string) (entity.ShortUrl, error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	// id := ps_p_shortUrl

	// redisKey := fmt.Sprintf("createdUser:%d", id)
	// val, err := s.rdb.Get(redisKey).Result()
	// if err == nil {
	// 	log.Println("data tersedia di redis")
	// 	err = json.Unmarshal([]byte(val), &user)
	// 	if err != nil {
	// 		log.Println("gagal marshall data di redis, coba ambil ke database")
	// 	}
	// }

	// if err != nil {
	// 	log.Println("data tidak tersedia di redis, ambil dari database")
	// 	user, err = s.userRepo.GetUserByID(ctx, id)
	// 	if err != nil {
	// 		log.Println("gagal ambil data di database")
	// 		return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
	// 	}
	// }
	user := entity.ShortUrl{}

	return user, nil
}
