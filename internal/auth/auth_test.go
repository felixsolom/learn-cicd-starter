package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func stringToHeader(key, value string) http.Header {
	header := http.Header{}
	header.Set(key, value)
	return header
}

func TestAuth(t *testing.T) {
	type test struct {
		name    string
		header  http.Header
		want    string
		wantErr bool
		errMsg  string
	}

	tests := []test{
		{
			name:    "missing header",
			header:  http.Header{},
			want:    "",
			wantErr: true,
			errMsg:  "no authorization header included",
		},
		{
			name:    "not an api key",
			header:  stringToHeader("Authorization", "notApiKey sfgsfgsfklj"),
			want:    "",
			wantErr: true,
			errMsg:  "malformed authorization header",
		},
		{
			name:    "valid api key",
			header:  stringToHeader("Authorization", "ApiKey sodifjdsigosdifjgosdifjfdsg"),
			want:    "sodifjdsigosdifjgosdifjfdsg",
			wantErr: false,
			errMsg:  "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.header)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey() err= %v, wanted error: %v", err, tc.wantErr)
			}

			if tc.wantErr && tc.errMsg != "" && err.Error() != tc.errMsg {
				t.Errorf("GetAPIKey() errMsg= %q, want: %q", err.Error(), tc.errMsg)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetAPIKey() got= %v, want: %v", got, tc.want)
			}
		})
	}
}
