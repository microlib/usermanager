package usermanager

type ApiRequest interface {
	Params() []string
}