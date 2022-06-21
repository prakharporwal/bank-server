package token

type Auth interface {
	//Login()
	//SignUp()
	GenerateToken()
	GetUserIdFromToken()
	ValidateToken()
	RefreshToken()
}

func GenerateToken() {

}

func RefreshToken() {

}
