package dto

type UserProfileResponse struct {
	Sub        string `json:"sub"`
	Nickname   string `json:"nickname"`
	Name       string `json:"name"`
	PictureUrl string `json:"picture"`
	Updated_at string `json:"updated_at"`
	Email      string `json:"email"`
	Verified   bool   `json:"email_verified"`
}
