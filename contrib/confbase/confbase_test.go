package confbase_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/config"

	"github.com/sendya/kratos-helper/contrib/confbase"
)

func TestDecoder(t *testing.T) {
	type args struct {
		src    *config.KeyValue
		target map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test failed format",
			args: args{
				src: &config.KeyValue{
					Key: "server.http.timeout",
					Format: "yaml~",
					Value: []byte(""),
				},
				target: map[string]interface{}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := confbase.Decoder(tt.args.src, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("Decoder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
