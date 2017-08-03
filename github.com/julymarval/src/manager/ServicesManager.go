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
	"github.com/julymarval/services/model"

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
	claims["exp"] = time.Now().Add(time.Hour).Unix()
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
func CreateAccount(data []byte) (string, string) {

	var user User
	var bol = true

	json.Unmarshal(data, &user)

	bol = model.ValidateEmail(user.Email)

	if !bol {
		response := DoGeneralJSONResponse(model.InvalidInputCode, model.InvalidInputMsg+": Email")
		return response, "error"
	}

	search := SelectDataByID(user.Email)

	if search != "" {
		response := DoGeneralJSONResponse(model.UserAlreadyExistCode, model.UserAlreadyExistMsg)
		return response, "error"
	}

	pass := parsePassword(user.Password)

	if pass == "error" {
		response := DoGeneralJSONResponse(model.InternalErrorCode, model.InternalErrorMsg)
		return response, "error"
	}

	user.Password = pass

	j, err := json.Marshal(user)

	if err != nil {
		response := DoGeneralJSONResponse(model.InternalErrorCode, model.InternalErrorMsg)
		return response, "error"
	}

	result := InsertData(j)

	if result != "ok" {
		response := DoGeneralJSONResponse(model.InternalErrorCode, model.InternalErrorMsg)
		return response, "error"
	}

	response := DoGeneralJSONResponse(model.OkCode, model.OkMsg)
	return response, "ok"

}

// Login = func that manages the login service
func Login(username string, password string) (string, string) {

	var user User
	var tok string
	var bol = true
	var isPass = true

	bol = model.ValidateEmail(username)

	if !bol {
		response := DoGeneralJSONResponse(model.InvalidInputCode, model.InvalidInputMsg+": Email")
		return response, "error"
	}

	result := SelectDataLogin(username)

	if result == "" {
		response := DoGeneralJSONResponse(model.NonExistingUserCode, model.NonExistingUserMsg)
		return response, "error"
	}

	resp := json.RawMessage(result)

	bytes, err := resp.MarshalJSON()

	if err != nil {
		response := DoGeneralJSONResponse(model.InternalErrorCode, model.InternalErrorMsg)
		return response, "error"
	}

	err = json.Unmarshal(bytes, &user)

	isPass = unParsePassword([]byte(password), []byte(user.Password))

	if !isPass {
		response := DoGeneralJSONResponse(model.InvalidPasswordCode, model.InvalidPasswordMsg)
		return response, "error"
	}

	if err != nil {
		response := DoGeneralJSONResponse(model.InternalErrorCode, model.InternalErrorMsg)
		return response, "error"
	}

	tok, err = generateSessionToken(username)

	if err != nil {
		response := DoGeneralJSONResponse(model.InternalErrorCode, model.InternalErrorMsg)
		return response, "error"
	}

	response := DoDataJSONResponse(model.OkCode, model.OkMsg, tok)
	return response, "ok"

}

// GetUser = func that get a user from database using session token
func GetUser(token string) (string, string) {

	var bol = true
	var username string

	bol, username = validateSessionToken(token)

	if !bol {
		response := DoGeneralJSONResponse(model.InvalidTokenCode, model.InvalidTokenMsg)
		return response, "error"
	}

	result := SelectDataByID(username)

	if result == "" {
		response := DoGeneralJSONResponse(model.NonExistingUserCode, model.NonExistingUserMsg)
		return response, "error"
	}

	response := DoDataJSONResponse(model.OkCode, model.OkMsg, result)

	return response, "ok"
}

// ResetPassword = func that updates a pass using session token
func ResetPassword(token string, pass string, confirm string) (string, string) {

	var bol = true
	var username string

	if pass != confirm {
		response := DoGeneralJSONResponse(model.PasswordMismatchCode, model.PasswordMismatchMsg)
		return response, "error"
	}

	bol, username = validateSessionToken(token)

	if !bol {
		response := DoGeneralJSONResponse(model.InvalidTokenCode, model.InvalidTokenMsg)
		return response, "error"
	}

	result := SelectDataByID(username)

	if result == "" {
		response := DoGeneralJSONResponse(model.NonExistingUserCode, model.NonExistingUserMsg)
		return response, "error"
	}

	password := parsePassword(pass)

	if password == "error" {
		response := DoGeneralJSONResponse(999, "InternalError")
		return response, "error"
	}

	res := UpdateDataByID(username, "password", password)

	if res != "ok" {
		return res, "error"
	}

	return res, "ok"

}
