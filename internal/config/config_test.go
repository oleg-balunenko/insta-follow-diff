package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	path string
}

type test struct {
	name    string
	args    args
	want    Config
	wantErr bool
}

func TestLoad(t *testing.T) {
	tests := []test{
		{
			name: "load config from file",
			args: args{
				path: filepath.Clean(filepath.Join("testdata", "config-test.json")),
			},
			want: Config{
				user: user{
					instagram: instagram{
						username: "user",
						password: "pass",
					},
				},
				db: db{
					local:               true,
					mongoURL:            "mongoURL:test",
					mongoDBName:         "testing",
					mongoCollectionName: "users",
				},
				whitelist: []string{
					"user1",
					"user2",
					"user3",
				},
				limits: limits{
					unfollow: 100,
				},
				debug: false,
			},
			wantErr: false,
		},
		{
			name: "error for not exist file",
			args: args{
				path: filepath.Clean(filepath.Join("testdata", "config-test-not-exist.json")),
			},
			want:    Config{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.path)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
