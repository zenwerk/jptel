package jptel

import "fmt"

// JPTelephoneNumber は電話番号1件を表す
type JPTelephoneNumber struct {
	AreaCode       string
	CityCode       string
	SubscriberCode string
}

// Format は電話番号をハイフン区切りの文字列にして返す
func (n *JPTelephoneNumber) Format() string {
	return fmt.Sprintf("%s-%s-%s", n.AreaCode, n.CityCode, n.SubscriberCode)
}
