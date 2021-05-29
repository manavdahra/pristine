package request

type LoginReq struct {
	IdToken string `json:"idToken"`
	OrgName string `json:"orgName"`
}

type LogoutReq struct {
	Email string `json:"email"`
}
