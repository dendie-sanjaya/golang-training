package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"praisindo/entity"
	"praisindo/service"

	"github.com/gin-gonic/gin"
)

// ISubmissionHandler mendefinisikan interface untuk handler Submission
type ISubmissionHandler interface {
	CreateSubmission(c *gin.Context)
	GetSubmission(c *gin.Context)
	DeleteSubmission(c *gin.Context)
	GetAllSubmission(c *gin.Context)
}

type SubmissionHandler struct {
	submissionService service.ISubmissionService
}

// NewSubmissionHandler membuat instance baru dari SubmissionHandler
func NewSubmissionHandler(submissionService service.ISubmissionService) ISubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
	}
}

// CreateSubmission menghandle permintaan untuk membuat Submission baru
func (h *SubmissionHandler) CreateSubmission(c *gin.Context) {
	var Submission entity.Submission
	if err := c.ShouldBindJSON(&Submission); err != nil {
		errMsg := err.Error()
		errMsg = convertSubmissionMandatoryFieldErrorString(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	createdSubmission, err := h.submissionService.CreateSubmissions(c.Request.Context(), &Submission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdSubmission)
}

// GetSubmission menghandle permintaan untuk mendapatkan Submission berdasarkan ID
func (h *SubmissionHandler) GetSubmission(c *gin.Context) {

	if c.Query("user_id") != "" {
		user_id, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID xxx"})
			return
		}

		Submission, err := h.submissionService.GetSubmissionsByID(c.Request.Context(), user_id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, Submission)
	} else {
		Submission, err := h.submissionService.GetAllSubmissions(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, Submission)
	}

}

// DeleteSubmission menghandle permintaan untuk menghapus Submission
func (h *SubmissionHandler) DeleteSubmission(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.submissionService.DeleteSubmissions(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission deleted"})
}

type getAllSubmissionData struct {
	ID           string          `json:"ID"`            // ID pengguna sebagai primary key
	UserId       int             `json:"user_id"`       // Kata sandi pengguna (wajib diisi)
	Answer       json.RawMessage `json:"answer"`        // Kata sandi pengguna (wajib diisi)
	RiskScore    int             `json:"risk_scrore"`   // Kata sandi pengguna (wajib diisi)
	RiskCategory string          `json:"risk_category"` // Kata sandi pengguna (wajib diisi)
	CreatedAt    time.Time       `json:"created_at"`    // Waktu pembuatan pengguna
	UpdatedAt    time.Time       `json:"updated_at"`    // Waktu pembaruan terakhir pengguna
}

type getAllSubmissionResponse struct {
	Data    []getAllSubmissionData `json:"Submissions`
	Message string
	Code    int
}

// GetAllSubmissions menghandle permintaan untuk mendapatkan semua Submission
func (h *SubmissionHandler) GetAllSubmission(c *gin.Context) {
	fmt.Print("Masuk ke GetAllSubmissions")
	Submissions, err := h.submissionService.GetAllSubmissions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res getAllSubmissionResponse
	for _, u := range Submissions {
		res.Data = append(res.Data, getAllSubmissionData{
			ID:           strconv.Itoa(u.ID),
			UserId:       u.UserId,
			Answer:       u.Answers,
			RiskScore:    u.RiskScore,
			RiskCategory: u.RiskCategory,
		})
	}
	res.Code = 200
	res.Message = "Sukses mendapatkan semua pengguna"
	c.JSON(http.StatusOK, res)
}

func convertSubmissionMandatoryFieldErrorString(oldErrorMsg string) string {
	switch {
	case strings.Contains(oldErrorMsg, "'Name' failed on the 'required' tag"):
		return "name is mandatory"
	case strings.Contains(oldErrorMsg, "'Email' failed on the 'required' tag"):
		return "email is mandatory"
	}
	return oldErrorMsg
}
