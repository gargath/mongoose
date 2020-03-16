package entities

type User struct {
	Sub        string `json:"sub" bson:"sub"`
	Name       string `json:"name" bson:"name"`
	FamilyName string `json:"family_name" bson:"family_name"`
	GivenName  string `json:"given_name" bson:"given_name"`
	Token      Token  `json:"token" bson:"token"`
}

type Token struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	Expiry       string `json:"expiry" bson:"expiry"`
	TokenType    string `json:"token_type" bson:"token_type"`
}
