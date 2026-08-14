package model

type Request struct {
	Token string
	Role  string
}

type Claims struct {
	User string
	Role string
}
