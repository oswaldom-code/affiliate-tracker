package affiliate_tracker

import "fmt"

type ReferralLink struct {
	Agent string `json:"agent"`
	Url   string `json:"url"`
}

func (r *ReferralLink) ToString() string {
	return fmt.Sprintf("%s:::%s", r.Agent, r.Url)
}

type ReferredDto struct {
	Id        int64
	Agent     string `json:"agent"`
	Url       string `json:"url"`
	CreatedAt string `json:"cread_at"`
}
