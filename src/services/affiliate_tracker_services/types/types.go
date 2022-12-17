package types

import "fmt"

type ReferralLink struct {
	Agent string `json:"agent"`
	Url   string `json:"url"`
}

func (r *ReferralLink) ToString() string {
	return fmt.Sprintf("%s:::%s", r.Agent, r.Url)
}

type Referred struct {
	Id               int64
	Agent            string `json:"agent"`
	Url              string `json:"url"`
	RequestReferer   string `json:"referer"`
	RequestIp        string `json:"ip"`
	RequestUserAgent string `json:"user_agent"`
	CreatedAt        string `json:"cread_at"`
}
