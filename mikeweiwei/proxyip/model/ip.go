package model

type Ip struct {
	Ip string
	IpType string
}

func Newip() *Ip {
	return &Ip{}
}