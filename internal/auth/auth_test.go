package auth

import (
	"fmt"
	"net/http"
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
		"simple":           {key: "Authorization", value: "ApiKey xxxxxx", want: "xxxxxx"},
		"incorrect header": {key: "Bearer", value: "ApiKey xxxxxx", wantErr: "no authorization header included"},
		"no header":        {wantErr: "no authorization header included"},
	}

	for name, tc := range tests {

		t.Run(fmt.Sprintf("TestGetAPIKey Case %v", name), func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tc.wantErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != tc.want {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})

	}
}
