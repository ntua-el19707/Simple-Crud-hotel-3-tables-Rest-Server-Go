package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/scrypt"
)

type password interface {
	genarate(pass string) error
	valid(candidate string) bool
}

type PassHashSalt struct {
	Hash string `json:"hash"`
	Salt string `json:"salt"`
}

// Generate generates the hashed password and salt for a given password.
func (p * PassHashSalt) genarate(pass string)  error {
	password := []byte(pass)
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return err
	}
	
	N := 16384       // CPU/memory cost factor
	r := 8           // Block size
	keyLen := 32     // Length of the derived key

	// Generate the RS356 hash
	hash, err := scrypt.Key(password, salt, N, r, 1, keyLen)
	if err != nil {
		return err
	}

	// Convert the hash and salt bytes to hexadecimal strings
	 p.Hash = hex.EncodeToString(hash)
	p.Salt = hex.EncodeToString(salt)

	return nil
}

// Validate checks if a candidate password matches the stored hash and salt.
func (p *PassHashSalt) valid(candidate string) bool {
	salt ,err := hex.DecodeString(p.Salt)
	if err != nil{
		return false
	}
	candidateHash, err := scrypt.Key([]byte(candidate),salt , 16384, 8, 1, 32)
	if err != nil {
		return false
	}

	return hex.EncodeToString(candidateHash) == p.Hash
}

func genarate(p password ,pass string )(error){
	err := p.genarate(pass)
	return err
}
func valid(p password ,pass string )(bool){
	rsp := p.valid(pass)
	return  rsp
}
type  userPayload struct {
	id int  
	userName string  
}

func issueJwt(ttl time.Duration, content userPayload)(string,  error){
	private_key ,  err :=  readFile("private_key.pem")
	if err != nil {
		return  "" ,  errors.New("Failed to get private key ")
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(private_key))
	if err != nil {
		return  "" ,  errors.New("Failed to get private key ")
	}

    if err != nil {
        return "", err
    }
	now := time.Now().UTC()

	
	claimData := jwt.MapClaims {
		"Id": strconv.Itoa(content.id),
		"ExpiresAt": now.Add(ttl).Unix(),
		"IssuedAt":  now.Unix(),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claimData)
	signedToken, err := token.SignedString(key)

	if err != nil {
		return "", err
	}
 
	return signedToken, nil
}
	type userPayloadSigned struct{
		Id  int 
		ExpiresAt  int64
		IssuedAt    int64
	}
 func validate(token string) (userPayloadSigned  , error) {
	public ,  err :=  readFile("public_key.pem")
	if err != nil {
		return  userPayloadSigned {} ,  errors.New("Failed to get private key ")
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(public))
	if err != nil {
		return userPayloadSigned {} , err
	}
 
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return 0 , err
		}
 
		return key, nil
	})
	if err != nil {
		return userPayloadSigned {} , err
	}
 

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return userPayloadSigned {}, errors.New("Invalid token or claims")
	}

	id, ok1 := claims["Id"].(string)
	ExpiresAt, ok2 := claims["ExpiresAt"].(float64)
	
	IssuedAt, ok3 := claims["IssuedAt"].(float64)

	if !ok1 || !ok2 || !ok3{
		return userPayloadSigned {}, errors.New("failed  to get  sign  info ")
	}
	idInt , err := strconv.Atoi(id)  
	if err != nil {
		return userPayloadSigned {}, errors.New("failed  to convert id string  to int")

	}

   
	userInfo := userPayloadSigned{
		Id: idInt,
		ExpiresAt: int64(ExpiresAt),
		IssuedAt: int64(IssuedAt),
  	 }
   return userInfo, nil 
 }