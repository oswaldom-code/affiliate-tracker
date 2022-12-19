package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/affiliate-tracker/src/adapters/http/rest/dto"
	AffiliateTrackerServices "github.com/oswaldom-code/affiliate-tracker/src/services/affiliate_tracker_services"
	AffiliateTrackerTypes "github.com/oswaldom-code/affiliate-tracker/src/services/affiliate_tracker_services/types"
)

func (h *Handler) GetReferralLink(c *gin.Context) {
	// dependency injection AffiliateTrackerServices
	affiliateTrackerServices := AffiliateTrackerServices.NewTrackerService()
	referralLink := AffiliateTrackerTypes.ReferralLink{}
	err := c.BindJSON(&referralLink) // get payload
	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.Response{
				Success: false,
				Message: "Your request could not be processed successfully",
				Error:   err.Error(),
			})
		return
	}
	identifierCreated, err := affiliateTrackerServices.GenerateIdentifier(referralLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			dto.Response{
				Success: false,
				Message: "Oops at this time we can not process requests, please try again later",
				Error:   err.Error(),
			})
		return
	}
	// response
	c.JSON(http.StatusOK, &dto.Response{
		Success: true,
		Message: "ID generated successfully",
		Data:    identifierCreated,
	})
}

func (h *Handler) ProcessRequest(c *gin.Context) {
	// dependency injection AffiliateTrackerServices
	affiliateTrackerServices := AffiliateTrackerServices.NewTrackerService()
	ref, found := c.GetQuery("ref")
	if !found {
		c.JSON(http.StatusInternalServerError,
			dto.Response{
				Success: false,
				Message: "Bad request",
				Error:   "ref parameter not found in request",
			})
		return
	}
	// get header
	httpHeaders := dto.HTTPHeaders{}
	httpHeaders.GetHeader(c)
	httpHeaders.HTTPHeadersToReferred()
	url, err := affiliateTrackerServices.IdentifierDecoding(ref, httpHeaders.HTTPHeadersToReferred())
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			dto.Response{
				Success: false,
				Message: "Oops at this time we can not process requests, please try again later",
				Error:   err.Error(),
			})
		return
	}
	// redirect to job
	c.Redirect(http.StatusMovedPermanently, url)
}

func (h *Handler) ReferredGetAll(c *gin.Context) {
	// dependency injection AffiliateTrackerServices
	affiliateTrackerServices := AffiliateTrackerServices.NewTrackerService()

	referredList, err := affiliateTrackerServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			dto.Response{
				Success: false,
				Message: "Oops at this time we can not process requests, please try again later",
				Error:   err.Error(),
			})
		return
	}

	var message string
	if len(referredList) == 0 {
		message = "There are no records to display"
	} else {
		message = fmt.Sprintf("There are %v records to display", len(referredList))
	}
	// response format
	c.JSON(http.StatusOK, &dto.Response{
		Success: true,
		Message: message,
		Data:    referredList,
	})
}

func (h *Handler) ReferredGetById(c *gin.Context) {
	// dependency injection AffiliateTrackerServices
	affiliateTrackerServices := AffiliateTrackerServices.NewTrackerService()
	// get parameters from request
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.Response{
				Success: false,
				Message: "Something has gone wrong, please check the request parameters",
				Error:   err.Error(),
			})
		return
	}
	referred, err := affiliateTrackerServices.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound,
			dto.Response{
				Success: false,
				Message: "Oops at this time we can not process requests, please try again later",
				Error:   err.Error(),
			})
		return
	}

	// response format
	c.JSON(http.StatusFound, &dto.Response{
		Success: true,
		Message: "Query completed successfully",
		Data:    referred,
	})
}

func (h *Handler) ReferredGetByAgentId(c *gin.Context) {
	// dependency injection AffiliateTrackerServices
	affiliateTrackerServices := AffiliateTrackerServices.NewTrackerService()
	// get parameters from request
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest,
			dto.Response{
				Success: false,
				Message: "Something has gone wrong, please check the request parameters",
				Error:   "The id parameter is required",
			})
		return
	}
	referred, err := affiliateTrackerServices.GetByAgentId(id)
	if err != nil {
		c.JSON(http.StatusNotFound,
			dto.Response{
				Success: false,
				Message: "Oops at this time we can not process requests, please try again later",
				Error:   err.Error(),
			})
		return
	}

	// response format
	c.JSON(http.StatusFound, &dto.Response{
		Success: true,
		Message: "Query completed successfully",
		Data:    referred,
	})
}
