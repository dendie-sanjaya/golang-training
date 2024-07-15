package slice

import (
	"praisindo/entity"
	"praisindo/service"
	"time"
)

// userRepository adalah implementasi dari IUserRepository yang menggunakan slice untuk menyimpan data pengguna
type userRepository struct {
	db     []entity.User // slice untuk menyimpan data pengguna
	nextID int           // ID berikutnya yang akan digunakan untuk pengguna baru
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db []entity.User) service.IUserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateUser menambahkan pengguna baru ke repository
func (r *userRepository) CreateUser(user *entity.User) entity.User {
	user.ID = r.nextID          // Set ID pengguna baru
	r.nextID++                  // Increment ID berikutnya
	user.CreatedAt = time.Now() // Set waktu pembuatan
	user.UpdatedAt = time.Now() // Set waktu pembaruan
	r.db = append(r.db, *user)  // Tambahkan pengguna ke slice
	return *user                // Kembalikan pengguna yang dibuat
}

// GetUserByID mencari pengguna berdasarkan ID
func (r *userRepository) GetUserByID(id int) (entity.User, bool) {
	for _, user := range r.db {
		if user.ID == id {
			return user, true // Kembalikan pengguna jika ditemukan
		}
	}
	return entity.User{}, false // Kembalikan false jika tidak ditemukan
}

// UpdateUser memperbarui pengguna berdasarkan ID
func (r *userRepository) UpdateUser(id int, user entity.User) (entity.User, bool) {
	for i, u := range r.db {
		if u.ID == id {
			user.ID = id                 // Pertahankan ID yang sama
			user.CreatedAt = u.CreatedAt // Pertahankan waktu pembuatan asli
			user.UpdatedAt = time.Now()  // Set waktu pembaruan
			r.db[i] = user               // Perbarui pengguna di slice
			return user, true            // Kembalikan pengguna yang diperbarui
		}
	}
	return entity.User{}, false // Kembalikan false jika pengguna tidak ditemukan
}

// DeleteUser menghapus pengguna berdasarkan ID
func (r *userRepository) DeleteUser(id int) bool {
	for i, user := range r.db {
		if user.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...) // Hapus pengguna dari slice
			return true                            // Kembalikan true jika berhasil
		}
	}
	return false // Kembalikan false jika pengguna tidak ditemukan
}

// GetAllUsers mengembalikan semua pengguna
func (r *userRepository) GetAllUsers() []entity.User {
	return r.db // Kembalikan slice semua pengguna
}
