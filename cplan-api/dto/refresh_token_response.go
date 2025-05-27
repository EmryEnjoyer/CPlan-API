package dto

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}
