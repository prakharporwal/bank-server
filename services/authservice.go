package services

type Token string

type Auth interface {
	LoginService(string, string) *LoginResponse
	generateToken(string, string) Token
	RefreshToken(Token, string) Token
}

type LoginResponse struct {
	Token Token `json:"token"`
}

func LoginService(username string, password string) *LoginResponse {

	token := generateToken(username, password)
	return &LoginResponse{
		Token: token,
	}
}

func generateToken(username string, password string) Token {
	// JWTService
	return "hello"
}

type SignUpResponse struct {
	Token      Token `json:"token"`
	IsVerified bool  `json:"is_verified"`
}

func SignUpService() *SignUpResponse {
	return &SignUpResponse{
		Token:      "hello",
		IsVerified: false,
	}
}
