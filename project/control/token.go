package control

import (
	"github.com/dgrijalva/jwt-go"
	"project/database"
	"project/models"
	"project/configs"
	"project/validation"
	"net/http"
	"errors"
	"time"
	"log"
)

// Senha da criptografia
var jwtKey = []byte("my_secret_key")

// 'Reclames' customizados uteis a nossa aplicação
type UserClaims struct {
	UserEmail string `json:"useremail"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request){
	
	// Instanciar o modelo a ser usado
	var creds models.SecureUser

	// Pegar o JSON da requisição
	body:= r.Body

	// Converter o JSON no modelo criado
	err := JsonStruct(&creds,body)
	if err != nil {
		log.Printf(configs.UNMARSHAL_ERROR+"%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(configs.RESPONSE_UNMARSHAL))
		return
	}

	// Validar o modelo preenchido
	err = validation.Validator.Struct(creds)
	if err != nil {
		log.Printf(configs.VALIDATION_ERROR+"%v\n", err)
		w.Write([]byte(configs.RESPONSE_VALIDATION))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Buscar usuário do modelo 
	user, err := database.SearchUserEmail(creds.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.RESPONSE_DATABASE))
		return
	}

	// Verificar se a senha passada coincide com a do usuário encontrado
	if user.Pass != creds.Pass {
		log.Println("Password not match")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Atribuir um tempo de sessão
	expirationTime := time.Now().Add(5 * time.Minute)

	// Instanciar os 'reclames' a ser usuado no token
	claims := &UserClaims{
		UserEmail: creds.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Criar o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Stringificar o token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.TOKEN_ERROR))
		return
	}

	// Colocar o token no cabeçalho da resposta
	http.SetCookie(w, &http.Cookie{
		Name:    configs.COOKIE_NAME,
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func ValidateLogin (c *http.Cookie) (s string, err error){
	
	// Converter o Cookie em token
	tkn, err := CookieToToken(c)
	if err != nil {
		log.Printf("[ERROR] Failed to parse Cookie: ", err)
		return
	}

	//Se for um token valido retorna o email para as operações
	s, err = TokenEmail(tkn)
	return
}

func CookieToToken (c *http.Cookie) (tkn *jwt.Token, err error){
	// Pegar o token do cookie
	tknStr := c.Value

	// Instanciar um novo 'reclame' neutro
	claims := &UserClaims{}

	// Converter o token passado em um novo token
	tkn, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return 
}

func TokenEmail (tkn *jwt.Token) (s string, err error){
	if claims, ok := tkn.Claims.(*UserClaims); ok && tkn.Valid {
        // Passar o email do 'reclame' do token
		s = claims.UserEmail
    }else {
    	err = errors.New("Not possible take the email")
    }
    return
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	
	c, err := r.Cookie(configs.COOKIE_NAME)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	tkn, err := CookieToToken(c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims, ok := tkn.Claims.(*UserClaims)
	if !ok || !tkn.Valid{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    configs.COOKIE_NAME,
		Value:   tokenString,
		Expires: expirationTime,
	})
}
