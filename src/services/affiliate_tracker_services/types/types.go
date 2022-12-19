package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/oswaldom-code/affiliate-tracker/src/domain/models"
)

type ReferralLink struct {
	Agent string `json:"agent"`
	Url   string `json:"url"`
}

func (r *ReferralLink) ToString() string {
	return fmt.Sprintf("%s:::%s", r.Agent, r.Url)
}

type ReferredDTO struct {
	Id               int64     `json:"id"`
	AgentId          string    `json:"agent"`
	JobUrl           string    `json:"url"`
	RequestReferer   string    `json:"referer"`
	RequestIp        string    `json:"ip"`
	RequestUserAgent string    `json:"user_agent"`
	CreatedAt        time.Time `json:"cread_at"`
}

func (r *ReferredDTO) ToModel() models.Referred {
	return models.Referred{
		AgentId:          strings.ToUpper(r.AgentId),
		JobUrl:           r.JobUrl,
		RequestReferer:   r.RequestReferer,
		RequestIp:        r.RequestIp,
		RequestUserAgent: r.RequestUserAgent,
		CreatedAt:        time.Now(),
	}
}

func (r *ReferredDTO) ReferredModelToReferredDTO(referred models.Referred) {
	r.Id = referred.ID
	r.AgentId = referred.AgentId
	r.JobUrl = referred.JobUrl
	r.RequestReferer = referred.RequestReferer
	r.RequestIp = referred.RequestIp
	r.RequestUserAgent = referred.RequestUserAgent
	r.CreatedAt = referred.CreatedAt
}

func (r *ReferredDTO) ReferredModelsToReferredDTOs(referredList []models.Referred) []ReferredDTO {
	var referredDTOList []ReferredDTO
	for _, referred := range referredList {
		referredDTO := ReferredDTO{}
		referredDTO.ReferredModelToReferredDTO(referred)
		referredDTOList = append(referredDTOList, referredDTO)
	}
	return referredDTOList
}
