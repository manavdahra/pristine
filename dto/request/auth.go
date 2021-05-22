package request

type LoginReq struct {
	IdToken string `json:"idToken"`
	OrgName string `json:"orgName"`
}
