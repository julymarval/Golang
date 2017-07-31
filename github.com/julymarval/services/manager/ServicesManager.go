/********************************************
* File that handles services requests    	*
* @author: Julyamnis Marval        			*
* @version: 1.0                    			*
********************************************/

package manager

import (
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

/******************************** GLOBAL VARIABLES ******************************/

// change this for a secret word
// you can have it in a config file.
var tokenEncodeString = "MYSAFEWORD"

/************************************ STRUCTS ***********************************/

// User = Struct
type User struct {
	ID          bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	Name        string        `json:"name"`
	LastName    string        `json:"lastName"`
	Email       string        `json:"email"`
	DateofBirth string        `json:"dateOfBirth"`
	Password    string        `json:"password,omitempty"`
}

/************************************ FUNCS ***********************************/

// parsePassword = func that generates a user password parsed
func parsePassword(pass string) string {

	password := []byte(pass)

	// Hashing the password with default cost
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return "error"
	}

	return string(hashedPassword)
}

func unParsePassword(pass []byte, hash []byte) bool {

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword(hash, pass); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// generateSessionToken = func that generates a session token
func generateSessionToken(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS512)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 7).Unix()
	claims["username"] = username
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(tokenEncodeString))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateSessionToken(Token string) (bool, string) {

	keyLookupFn := func(token *jwt.Token) (interface{}, error) {
		// Check for expected signing method.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenEncodeString), nil
	}

	// Parse and validate the token.
	token, err := jwt.Parse(Token, keyLookupFn)
	if err != nil {
		fmt.Println(err)
		return false, ""
	} else if !token.Valid {
		return false, ""
	}

	claims := token.Claims.(jwt.MapClaims)

	return true, claims["username"].(string)

}

// CreateAccount = func thats manages user's data and creates an account
func CreateAccount(data []byte) string {

	var user User

	json.Unmarshal(data, &user)

	search := SelectDataByID(user.Email)

	if search != "" {
		response := DoGeneralJSONResponse(104, "UserAlreadyExists")
		return response
	}

	pass := parsePassword(user.Password)

	if pass == "error" {
		response := DoGeneralJSONResponse(999, "InternalError")
		return response
	}

	user.Password = pass

	j, err := json.Marshal(user)

	if err != nil {
		response := DoGeneralJSONResponse(999, "InternalError")
		return response
	}

	result := InsertData(j)

	if result != "ok" {
		response := DoGeneralJSONResponse(999, "InternalErrorok")
		return response
	}

	response := DoGeneralJSONResponse(0, "ok")
	return response

}

// Login = func that manages the login service
func Login(username string, password string) string {

	var user User
	var tok string
	var bol = true

	result := SelectDataLogin(username)

	resp := json.RawMessage(result)

	bytes, err := resp.MarshalJSON()

	if err != nil {
		response := DoGeneralJSONResponse(999, "InternalError")
		return response
	}

	err = json.Unmarshal(bytes, &user)

	bol = unParsePassword([]byte(password), []byte(user.Password))

	if !bol {
		response := DoGeneralJSONResponse(100, "InvalidPassword")
		return response
	}

	if err != nil {
		response := DoGeneralJSONResponse(999, "InternalError")
		return response
	}

	tok, err = generateSessionToken(username)

	if err != nil {
		response := DoGeneralJSONResponse(999, "InternalError")
		return response
	}

	response := DoDataJSONResponse(0, "ok", tok)
	return response

}

// GetUser = func that get a user from database using session token
func GetUser(token string) string {

	var bol = true
	var username string

	bol, username = validateSessionToken(token)

	if !bol {
		response := DoGeneralJSONResponse(003, "InvalidSessionToken")
		return response
	}

	result := SelectDataByID(username)

	response := DoDataJSONResponse(0, "ok", result)

	return response
}

// ResetPassword = func that updates a pass using session token
func ResetPassword(token string, pass string, confirm string) string {

	var bol = true
	var username string

	if pass != confirm {
		response := DoGeneralJSONResponse(106, "PasswordsMismatch")
		return response
	}

	bol, username = validateSessionToken(token)

	if !bol {
		response := DoGeneralJSONResponse(003, "InvalidSessionToken")
		return response
	}

	result := SelectDataByID(username)

	if result == "" {
		response := DoGeneralJSONResponse(105, "NonExistingUser")
		return response
	}

	password := parsePassword(pass)

	if password == "error" {
		response := DoGeneralJSONResponse(999, "InternalError")
		return response
	}

	res := UpdateDataById(username, "password", password)

	return res

}
