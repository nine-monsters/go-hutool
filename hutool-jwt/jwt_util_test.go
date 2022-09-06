package jwt

import (
	"github.com/nine-monsters/go-hutool/hutool-jwt/signer"
	"reflect"
	"testing"
)

func TestCreateToken(t *testing.T) {
	type args struct {
		payload map[string]any
		key     []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{payload: map[string]any{"uid": 123, "expire_time": 1000 * 60 * 60 * 24 * 15}, key: []byte("123")}, "eyJhbGciOiJIUzI1NiJ9.eyIxIjoiMSJ9.gi00We2cXmQSIhmXdssCBYLbsRD2ZCjwzOeE4q4qKwA="},
		{"test2", args{payload: map[string]any{"typ": "JWT"}, key: []byte("1234567890")}, "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiYWRtaW4iOnRydWUsIm5hbWUiOiJsb29seSJ9.U2aQkC2THYV9L0fTN-yBBI7gmo5xhmvMhATtu8v0zEA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateToken(tt.args.payload, tt.args.key); got != tt.want {
				t.Errorf("CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want *Jwt
	}{
		{name: "test1", args: args{
			token: "eyJhbGciOiJIUzI1NiJ9.eyIxIjoiMSJ9.gi00We2cXmQSIhmXdssCBYLbsRD2ZCjwzOeE4q4qKwA=",
		}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		token string
		key   []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{
			token: "eyJhbGciOiJIUzI1NiJ9.eyIxIjoiMSJ9.gi00We2cXmQSIhmXdssCBYLbsRD2ZCjwzOeE4q4qKwA=",
			key:   []byte("123"),
		}, want: true},
		{name: "test2", args: args{
			token: "eyJhbGciOiJIUzI1NiJ9.eyIxIjoiMSJ9.gi00We2cXmQSIhmXdssCBYLbsRD2ZCjwzOeE4q4qKwA=",
			key:   []byte("1233"),
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Verify(tt.args.token, tt.args.key); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTokenWithSigner(t *testing.T) {
	type args struct {
		payload map[string]any
		signer  signer.Signer
	}
	hs384 := signer.HS384([]byte("123"))
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{payload: map[string]any{"uid": 123, "expire_time": 1000 * 60 * 60 * 24 * 15}, signer: &signer.NoneSigner{}}, "eyJhbGciOiJIUzI1NiJ9.eyIxIjoiMSJ9.gi00We2cXmQSIhmXdssCBYLbsRD2ZCjwzOeE4q4qKwA="},
		{"test2", args{payload: map[string]any{"typ": "JWT"}, signer: hs384}, "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiYWRtaW4iOnRydWUsIm5hbWUiOiJsb29seSJ9.U2aQkC2THYV9L0fTN-yBBI7gmo5xhmvMhATtu8v0zEA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTokenWithSigner(tt.args.payload, tt.args.signer); got != tt.want {
				t.Errorf("CreateTokenWithSigner() = %v, want %v", got, tt.want)
			}
		})
	}
}
