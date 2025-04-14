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
func (a *Account) SetUsername(username string) {
	a.username = username
}
func (a *Account) GetUsername() string {
	return a.username
}
func (a *Account) SetAccountName(account_name string) {
	a.accountName = account_name
}
func (a *Account) GetAccountName() string {
	return a.accountName
}
func (a *Account) SetPassword(password string) {
	a.password = password
}
func (a *Account) GetPassword() string {
	return a.password
}
func (a *Account) SetSalt(salt []byte) {
	a.salt = salt
}
func (a *Account) GetSalt() []byte {
	return a.salt
}
func (a *Account) SetKey(secret_key string) {
	a.secret_key = secret_key
}
func (a *Account) GetKey() string {
	return a.secret_key
}
func (a *Account) GetSite() string {
	return a.site
}
func (a *Account) SetSite(site string) {
	a.site = site
}

type Website struct {
	Site string `json:"site"`
	Url  string `json:"url"`
}

// Konstroktors
func CreateWebsite() *Website { //rādītājs, saglabā atmiņas adresi(dereferencētu rādītāju, ļaujot piekļūt vērtībai šajā atmiņas adresē). atgriež rādītāju uz jaunu Website struktūru
	return &Website{} //iegūst mainīga adresi atmiņas adresi. atgriež rādītāju uz jauniem `Website` objektiem
}

func (w *Website) GetSite() string {
	return w.Site
}
func (w *Website) SetSite(site string) {
	w.Site = site
}

func (w *Website) GetURL() string {
	return w.Url
}

func (w *Website) SetURL(url string) {
	w.Url = url
}
