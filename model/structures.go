package model

type Account struct {
	username     string
	account_name string
	password     string
	salt         string
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
