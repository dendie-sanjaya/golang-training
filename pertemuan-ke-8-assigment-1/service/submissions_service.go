package service

import (
	"context"
	"fmt"
	"praisindo/entity"
)

// ISubmissionService mendefinisikan interface untuk layanan pengguna
type ISubmissionService interface {
	GetSubmissionsByID(ctx context.Context, id int, limit int, offset int) ([]entity.SubmissionData, error)
	CreateSubmissions(ctx context.Context, user *entity.Submission) (entity.Submission, error)
	DeleteSubmissions(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.SubmissionData, error)
	GetSubmissionsByIDTotal(ctx context.Context, id int) (totaSubmission int64, error error)
}

// ISubmissionRepository mendefinisikan interface untuk repository pengguna
type ISubmissionRepository interface {
	GetSubmissionsByID(ctx context.Context, id int, limit int, offset int) ([]entity.SubmissionData, error)
	CreateSubmissions(ctx context.Context, user *entity.Submission) (entity.Submission, error)
	DeleteSubmissions(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.SubmissionData, error)
	GetSubmissionsByIDTotal(ctx context.Context, id int) (totaSubmission int64, error error)
}

// submissionService adalah implementasi dari ISubmissionService yang menggunakan ISubmissionRepository
type submissionService struct {
	SubmissionRepo ISubmissionRepository
}

type Question struct {
	ID       int
	Question string
	Options  []Option
}
type Option struct {
	Answer string
	Weight int
}

type Answer struct {
	QuestionID int
	Answer     string
}

type answareJSON struct {
	user_od int
	Answers []Answer
}

type ProfileRisk struct {
	MinScore   int
	MaxScore   int
	Category   ProfileRiskCategory
	Definition string
}

type ProfileRiskCategory string

const (
	ProfileRiskCategoryConservative ProfileRiskCategory = "Conservative"
	ProfileRiskCategoryModerate     ProfileRiskCategory = "Moderate"
	ProfileRiskCategoryBalanced     ProfileRiskCategory = "Balanced"
	ProfileRiskCategoryGrowth       ProfileRiskCategory = "Growth"
	ProfileRiskCategoryAggresive    ProfileRiskCategory = "Aggresive"
)

// NewsubmissionService membuat instance baru dari submissionService
func NewSubmissionService(userRepo ISubmissionRepository) ISubmissionService {
	return &submissionService{SubmissionRepo: userRepo}
}

// GetUserByID mendapatkan pengguna berdasarkan ID
func (s *submissionService) GetSubmissionsByID(ctx context.Context, id int, limit int, offset int) ([]entity.SubmissionData, error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	fmt.Print("masuk ke GetSubmissionsByID gorm submission service \n ")
	user, err := s.SubmissionRepo.GetSubmissionsByID(ctx, id, limit, offset)
	if err != nil {
		return user, fmt.Errorf("gagal mendapatkan submission berdasarkan user id: %v", err)
	}
	return user, nil
}

func (s *submissionService) GetSubmissionsByIDTotal(ctx context.Context, id int) (TotaSubmission int64, error error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	fmt.Print("masuk ke GetSubmissionsByID gorm submission service  total \n ")
	user, err := s.SubmissionRepo.GetSubmissionsByIDTotal(ctx, id)
	if err != nil {
		return user, fmt.Errorf("gagal mendapatkan submission berdasarkan user id: %v", err)
	}
	return user, nil
}

// UpdateUser memperbarui data pengguna
func (s *submissionService) CreateSubmissions(ctx context.Context, user *entity.Submission) (entity.Submission, error) {
	// Memanggil UpdateUser dari repository untuk memperbarui data pengguna
	fmt.Print("masuk ke CreateUser gorm submission service \n ")
	var Questions = []Question{
		{
			ID:       1,
			Question: "Apakah tujuan investasi Anda?",
			Options: []Option{
				{Answer: "Pertumbuhan kekayaan untuk jangka panjang", Weight: 5},
				{Answer: "Pendapatan dan pertumbuhan dalam jangka panjang", Weight: 4},
				{Answer: "Pendapatan berkala", Weight: 3},
				{Answer: "Pendapatan dan keamanan dana investasi", Weight: 2},
				{Answer: "Keamanan dana investasi", Weight: 1},
			},
		},
		{
			ID:       2,
			Question: "Berdasarkan tujuan investasi Anda, dana Anda akan diinvestasikan untuk jangka waktu?",
			Options: []Option{
				{Answer: "≥ 10 tahun", Weight: 5},
				{Answer: "7 - 10 tahun", Weight: 4},
				{Answer: "4 - ≥ 6 tahun", Weight: 3},
				{Answer: "1 - ≥ 3 tahun", Weight: 2},
				{Answer: "< 1 tahun", Weight: 1},
			},
		},
		{
			ID:       3,
			Question: "Berapa lama pengalaman Anda berinvestasi dalam produk yang nilainya berfluktuasi?",
			Options: []Option{
				{Answer: "> 10 tahun", Weight: 5},
				{Answer: "8 - 10 tahun", Weight: 4},
				{Answer: "4 - 7 tahun", Weight: 3},
				{Answer: "< 4 tahun", Weight: 2},
				{Answer: "0 tahun (tidak memiliki pengalaman)", Weight: 1},
			},
		},
		{
			ID:       4,
			Question: "Jenis investasi apa yang pernah Anda miliki?",
			Options: []Option{
				{Answer: "Saham, Reksa Dana terbuka, equity linked structure product", Weight: 5},
				{Answer: "Mata uang asing, currency linked structured product", Weight: 4},
				{Answer: "Uang tunai, deposito, produk dengan proteksi modal", Weight: 3},
			},
		},
		{
			ID:       5,
			Question: "Berapa persen dari aset Anda yang disimpan dalam produk investasi berfluktuasi?",
			Options: []Option{
				{Answer: "> 50%", Weight: 5},
				{Answer: "> 25% - ≥ 50%", Weight: 4},
				{Answer: "> 10% - ≥ 25%", Weight: 3},
				{Answer: "> 0% - ≥ 10%", Weight: 2},
				{Answer: "0%", Weight: 1},
			},
		},
		{
			ID:       6,
			Question: "Tingkat kenaikan dan penurunan nilai investasi yang dapat Anda terima?",
			Options: []Option{
				{Answer: "< -20% - > +20%", Weight: 5},
				{Answer: "-20% - +20%", Weight: 4},
				{Answer: "-15% - +15%", Weight: 3},
				{Answer: "-10% - +10%", Weight: 2},
				{Answer: "-5% - +5%", Weight: 1},
			},
		},
		{
			ID:       7,
			Question: "Ketergantungan Anda pada hasil investasi untuk biaya hidup sehari-hari?",
			Options: []Option{
				{Answer: "Tidak bergantung pada hasil investasi", Weight: 5},
				{Answer: "Tidak bergantung pada hasil investasi, minimal 5 tahun ke depan", Weight: 4},
				{Answer: "Sedikit bergantung pada hasil investasi", Weight: 3},
				{Answer: "Bergantung pada hasil investasi", Weight: 2},
				{Answer: "Sangat bergantung pada hasil investasi", Weight: 1},
			},
		},
		{
			ID:       8,
			Question: "Persentase pendapatan bulanan yang dapat Anda sisihkan untuk investasi/tabungan?",
			Options: []Option{
				{Answer: "> 50%", Weight: 5},
				{Answer: "> 25% - 50%", Weight: 4},
				{Answer: "> 10% - 25%", Weight: 3},
				{Answer: "> 0% - 10%", Weight: 2},
				{Answer: "0%", Weight: 1},
			},
		},
	}
	var total_weight int
	for _, list_answare := range user.Answers {
		for _, list_questions := range Questions {
			if list_answare.QuestionID == list_questions.ID {
				for _, list_questions_option := range list_questions.Options {
					if list_answare.Answer == list_questions_option.Answer {
						total_weight = total_weight + list_questions_option.Weight
					}
				}
			}
		}
	}
	user.RiskScore = total_weight

	var RiskMapping = []ProfileRisk{
		{
			MinScore: 0,
			MaxScore: 11,
			Category: ProfileRiskCategoryConservative,
			Definition: "Tujuan utama Anda adalah untuk melindungi modal/dana yang ditempatkan dan Anda tidak memiliki toleransi " +
				"sama sekali terhadap perubahan harga/nilai dari dana investasinya tersebut. " +
				"Anda memiliki pengalaman yang sangat terbatas atau tidak memiliki pengalaman sama sekali mengenai produk investasi.",
		},
		{
			MinScore:   12,
			MaxScore:   19,
			Category:   ProfileRiskCategoryModerate,
			Definition: "Anda memiliki toleransi yang rendah dengan perubahan harga/nilai dari dana investasi dan risiko investasi.",
		},
		{
			MinScore: 20,
			MaxScore: 28,
			Category: ProfileRiskCategoryBalanced,
			Definition: "Anda memiliki toleransi yang cukup terhadap produk investasi dan dapat menerima perubahan yang besar dari " +
				"harga/nilai dari harga yang diinvestasikan.",
		},
		{
			MinScore: 29,
			MaxScore: 35,
			Category: ProfileRiskCategoryGrowth,
			Definition: "Anda memiliki toleransi yang cukup tinggi dan dapat menerima perubahan yang besar dari harga/nilai portfolio" +
				"pada produk investasi yang diinvestasikan." +
				"Pada umumnya Anda sudah pernah atau berpengalaman dalam berinvestasi di produk investasi.",
		},
		{
			MinScore: 36,
			MaxScore: 40,
			Category: ProfileRiskCategoryAggresive,
			Definition: "Anda sangat berpengalaman terhadap produk investasi dan memiliki toleransi yang sangat tinggi atas" +
				"produk-produk investasi. Anda bahkan dapat menerima perubahan signifikan pada modal/nilai investasi." +
				"Pada umumnya portfolio Anda sebagian besar dialokasikan pada produk investasi.",
		},
	}

	for _, list_risk_mapping := range RiskMapping {
		if total_weight >= list_risk_mapping.MinScore && total_weight <= list_risk_mapping.MaxScore {
			user.RiskCategory = string(list_risk_mapping.Category)
		}
	}

	return s.SubmissionRepo.CreateSubmissions(ctx, user)
}

// DeleteUser menghapus pengguna berdasarkan ID
func (s *submissionService) DeleteSubmissions(ctx context.Context, id int) error {
	// Memanggil DeleteUser dari repository untuk menghapus pengguna berdasarkan ID
	err := s.SubmissionRepo.DeleteSubmissions(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus pengguna: %v", err)
	}
	return nil
}

// GetAllUsers mendapatkan semua pengguna
func (s *submissionService) GetAllSubmissions(ctx context.Context) ([]entity.SubmissionData, error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua pengguna
	users, err := s.SubmissionRepo.GetAllSubmissions(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua submission: %v", err)
	}
	return users, nil
}
