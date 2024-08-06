package postgres_gorm

import (
	"context"
	"log"
	"praisindo/entity"
	pb "praisindo/proto"

	"gorm.io/gorm"
)

type UserWalletHandler struct {
	pb.UnimplementedUserWalletServiceServer
	db *gorm.DB
}

func (handler *UserWalletHandler) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {

	userIDFrom := int(req.From)
	userIDTo := int(req.To)
	amount := float32(req.Amount)

	// Input Credit
	historyCredit := entity.UserSaldoHistory{
		UserIdFrom:      userIDFrom,
		UserIdTo:        userIDTo,
		TypeTransaction: "credit",
		TypeCredit:      "transfer",
		Total:           amount,
	}

	err := handler.db.Table("user_saldo_histories").Create(&historyCredit).Error
	if err != nil {
		log.Fatal("Failed to history credit:", err)
		//return historyCredit, historyCredit, err

	}

	// update saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ?", userIDTo).First(&saldo).Error; err != nil {
		// Jika tidak ditemukan, buat entri baru dengan saldo awal
		if err == gorm.ErrRecordNotFound {
			saldo = entity.UserSaldo{
				UserId: userIDTo,
				Saldo:  0.0,
			}
			if err := handler.db.Table("user_saldos").Create(&saldo).Error; err != nil {
				log.Fatal("Failed to create user saldo credit:", err)
				//return historyCredit, historyCredit, err
			}
		} else {
			log.Fatal("Failed to retrieve user saldo credit:", err)
			//return historyCredit, historyCredit, err
		}
	}

	// Perbarui saldo dengan menambahkan jumlah credit
	saldo.Saldo += amount
	if err := handler.db.Table("user_saldos").Where("user_id = ?", userIDTo).Update("saldo", saldo.Saldo).Error; err != nil {
		log.Fatal("Failed to update user saldo credit:", err)
		//return historyCredit, historyCredit, nil
	}

	// Input Debit
	historyDebit := entity.UserSaldoHistory{
		UserIdFrom:      userIDTo,
		UserIdTo:        userIDFrom,
		TypeTransaction: "debit",
		TypeCredit:      "transfer",
		Total:           amount,
	}

	err = handler.db.Table("user_saldo_histories").Create(&historyDebit).Error
	if err != nil {
		log.Fatal("Failed to top-up debit:", err)
		//return historyDebit, historyDebit, err

	}

	// Ambil saldo saat ini dari tabel user_saldo
	saldo = entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ?", userIDFrom).First(&saldo).Error; err != nil {
		log.Fatal("Failed to retrieve user saldo debit:", err)
		//return historyDebit, historyDebit, err
	}

	// Perbarui saldo dengan mengurangi jumlah debit
	saldo.Saldo -= amount
	if err := handler.db.Table("user_saldos").Where("user_id = ?", userIDFrom).Update("saldo", saldo.Saldo).Error; err != nil {
		log.Fatal("Failed to update user saldo debit:", err)
		//return historyDebit, historyDebit, nil
	}

	historyCreditResponse := &pb.HistoryTransaction{
		Id:              int32(historyCredit.Id),
		UserIdFrom:      int32(historyCredit.UserIdFrom),
		UserIdTo:        int32(historyCredit.UserIdTo),
		TypeTransaction: historyCredit.TypeTransaction,
		TypeCredit:      historyCredit.TypeCredit,
		Total:           float32(historyCredit.Total),
	}

	historyDebitResponse := &pb.HistoryTransaction{
		Id:              int32(historyCredit.Id),
		UserIdFrom:      int32(historyCredit.UserIdFrom),
		UserIdTo:        int32(historyCredit.UserIdTo),
		TypeTransaction: historyDebit.TypeTransaction,
		TypeCredit:      historyDebit.TypeCredit,
		Total:           float32(historyCredit.Total),
	}

	return &pb.TransferResponse{
			History1: historyCreditResponse,
			History2: historyDebitResponse,
		},
		nil
}

func (handler *UserWalletHandler) Topup(ctx context.Context, req *pb.TopupRequest) (*pb.TopupResponse, error) {
	// Input history top-up
	walletId := 1

	history := entity.UserSaldoHistory{
		UserIdFrom:      int(req.Id),
		UserIdTo:        int(req.Id),
		WalletId:        walletId,
		TypeTransaction: "credit",
		TypeCredit:      "topup",
		Total:           float32(req.Amount),
	}

	_, err := history, handler.db.Table("user_saldo_histories").Create(&history).Error
	if err != nil {
		log.Fatal("Failed to top-up:", err)
		//return history, err

	}

	// Ambil saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ?", req.Id).First(&saldo).Error; err != nil {
		// Jika tidak ditemukan, buat entri baru dengan saldo awal
		if err == gorm.ErrRecordNotFound {
			saldo = entity.UserSaldo{
				UserId:   int(req.Id),
				WalletId: walletId,
				Saldo:    0.0,
			}
			if err := handler.db.Table("user_saldos").Create(&saldo).Error; err != nil {
				log.Fatal("Failed to create user saldo:", err)
				//return history, err
			}
		} else {
			log.Fatal("Failed to retrieve user saldo:", err)
			//return history, err
		}
	}

	// Perbarui saldo dengan menambahkan jumlah top-up
	saldo.Saldo += float32(req.Amount)
	if err := handler.db.Table("user_saldos").Where("user_id = ? AND wallet_id = ?", req.Id, walletId).Update("saldo", saldo.Saldo).Error; err != nil {
		log.Fatal("Failed to update user saldo:", err)
		//return history, nil
	}

	// Convert entity.UserSaldoHistory to *user_wallet.HistoryTransaction
	historyTransaction := &pb.HistoryTransaction{
		Id:              int32(history.Id),
		UserIdFrom:      int32(history.UserIdFrom),
		UserIdTo:        int32(history.UserIdTo),
		TypeTransaction: history.TypeTransaction,
		TypeCredit:      history.TypeCredit,
		Total:           float32(history.Total),
	}

	return &pb.TopupResponse{
		History: historyTransaction,
	}, nil
}

func (handler *UserWalletHandler) GetUserBalance(ctx context.Context, req *pb.GetUserBalanceRequest) (*pb.GetUserBalanceResponse, error) {

	// 	// 	// Ambil saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ?", req.Id).First(&saldo).Error; err != nil {
		log.Fatal("Failed to retrieve user saldo:", err)
	}

	return &pb.GetUserBalanceResponse{
		UserId: req.Id,
		Saldo:  float32(saldo.Saldo),
	}, nil
}

func (handler *UserWalletHandler) GetTransactionHistory(ctx context.Context, req *pb.GetTransactionHistoryRequest) (*pb.GetTransactionHistoryResponse, error) {
	var transactions []entity.UserSaldoHistory

	result := handler.db.Table("user_saldo_histories").Where("user_id_from = ? OR user_id_to = ?", req.UserId, req.UserId).Order("created_at DESC").Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	// Konversi transactions ke format yang diinginkan oleh gRPC response
	var transactionResponses []*pb.HistoryTransaction
	for _, transaction := range transactions {
		transactionResponse := &pb.HistoryTransaction{
			Id:              int32(transaction.Id),
			UserIdFrom:      int32(transaction.UserIdFrom),
			UserIdTo:        int32(transaction.UserIdTo),
			TypeTransaction: transaction.TypeTransaction,
			TypeCredit:      transaction.TypeCredit,
			Total:           float32(transaction.Total),
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return &pb.GetTransactionHistoryResponse{
		History: transactionResponses,
	}, nil
}

func (handler *UserWalletHandler) CreateWallet(ctx context.Context) {
	userId := 1
	typeWallet := "in"
	name := "gaji"

	// Input Credit
	dataWallet := entity.UserWalet{
		Id:   userId,
		Type: typeWallet,
		Name: name,
	}

	err := handler.db.Table("user_walets").Create(&dataWallet).Error
	if err != nil {
		log.Fatal("Failed to create wallet :", err)

	}
}
func (handler *UserWalletHandler) UpdateWallet(ctx context.Context) {
	userId := 1
	typeWallet := "in"
	name := "gaji edit"

	// Update Credit
	dataWallet := entity.UserWalet{
		Type: typeWallet,
		Name: name,
	}

	err := handler.db.Table("user_walets").Where("id = ?", userId).Updates(&dataWallet).Error
	if err != nil {
		log.Fatal("Failed to update wallet:", err)
	} else {
		log.Println("Wallet updated successfully")
	}
}

func (handler *UserWalletHandler) DeleteWallet(ctx context.Context) {
	userId := 1

	err := handler.db.Table("user_walets").Where("id = ?", userId).Delete(nil).Error
	if err != nil {
		log.Fatal("Failed to delete wallet:", err)
	} else {
		log.Println("Wallet deleted successfully")
	}
}

func (handler *UserWalletHandler) GetUserBalanceByWallet(ctx context.Context) {
	walletId := 1
	userId := 1
	// 	// 	// Ambil saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ? and  wallet_id = ?", userId, walletId).First(&saldo).Error; err != nil {
		log.Fatal("Failed to retrieve user saldo:", err)
	}

	// return &pb.GetUserBalanceResponse{
	// 	UserId: req.Id,
	// 	Saldo:  float32(saldo.Saldo),
	// }, nil
}

func (handler *UserWalletHandler) GetTransactionHistoryByWallet(ctx context.Context) {
	walletId := 1
	userId := 1

	var transactions []entity.UserSaldoHistory

	result := handler.db.Table("user_saldo_histories").Where("user_id_from = ? OR user_id_to = ? AND wallet_id = ? ", userId, userId, walletId).Order("created_at DESC").Find(&transactions)
	if result.Error != nil {
		log.Fatal("Failed to retrieve user saldo:", result.Error)
		//return nil, result.Error
	}

	// Konversi transactions ke format yang diinginkan oleh gRPC response
	var transactionResponses []*pb.HistoryTransaction
	for _, transaction := range transactions {
		transactionResponse := &pb.HistoryTransaction{
			Id:              int32(transaction.Id),
			UserIdFrom:      int32(transaction.UserIdFrom),
			UserIdTo:        int32(transaction.UserIdTo),
			TypeTransaction: transaction.TypeTransaction,
			TypeCredit:      transaction.TypeCredit,
			Total:           float32(transaction.Total),
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	/*
		return &pb.GetTransactionHistoryResponse{
			History: transactionResponses,
		}, nil
	*/
}

func (handler *UserWalletHandler) Spend(ctx context.Context) {

	// userIDFrom := int(req.From)
	// userIDTo := int(req.To)
	// amount := float32(req.Amount)

	userIDFrom := 1
	userIDTo := 1
	amount := float32(1000.00)
	walletIdFrom := 1
	walletIdTo := 2

	// Input Credit
	historyCredit := entity.UserSaldoHistory{
		UserIdFrom:      userIDFrom,
		UserIdTo:        userIDTo,
		WalletId:        walletIdTo,
		TypeTransaction: "credit",
		TypeCredit:      "",
		Total:           amount,
	}

	err := handler.db.Table("user_saldo_histories").Create(&historyCredit).Error
	if err != nil {
		log.Fatal("Failed to history credit:", err)
		//return historyCredit, historyCredit, err

	}

	// update saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ? and wallet_id = ? ", userIDTo, walletIdTo).First(&saldo).Error; err != nil {
		// Jika tidak ditemukan, buat entri baru dengan saldo awal
		if err == gorm.ErrRecordNotFound {
			saldo = entity.UserSaldo{
				UserId:   userIDTo,
				WalletId: walletIdTo,
				Saldo:    0.0,
			}
			if err := handler.db.Table("user_saldos").Create(&saldo).Error; err != nil {
				log.Fatal("Failed to create user saldo credit:", err)
				//return historyCredit, historyCredit, err
			}
		} else {
			log.Fatal("Failed to retrieve user saldo credit:", err)
			//return historyCredit, historyCredit, err
		}
	}

	// Perbarui saldo dengan menambahkan jumlah credit
	saldo.Saldo += amount
	if err := handler.db.Table("user_saldos").Where("user_id = ? and wallet_id = ? ", userIDTo, walletIdTo).Update("saldo", saldo.Saldo).Error; err != nil {
		log.Fatal("Failed to update user saldo credit:", err)
		//return historyCredit, historyCredit, nil
	}

	// Input Debit
	historyDebit := entity.UserSaldoHistory{
		UserIdFrom:      userIDTo,
		UserIdTo:        userIDFrom,
		WalletId:        walletIdFrom,
		TypeTransaction: "debit",
		TypeCredit:      "",
		Total:           amount,
	}

	err = handler.db.Table("user_saldo_histories").Create(&historyDebit).Error
	if err != nil {
		log.Fatal("Failed to top-up debit:", err)
		//return historyDebit, historyDebit, err

	}

	// Ambil saldo saat ini dari tabel user_saldo
	saldo = entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ? and wallet_id = ?", userIDFrom, walletIdFrom).First(&saldo).Error; err != nil {
		log.Fatal("Failed to retrieve user saldo debit:", err)
		//return historyDebit, historyDebit, err
	}

	// Perbarui saldo dengan mengurangi jumlah debit
	saldo.Saldo -= amount
	if err := handler.db.Table("user_saldos").Where("user_id = ? and wallet_id = ?", userIDFrom, walletIdFrom).Update("saldo", saldo.Saldo).Error; err != nil {
		log.Fatal("Failed to update user saldo debit:", err)
		//return historyDebit, historyDebit, nil
	}

	// historyCreditResponse := &pb.HistoryTransaction{
	// 	Id:              int32(historyCredit.Id),
	// 	UserIdFrom:      int32(historyCredit.UserIdFrom),
	// 	UserIdTo:        int32(historyCredit.UserIdTo),
	// 	TypeTransaction: historyCredit.TypeTransaction,
	// 	TypeCredit:      historyCredit.TypeCredit,
	// 	Total:           float32(historyCredit.Total),
	// }

	// historyDebitResponse := &pb.HistoryTransaction{
	// 	Id:              int32(historyCredit.Id),
	// 	UserIdFrom:      int32(historyCredit.UserIdFrom),
	// 	UserIdTo:        int32(historyCredit.UserIdTo),
	// 	TypeTransaction: historyDebit.TypeTransaction,
	// 	TypeCredit:      historyDebit.TypeCredit,
	// 	Total:           float32(historyCredit.Total),
	// }

	// return &pb.TransferResponse{
	// 		History1: historyCreditResponse,
	// 		History2: historyDebitResponse,
	// 	},
	// 	nil
}

func NewUserWalletHandler(db *gorm.DB) *UserWalletHandler {
	return &UserWalletHandler{db: db}
}
