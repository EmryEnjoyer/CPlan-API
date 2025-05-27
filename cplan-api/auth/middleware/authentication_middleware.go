package middleware

import (
	"bytes"
	"cplan-api/auth"
	"cplan-api/dto"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/google/go-querystring/query"
)

func TokenIsNotExpired(access_token string) bool {
	parser := new(jwt.Parser)

	token, err := parser.Parse(access_token, nil)
	if err != nil && err.Error() != "no Keyfunc was provided." {
		log.Fatal(err.Error())
	}

	claims, claims_ok := token.Claims.(jwt.MapClaims)
	if !claims_ok {
		log.Fatal("failed to read claims")
	}

	exp, exp_ok := claims["exp"].(float64)
	if !exp_ok {
		log.Fatal("exp claim not present")
	}

	return time.Now().Before(time.Unix(int64(exp), 0))
}

func GetUserProfile(context *gin.Context) {

	session := sessions.Default(context)
	access_token := session.Get("access_token")
	user_profile_client := http.Client{}

	user_profile_url, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + os.Getenv("AUTH0_USER_INFO_ENDPOINT"))

	if err != nil {
		log.Println("Failed to build user profile url " + err.Error())
		context.Abort()
		return
	}

	user_profile_request, user_profile_request_error := http.NewRequest("GET", user_profile_url.String(), nil)
	if user_profile_request_error != nil {
		log.Println("Failed to initialize validation request", user_profile_request_error.Error())
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	user_profile_request.Header.Add("Accept", "application/json")
	user_profile_request.Header.Add("Authorization", "Bearer "+access_token.(string))

	user_profile_response, user_profile_response_err := user_profile_client.Do(user_profile_request)
	if user_profile_response_err != nil {
		log.Println("Failed to validate user")
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	defer user_profile_response.Body.Close()
	user_profile_bytes, _ := io.ReadAll(user_profile_response.Body)

	var user_profile dto.UserProfileResponse
	json.Unmarshal(user_profile_bytes, &user_profile)
	context.Set("user_profile", user_profile)
	context.Next()
}

func HandleRefreshToken(session sessions.Session) bool {
	token_url, token_url_err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + os.Getenv("AUTH0_TOKEN_ENDPOINT"))
	if token_url_err != nil {
		log.Fatal("Failed to parse token url", token_url_err.Error())
	}

	refresh_token := session.Get("refresh_token")
	refresh_request_dto := dto.RefreshTokenRequest{
		GrantType:    "refresh_token",
		ClientId:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RefreshToken: refresh_token.(string),
	}

	refresh_body, marshal_err := query.Values(refresh_request_dto)
	if marshal_err != nil {
		log.Fatal("failed to marshal object: ", marshal_err)
	}

	refresh_client := http.Client{}
	refresh_request, refresh_request_err := http.NewRequest("POST", token_url.String(), bytes.NewReader([]byte(refresh_body.Encode())))
	if refresh_request_err != nil {
		log.Fatal("Failed to create a request: ", refresh_request_err.Error())
	}
	refresh_request.Header.Add("content-type", "application/x-www-form-urlencoded")
	refresh_response, refresh_response_err := refresh_client.Do(refresh_request)
	if refresh_response_err != nil {
		log.Fatal("Failed to fetch refresh token and response: ", refresh_response_err.Error())
	}
	defer refresh_response.Body.Close()

	if refresh_response.StatusCode == http.StatusOK {
		var response_body dto.RefreshTokenResponse
		json_decoder := json.NewDecoder(refresh_response.Body)
		content_unmarshal_err := json_decoder.Decode(&response_body)

		if content_unmarshal_err != nil {
			log.Println("Faileed to unmarshal data")
			return false
		}

		session.Set("access_token", response_body.AccessToken)
		session.Set("refresh_token", response_body.RefreshToken)
		session.Options(sessions.Options{Path: "/"})
		session_save_error := session.Save()

		if session_save_error != nil {
			log.Println("Failed to set session: " + session_save_error.Error())
			return false
		}
		return true
	} else {
		return false
	}
}

func IsAuthenticated(auth *auth.Authenticator) gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)

		if session.Get("profile") == nil {
			context.Redirect(http.StatusSeeOther, "/auth/login")
			context.Abort()
			return
		}

		access_token := session.Get("access_token")

		if access_token == nil {
			context.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		if TokenIsNotExpired(access_token.(string)) {
			context.Next()
		} else {
			if !HandleRefreshToken(session) {
				context.String(http.StatusUnauthorized, "Failed to refresh access token")
			} else {
				context.Next()
			}
		}
	}
}
