package dto

import (
	"github.com/gin-gonic/gin"
	AffiliateTrackerTypes "github.com/oswaldom-code/affiliate-tracker/src/services/affiliate_tracker_services/types"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

type HTTPHeaders struct {
	Referer       string `json:"referer"`
	UserAgent     string `json:"user_agent"`
	XForwardedFor string
}

func (r *HTTPHeaders) GetHeader(c *gin.Context) {
	r.Referer = c.GetHeader("Referer")
	r.UserAgent = c.GetHeader("User-Agent")
	r.XForwardedFor = c.GetHeader("X-Forwarded-For")
}

func (r *HTTPHeaders) HTTPHeadersToReferred() AffiliateTrackerTypes.ReferredDTO {
	return AffiliateTrackerTypes.ReferredDTO{
		RequestReferer:   r.Referer,
		RequestIp:        r.XForwardedFor,
		RequestUserAgent: r.UserAgent,
	}
}
