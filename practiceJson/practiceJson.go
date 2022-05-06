package practiceJson

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Person struct {
	//struct => jsonに変換　=> リテラル（"Name" => "name"）を指定可能
	//omitempty => 配列のvalueが空の場合、該当のjsonを非表示
	Name      string   `json:"name"`
	Age       int      `json:"age,omitempty"`
	NickNames []string `json:"nicknames"`
	T         *T       `json:"T,omitempty"`
}

type T struct{}

//jsonに変換時、以下のメソッドでカスタマイズ可能
func (p Person) MarshalJSON() ([]byte, error) {
	//下記の様に１行でstructを初期化可能
	// a := &struct{ Name string }{Name: "Jon"}
	v, err := json.Marshal(&struct {
		Name string
	}{Name: "Mr." + p.Name})
	return v, err
}

//json => struct (Unmarshal)時、以下のメソッドでカスタマイズ可能
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

func PracticeJson() {
	//byte配列のKey大文字or小文字　=> スペルOK => structに代入可能
	b := []byte(`{"name" : "Mike" , "age" : 0, "nicknames" : ["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}

	fmt.Println(p.Name, p.Age, p.NickNames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))

}

//--------------------hmac--------------------


var DB = map[string]string {
	"User1Key" : "User1Secret",
	"User2Key" : "User2Secret",
}

//サーバー側処理
func Sever (apikey,sign string,data []byte) {
	apiSecret := DB[apikey]
	//DB格納apiSecretをbyteに変換しハッシュ生成
	h := hmac.New(sha256.New,[]byte(apiSecret))
	//生成したハッシュにクライアントデータ書込
	h.Write(data)
	//生成したサーバーのハッシュをEncodeToString => ハッシュ表示可能に
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	//クライアントのハッシュとサーバー側のハッシュ一致確認
	fmt.Println(sign == expectedHMAC)
}

//クライアント側処理
func PracticeHmac() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	data := []byte("data")
	//sha256のアルゴリズムを使用し、apiSecretをbyteに変換しハッシュを生成
	h := hmac.New(sha256.New,[]byte(apiSecret))
	//生成したハッシュ　<= 送信data書込
	h.Write(data)
	//生成したハッシュをEncodeToString => ハッシュ表示可能に
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)

	//クライアント送信データ => サーバー側で正か確認
	Sever(apiKey,sign,data)
}
