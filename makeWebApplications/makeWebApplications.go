package makeWebApplications

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	//0600 => Webサーバー立ち上げユーザー読み書き可能パーミッション
	return ioutil.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	//上記のテンプレートを使用し、ResponseWriterにページの情報記述
	t.Execute(w, p)


}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/test <= 左記の場合、title変数でURLのPath取得可能
	//titleには、/view/以降の文字列を代入したい為、len(/view/)で文字列の長さ取得し、:で先程の文字列以降の文字列を取得する
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func MakeWebApplications() {
	//-----START-----
	// p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
	// p1.save()
	
	// p2, _ := loadPage(p1.Title)
	// fmt.Println(string(p2.Body)
	//-----END-----

	//~ Webサーバー立ち上げ編 ~

	// /view/が来た場合、viewHandlerメソッド呼び出して下さい
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	/*
		:8080 <= Webサーバーを立ち上げるポート記述
		:　<= コロンより前に未記述 => localhost
		nil記述　=> /でアクセス時デフォルトのハンドラーが Page Not Foundを返却
		ListenAndServe => サーバー動作中エラーが発生　=> エラー返却
		log.Fatalでラップした意味 => エラーが発生後、自動でhttp通信を終了する為
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}
