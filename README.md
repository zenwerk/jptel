# jptel 電話番号分割

jptel は日本の電話番号を市外局番・市内局番・加入者番号に分割して返します。

## インストール

```
$ go get github.com/zenwerk/jptel
```

## 使い方

```go
package main

import (
	"fmt"

	"github.com/zenwerk/jptel"
)

func main() {
	result, err := jptel.Split("0312345678")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(result)) // 3
	fmt.Println(result[0])   // 03
	fmt.Println(result[1])   // 1234
	fmt.Println(result[2])   // 5678

	result, err = jptel.Split("090987654323")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(result)) // 3
	fmt.Println(result[0])   // 090
	fmt.Println(result[1])   // 9876
	fmt.Println(result[2])   // 5432
}
```

## その他
固定電話の市外局番データは[総務省のサイト](http://www.soumu.go.jp/main_sosiki/joho_tsusin/top/tel_number/number_shitei.html#kotei-denwa)からダウンロードできるExcelから生成しています。
再生成する場合は以下の手順で行って下さい。

```
$ pip install -r freeze.txt
$ python generate_master_data.py
```

### thanks
- https://gist.github.com/kennyj/4966002