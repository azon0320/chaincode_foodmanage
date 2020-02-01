package models

type Credentials struct {
	AccountId string `json:"account_id"`
	Password  string `json:"password"`
	Token string `json:"token"`
}

func CreateCredentialsWithPassword(id string, passwd string) *Credentials{
	return &Credentials{AccountId: id, Password: passwd}
}

func CreateCredentialsWithToken(token string) *Credentials{
	return &Credentials{Token:token}
}