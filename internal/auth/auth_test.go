package auth

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	tests := map[string]struct {
		key     string
		value   string
		want    string
		wantErr string
	}{
		"simple":           {key: "Authorization", value: "XXXXXXXX", want: "XXXXXXXX", wantErr: ""},
		"incorrect header": {key: "Bearer", value: "XXXXXXX", want: "", wantErr: "malformed authorization header"},
		"no header":        {key: "", value: "", want: "", wantErr: "no authorization header included"},
	}

	for name, tc := range tests {
		header := http.Header{}
		header.Add(tc.key, tc.value)

		got, err := GetAPIKey(header)
		if err != nil {
			if strings.Contains(err.Error(), tc.wantErr) {
				return
			}
			t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
			return
		}
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
			return
		}
	}
}
