package usermanager

import (
	"fmt"
	"testing"
)

func TestApiRequest_Param(t *testing.T) {
	testCases := []struct {
		params map[string]string
		key    string
		value  string
	}{
		{map[string]string{"hello": "test"}, "hello", "test"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("value for key '%s'", tc.key), func(t *testing.T) {
			apireq := ApiRequest{params: tc.params}
			if got, _ := apireq.Param(tc.key); got != tc.value {
				t.Errorf("got %s; want %s", got, tc.value)
			}
		})
	}
}
