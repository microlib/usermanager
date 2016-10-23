package usermanager

import "errors"

type ApiRequestInterface interface {
	Params() map[string]string
	Param(param string) (string, error)
}

type ApiRequest struct {
	params map[string]string
}

func (a *ApiRequest) Params() map[string]string {
	return a.params
}

func (a *ApiRequest) Param(param string) (string, error) {
	v, ok := a.params[param]
	if !ok {
		return "", errors.New("Value not found")
	}
	return v, nil
}

