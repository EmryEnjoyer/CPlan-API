package dto

type RefreshTokenRequest struct {
	GrantType    string `url:"grant_type"`
	ClientId     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	RefreshToken string `url:"refresh_token"`
}
