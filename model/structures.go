package model

type Account struct {
	Username     string
	Account_name string
	Password     string
	Salt         string
}

type website struct {
	site string
	url  string
}

func CreateWebsite(site, email string) *website {
	return &website{}
}

func (w *website) GetSite() string {
	return w.site
}
func (w *website) SetSite() string {
	w.site = 
}