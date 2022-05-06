package practiceHttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PracticeHttp() {
	// resp, err := http.Get("http://example.com")
	// if err != nil {
	// 	defer resp.Body.Close()
	// }
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	//url.Parse => 正しいurlか判定
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	/*
		base.ResolveReference => http://example.com + /test?a=1&b=2
		正しいurl表記に変換
	*/
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	//POSTの場合は、nilの箇所にbyteを代入
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match","W/wyzzy")
	//リクエストURLのクエリのデータの中身取出
	q := req.URL.Query()
	fmt.Println(q)
	//クエリに新map配列代入
	q.Add("c","3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	/*
	クエリにmap配列代入 => 新map配列のデータのエンコード必須
	エンコード結果　=> a=1&b=2&c=3%26%25
	3&% 3 => 3 & => %26 % => %25
	*/
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	//実際にURLにアクセス
	resp, _ := client.Do(req)
	//レスポンスのbody読込
	body, _ := ioutil.ReadAll(resp.Body)
	//bodyをstringにキャストして出力
	fmt.Println(string(body))
}
