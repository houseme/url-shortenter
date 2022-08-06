package utility

import (
	"context"
	"encoding/hex"
	"reflect"
	"testing"
)

func Test_utilHelper_AESEncrypt(t *testing.T) {
	type args struct {
		ctx  context.Context
		key  []byte
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantDst string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ctx:  context.Background(),
				key:  []byte("kNsSjdNZDvXkXch2"),
				data: []byte("jNvDQnU8VLvXdRPytEBBTPmrekmZDZdaNKyX3SEPbEunWGYb"),
			},
			wantDst: "a87vEpKfvwsgk5+80Cmx5ip9do3O3SZ1h6Gd5EXGnkDwBCDbjUlQPQlEJwTdRMUHexEyan6orRPKM3eu/hMBMg==",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &utilHelper{}
			gotDst, err := u.AESEncrypt(tt.args.ctx, tt.args.key, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDst != tt.wantDst {
				t.Errorf("AESEncrypt() gotDst = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

// Benchmark_utilHelper_AESEncrypt-4   	 1000000	      904 ns/op	     832 B/op	       1 allocs/op
func Benchmark_utilHelper_AESEncrypt(b *testing.B) {
	ctx := context.Background()
	key := []byte("kNsSjdNZDvXkXch2")
	data := []byte("jNvDQnU8VLvXdRPytEBBTPmrekmZDZdaNKyX3SEPbEunWGYb")
	u := &utilHelper{}
	for i := 0; i < b.N; i++ {
		_, err := u.AESEncrypt(ctx, key, data)
		if err != nil {
			b.Error(err)
		}
	}
}

func Test_utilHelper_Sha256OfShort(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				input: "jNvDQnU8VLvXdRPytEBBTPmrekmZDZdaNKyX3SEPbEunWGYb",
			},
			want:    []byte("89581823aea94b230ccb5b9b0dcf92fb8aa77447fddd113708dcc1accd9b1ae9"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &utilHelper{}
			got, err := u.Sha256OfShort(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha256OfShort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(hex.EncodeToString(got), string(tt.want)) {
				t.Errorf("Sha256OfShort() got = %x, want %x", got, tt.want)
			}
		})
	}
}

func Test_utilHelper_PasswordBase58Hash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				password: "321003061742927872",
			},
			want:    "HAELVZ5n7T8VMDRUZQ62wPEFidmo5Gnj8qKKvV82rxbQ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &utilHelper{}
			got, err := u.PasswordBase58Hash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswordBase58Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PasswordBase58Hash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_utilHelper_CreateAccessToken(t *testing.T) {
	type args struct {
		ctx       context.Context
		accountNo uint64
	}
	tests := []struct {
		name      string
		args      args
		wantToken string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ctx:       context.Background(),
				accountNo: 1,
			},
			wantToken: "8687027c3411afd9069b56d72ea0e795e55ac30441f6dfa8f06385d360948b2d",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &utilHelper{}
			gotToken, err := u.CreateAccessToken(tt.args.ctx, tt.args.accountNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken == tt.wantToken {
				t.Errorf("CreateAccessToken() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
			t.Log("getToken:", gotToken)
		})
	}
}

// Benchmark_utilHelper_CreateAccessToken-4   	 1000000	      904 ns/op	     832 B/op	       1 allocs/op
func Benchmark_utilHelper_CreateAccessToken(b *testing.B) {
	ctx := context.Background()
	accountNo := uint64(1)
	u := &utilHelper{}
	for i := 0; i < b.N; i++ {
		_, err := u.CreateAccessToken(ctx, accountNo)
		if err != nil {
			b.Error(err)
		}
	}
}
