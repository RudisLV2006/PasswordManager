package model

type Account struct {
	username     string
	site         string
	account_name string
	password     string
	secret_key   string
	salt         []byte
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
	a.account_name = account_name
}
func (a *Account) GetAccountName() string {
	return a.account_name
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
	site string
	url  string
}

// Konstroktors
func CreateWebsite() *Website { //rādītājs, saglabā atmiņas adresi(dereferencētu rādītāju, ļaujot piekļūt vērtībai šajā atmiņas adresē). atgriež rādītāju uz jaunu Website struktūru
	return &Website{} //iegūst mainīga adresi atmiņas adresi. atgriež rādītāju uz jauniem `Website` objektiem
}

func (w *Website) GetSite() string {
	return w.site
}
func (w *Website) SetSite(site string) {
	w.site = site
}

func (w *Website) GetURL() string {
	return w.url
}

func (w *Website) SetURL(url string) {
	w.url = url
}
