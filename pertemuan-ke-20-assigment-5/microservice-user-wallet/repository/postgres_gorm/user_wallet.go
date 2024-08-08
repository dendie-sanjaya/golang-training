package postgres_gorm

import (
	"context"
	"fmt"
	"log"
	"praisindo/entity"
	pb "praisindo/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
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
	walletId := int32(req.WalletId)

	history := entity.UserSaldoHistory{
		UserIdFrom:      int(req.Id),
		UserIdTo:        int(req.Id),
		WalletId:        int(walletId),
		TypeTransaction: "credit",
		TypeCredit:      "topup",
		Total:           float32(req.Amount),
	}

	_, err := history, handler.db.Table("user_saldo_histories").Create(&history).Error
	if err != nil {
		log.Fatal("Failed to top-up:", err)
		return nil, nil
	}

	//log.Println(history.WalletId)
	// Ambil saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ? and wallet_id = ? ", req.Id, walletId).First(&saldo).Error; err != nil {
		// Jika tidak ditemukan, buat entri baru dengan saldo awal
		if err == gorm.ErrRecordNotFound {
			saldo = entity.UserSaldo{
				UserId:   int(req.Id),
				WalletId: int(walletId),
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
		WalletId:        int32(history.WalletId),
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

func (handler *UserWalletHandler) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletResponse, error) {
	log.Println("Create Wallet")
	userId := int32(req.UserId)
	typeWallet := req.Type
	name := req.Name

	// Input Credit
	dataWallet := entity.UserWalet{
		UserId: int(userId),
		Type:   typeWallet,
		Name:   name,
	}

	// Create the wallet entry
	err := handler.db.Table("user_walets").Create(&dataWallet).Error
	if err != nil {
		log.Fatal("Failed to Create wallet:", err)
		return nil, fmt.Errorf("failed to create wallet: %v", err)
	}

	// Retrieve the last inserted wallet entry
	var createdWallet entity.UserWalet
	err = handler.db.Table("user_walets").Where("id = ?", dataWallet.Id).First(&createdWallet).Error
	if err != nil {
		log.Fatal("Failed to retrieve created wallet:", err)
		return nil, fmt.Errorf("failed to retrieve created wallet: %v", err)
	}

	log.Println("Wallet Created successfully")

	return &pb.CreateWalletResponse{
		UserId:    int32(createdWallet.UserId),
		WalletId:  int32(createdWallet.Id),
		Type:      createdWallet.Type,
		Name:      createdWallet.Name,
		CreatedAt: timestamppb.New(createdWallet.CreatedAt),
	}, nil
}

func (handler *UserWalletHandler) UpdateWallet(ctx context.Context, req *pb.UpdateWalletRequest) (*pb.UpdateWalletResponse, error) {
	walletId := int32(req.WalletId)
	typeWallet := req.Type
	name := req.Name

	log.Println("Update Wallet: ", int32(req.WalletId))

	// Update Credit
	dataWallet := entity.UserWalet{
		Type: typeWallet,
		Name: name,
	}

	// Define a struct to hold the updated data
	var updatedWallet entity.UserWalet

	err := handler.db.Table("user_walets").Where("id = ?", walletId).Updates(&dataWallet).First(&updatedWallet).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Failed to update wallet: record not found")
			return nil, fmt.Errorf("wallet with ID %d not found", walletId)
		}
		log.Println("Failed to update wallet:", err)
		return nil, fmt.Errorf("failed to update wallet: %v", err)
	}

	log.Println("Wallet updated successfully :", walletId)

	return &pb.UpdateWalletResponse{
		UserId:    int32(updatedWallet.UserId),
		WalletId:  int32(updatedWallet.Id),
		Type:      updatedWallet.Type,
		Name:      updatedWallet.Name,
		CreatedAt: timestamppb.New(updatedWallet.CreatedAt),
	}, nil
}

func (handler *UserWalletHandler) DeleteWallet(ctx context.Context, req *pb.DeleteWalletRequest) (*pb.DeleteWalletResponse, error) {
	walletId := int32(req.WalletId)

	err := handler.db.Table("user_walets").Where("id = ?", walletId).Delete(nil).Error
	if err != nil {
		log.Fatal("Failed to delete wallet:", err)
		return nil, fmt.Errorf("wallet with ID %d not found", walletId)
	} else {
		log.Println("Wallet deleted successfully", req.WalletId)
	}

	return &pb.DeleteWalletResponse{
		WalletId: walletId,
	}, nil
}

func (handler *UserWalletHandler) GetUserBalanceByWallet(ctx context.Context, req *pb.GetUserBalanceByWalletRequest) (*pb.GetUserBalanceByWalletResponse, error) {
	walletId := req.WalletId
	userId := req.UserId
	// 	// 	// Ambil saldo saat ini dari tabel user_saldo
	saldo := entity.UserSaldo{}
	if err := handler.db.Table("user_saldos").Where("user_id = ? and  wallet_id = ?", userId, walletId).First(&saldo).Error; err != nil {
		return nil, fmt.Errorf("Failed to retrieve user saldo:", err)
	}

	return &pb.GetUserBalanceByWalletResponse{
		UserId:   int32(saldo.UserId),
		WalletId: int32(saldo.WalletId),
		Saldo:    float32(saldo.Saldo),
	}, nil
}

func (handler *UserWalletHandler) GetTransactionHistoryByWallet(ctx context.Context, req *pb.GetTransactionHistoryByWalletRequest) (*pb.GetTransactionHistoryByWalletResponse, error) {
	walletId := req.WalletId
	userId := req.UserId

	var transactions []entity.UserSaldoHistory

	result := handler.db.Table("user_saldo_histories").Where("(user_id_from = ? OR user_id_to = ?) AND wallet_id = ? ", userId, userId, walletId).Order("created_at DESC").Find(&transactions)
	if result.Error != nil {
		log.Fatal("Failed to retrieve user wallet saldo:", result.Error)
		return nil, fmt.Errorf("Failed to retrieve user wallet saldo:", result.Error)
	}

	// Konversi transactions ke format yang diinginkan oleh gRPC response
	var transactionResponses []*pb.HistoryTransaction
	for _, transaction := range transactions {
		transactionResponse := &pb.HistoryTransaction{
			Id:              int32(transaction.Id),
			WalletId:        int32(transaction.WalletId),
			UserIdFrom:      int32(transaction.UserIdFrom),
			UserIdTo:        int32(transaction.UserIdTo),
			TypeTransaction: transaction.TypeTransaction,
			TypeCredit:      transaction.TypeCredit,
			Total:           float32(transaction.Total),
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return &pb.GetTransactionHistoryByWalletResponse{
		History: transactionResponses,
	}, nil
}

func (handler *UserWalletHandler) GetSpend(ctx context.Context, req *pb.GetSpendRequest) (*pb.GetSpendResponse, error) {

	userIDFrom := int(req.UserIdFrom)
	userIDTo := int(req.UserIdTo)
	amount := float32(req.Amount)
	walletIdFrom := int(req.WalletIdFrom)
	walletIdTo := int(req.WalletIdTo)
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

	historyCreditResponse := &pb.HistoryTransaction{
		Id:              int32(historyCredit.Id),
		UserIdFrom:      int32(historyCredit.UserIdFrom),
		UserIdTo:        int32(historyCredit.UserIdTo),
		WalletId:        int32(historyCredit.WalletId),
		TypeTransaction: historyCredit.TypeTransaction,
		TypeCredit:      historyCredit.TypeCredit,
		Total:           float32(historyCredit.Total),
	}

	historyDebitResponse := &pb.HistoryTransaction{
		Id:              int32(historyCredit.Id),
		UserIdFrom:      int32(historyCredit.UserIdFrom),
		UserIdTo:        int32(historyCredit.UserIdTo),
		WalletId:        int32(historyDebit.WalletId),
		TypeTransaction: historyDebit.TypeTransaction,
		TypeCredit:      historyDebit.TypeCredit,
		Total:           float32(historyCredit.Total),
	}

	return &pb.GetSpendResponse{
			History1: historyCreditResponse,
			History2: historyDebitResponse,
		},
		nil
}

func NewUserWalletHandler(db *gorm.DB) *UserWalletHandler {
	return &UserWalletHandler{db: db}
}
