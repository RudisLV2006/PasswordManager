package model

type Account struct {
	Username    string `json:"username"`
	Site        string `json:"site"`
	AccountName string /* `json:"account_name"` */
	Password    string `json:"password"`
	Secret_key  string `json:"secretKey"`
	salt        []byte
}

func (a *Account) SetSalt(salt []byte) {
	a.salt = salt
}
func (a *Account) GetSalt() []byte {
	return a.salt
}

func CreateAccount() *Account {
	return &Account{}
}

type Website struct {
	Site string         `json:"site"`
	Url  string `json:"url"`
}

// Konstroktors
func CreateWebsite() *Website { //rādītājs, saglabā atmiņas adresi(dereferencētu rādītāju, ļaujot piekļūt vērtībai šajā atmiņas adresē). atgriež rādītāju uz jaunu Website struktūru
	return &Website{} //iegūst mainīga adresi atmiņas adresi. atgriež rādītāju uz jauniem `Website` objektiem
}
