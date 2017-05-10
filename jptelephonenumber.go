package jptel

import "fmt"

// JPTelephoneNumber represents japanese telephone number
type JPTelephoneNumber struct {
	AreaCode       string
	CityCode       string
	SubscriberCode string
}

// Format returns phone number separated by hyphens
func (n *JPTelephoneNumber) Format() string {
	return fmt.Sprintf("%s-%s-%s", n.AreaCode, n.CityCode, n.SubscriberCode)
}
