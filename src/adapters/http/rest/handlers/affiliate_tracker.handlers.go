package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/affiliate-tracker/src/adapters/http/rest/dto"
	AffiliateTrackerServices "github.com/oswaldom-code/affiliate-tracker/src/services/affiliate_tracker_services"
)

func (h *Handler) GetReferralLink(c *gin.Context) {
	// dependency injection AffiliateTrackerServices
	affiliateTrackerServices := AffiliateTrackerServices.NewTrackerService()

	// get payload
	referralLink := AffiliateTrackerServices.ReferralLink{}
	err := c.BindJSON(&referralLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"success": false})
		return
	}
	stringEncrypter := affiliateTrackerServices.ProcessInputUrl(referralLink)
	// response format
	c.JSON(http.StatusOK, &dto.Response{
		Success: true,
		Message: "Success",
		Data:    stringEncrypter,
	})
}
