package postgres_gorm

import (
	"log"
	"praisindo/entity"
	pb "praisindo/proto"

	"gorm.io/gorm"
)

// IUserRepository mendefinisikan interface untuk repository pengguna
// type IUserWalletRepository interface {
// 	TopUp(db *gorm.DB, userID int, amount float64) (entity.UserSaldoHistory, error)
// 	Transfer(db *gorm.DB, userIDFrom int, userIDTo int, amount float64) (entity.UserSaldoHistory, entity.UserSaldoHistory, error)
// 	GetUserBalance(db *gorm.DB, userID int) (float64, error)
// 	GetTransactionHistory(db *gorm.DB, userID int) ([]entity.UserSaldoHistory, error)
// }

// func NewUserWalletHandler(db *gorm.DB) IUserWalletRepository {
// 	return nil
// }

type UserWalletHandler struct {
	pb.UnimplementedUserWalletServiceServer
	db *gorm.DB
}

func NewUserWalletHandler(db *gorm.DB) *UserWalletHandler {
	return &UserWalletHandler{db: db}
}

// Top-up function
func TopUp(db *gorm.DB, userID int, amount float64) (entity.UserSaldoHistory, error) {

	// Input history top-up
	history := entity.UserSaldoHistory{
		UserIdFrom:      userID,
		UserIdTo:        userID,
		TypeTransaction: "credit",
		TypeCredit:      "topup",
		Total:           amount,
	}

	_, err := history, db.Table("user_saldo_history").Create(&history).Error
	if err != nil {
		log.Fatal("Failed to top-up:", err)
		return history, err

	}

	// Ambil saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := db.Table("user_saldo").Where("user_id = ?", userID).First(&saldo).Error; err != nil {
		// Jika tidak ditemukan, buat entri baru dengan saldo awal
		if err == gorm.ErrRecordNotFound {
			saldo = entity.UserSaldo{
				UserId: userID,
				Saldo:  0.0,
			}
			if err := db.Table("user_saldo").Create(&saldo).Error; err != nil {
				log.Fatal("Failed to create user saldo:", err)
				return history, err
			}
		} else {
			log.Fatal("Failed to retrieve user saldo:", err)
			return history, err
		}
	}

	// Perbarui saldo dengan menambahkan jumlah top-up
	saldo.Saldo += amount
	if err := db.Table("user_saldo").Where("user_id = ?", userID).Update("saldo", saldo.Saldo).Error; err != nil {
		log.Fatal("Failed to update user saldo:", err)
		return history, nil
	}
	return history, nil
}

// Transfer function
func Transfer(db *gorm.DB, userIDFrom int, userIDTo int, amount float64) (entity.UserSaldoHistory, entity.UserSaldoHistory, error) {

	// Input Credit
	historyCredit := entity.UserSaldoHistory{
		UserIdFrom:      userIDFrom,
		UserIdTo:        userIDTo,
		TypeTransaction: "credit",
		TypeCredit:      "transfer",
		Total:           amount,
	}

	err := db.Table("user_saldo_history").Create(&historyCredit).Error
	if err != nil {
		log.Fatal("Failed to history credit:", err)
		return historyCredit, historyCredit, err

	}

	// update saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := db.Table("user_saldo").Where("user_id = ?", userIDTo).First(&saldo).Error; err != nil {
		// Jika tidak ditemukan, buat entri baru dengan saldo awal
		if err == gorm.ErrRecordNotFound {
			saldo = entity.UserSaldo{
				UserId: userIDTo,
				Saldo:  0.0,
			}
			if err := db.Table("user_saldo").Create(&saldo).Error; err != nil {
				log.Fatal("Failed to create user saldo credit:", err)
				return historyCredit, historyCredit, err
			}
		} else {
			log.Fatal("Failed to retrieve user saldo credit:", err)
			return historyCredit, historyCredit, err
		}
	}

	// Perbarui saldo dengan menambahkan jumlah credit
	saldo.Saldo += amount
	if err := db.Table("user_saldo").Where("user_id = ?", userIDTo).Update("saldo", saldo.Saldo).Error; err != nil {
		log.Fatal("Failed to update user saldo credit:", err)
		return historyCredit, historyCredit, nil
	}

	// Input Debit
	historyDebit := entity.UserSaldoHistory{
		UserIdFrom:      userIDTo,
		UserIdTo:        userIDFrom,
		TypeTransaction: "debit",
		TypeCredit:      "transfer",
		Total:           amount,
	}

	err = db.Table("user_saldo_history").Create(&historyDebit).Error
	if err != nil {
		log.Fatal("Failed to top-up debit:", err)
		return historyDebit, historyDebit, err

	}

	// Ambil saldo saat ini dari tabel user_saldo
	saldo = entity.UserSaldo{}
	if err := db.Table("user_saldo").Where("user_id = ?", userIDFrom).First(&saldo).Error; err != nil {
		log.Fatal("Failed to retrieve user saldo debit:", err)
		return historyDebit, historyDebit, err
	}

	// Perbarui saldo dengan mengurangi jumlah debit
	saldo.Saldo -= amount
	if err := db.Table("user_saldo").Where("user_id = ?", userIDFrom).Update("saldo", saldo.Saldo).Error; err != nil {
		log.Fatal("Failed to update user saldo debit:", err)
		return historyDebit, historyDebit, nil
	}

	return historyCredit, historyDebit, nil
}

// Get user balance function
func GetUserBalance(db *gorm.DB, userID int) (float64, error) {
	var balance float64
	balance = 0.0

	// Ambil saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := db.Table("user_saldo").Where("user_id = ?", userID).First(&saldo).Error; err != nil {
		log.Fatal("Failed to retrieve user saldo:", err)
		return balance, err
	}

	return saldo.Saldo, nil
}

// Get transaction history function
func GetTransactionHistory(db *gorm.DB, userID int) ([]entity.UserSaldoHistory, error) {
	var history []entity.UserSaldoHistory
	err := db.Table("user_saldo_history").Where("user_id_from = ? OR user_id_to = ?", userID, userID).Order("created_at DESC").Find(&history).Error

	if err != nil {
		log.Fatal("Failed to retrieve history:", err)
		return history, err
	}
	return history, err
}
