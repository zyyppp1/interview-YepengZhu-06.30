package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zyyppp1/interview-YepengZhu-06.30/models"
	"github.com/zyyppp1/interview-YepengZhu-06.30/services"
)

// ProcessPayment 处理支付
func ProcessPayment(c *gin.Context) {
	var req struct {
		PlayerID       uuid.UUID              `json:"player_id" binding:"required"`
		PaymentMethod  string                 `json:"payment_method" binding:"required"`
		Amount         float64                `json:"amount" binding:"required,gt=0"`
		PaymentDetails map[string]interface{} `json:"payment_details,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 创建支付记录
	payment := &models.Payment{
		PlayerID:       req.PlayerID,
		PaymentMethod:  req.PaymentMethod,
		Amount:         req.Amount,
		Currency:       "CNY",
		PaymentDetails: req.PaymentDetails,
	}

	// 处理支付
	if err := services.Payment.ProcessPayment(payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回响应
	c.JSON(http.StatusCreated, gin.H{
		"payment_id":     payment.ID.String(),
		"transaction_id": payment.TransactionID,
		"status":         payment.Status,
		"message":        getPaymentMessage(payment.Status),
		"processed_at":   payment.CreatedAt,
	})
}

// GetPayment 获取支付详情
func GetPayment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payment ID",
		})
		return
	}

	payment, err := services.Payment.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    payment,
	})
}

// getPaymentMessage 根据支付状态返回消息
func getPaymentMessage(status string) string {
	messages := map[string]string{
		"pending":    "Payment is being processed",
		"processing": "Payment is being processed",
		"success":    "Payment completed successfully",
		"failed":     "Payment failed",
		"refunded":   "Payment has been refunded",
	}
	
	if msg, ok := messages[status]; ok {
		return msg
	}
	return "Unknown payment status"
}