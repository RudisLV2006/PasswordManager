package model

type Account struct {
	username    string `json:"username"`
	site        string `json:"site"`
	accountName string `json:"account_name"`
	password    string `json:"password"`
	secret_key  string `json:"secretKey"`
	salt        []byte
}

func CreateAccount() *Account {
	return &Account{}
}

type Website struct {
	Site string `json:"site"`
	Url  string `json:"url"`
}

// Konstroktors
func CreateWebsite() *Website { //rādītājs, saglabā atmiņas adresi(dereferencētu rādītāju, ļaujot piekļūt vērtībai šajā atmiņas adresē). atgriež rādītāju uz jaunu Website struktūru
	return &Website{} //iegūst mainīga adresi atmiņas adresi. atgriež rādītāju uz jauniem `Website` objektiem
}
