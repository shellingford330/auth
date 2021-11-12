package config

import (
	"os"
	"testing"
)

func TestConfiguration(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		env    string
		method func(c *Configuration) interface{}
		want   string
	}{
		{
			name:   "ENV",
			env:    "test",
			method: func(c *Configuration) interface{} { return c.Env() },
			want:   "test",
		},
		{
			name:   "HTTP_SERVER_HOST",
			env:    "host",
			method: func(c *Configuration) interface{} { return c.HTTPServerHost() },
			want:   "host",
		},
		{
			name:   "HTTP_SERVER_PORT",
			env:    "80",
			method: func(c *Configuration) interface{} { return c.HTTPServerPort() },
			want:   "80",
		},
		{
			name:   "GRPC_SERVER_HOST",
			env:    "localhost",
			method: func(c *Configuration) interface{} { return c.GRPCServerHost() },
			want:   "localhost",
		},
		{
			name:   "GRPC_SERVER_PORT",
			env:    "50051",
			method: func(c *Configuration) interface{} { return c.GRPCServerPort() },
			want:   "50051",
		},
		{
			name:   "DB_NAME",
			env:    "test_db",
			method: func(c *Configuration) interface{} { return c.DBName() },
			want:   "test_db",
		},
		{
			name:   "DB_HOST",
			env:    "test-server",
			method: func(c *Configuration) interface{} { return c.DBHost() },
			want:   "test-server",
		},
		{
			name:   "DB_PORT",
			env:    "13306",
			method: func(c *Configuration) interface{} { return c.DBPort() },
			want:   "13306",
		},
		{
			name:   "DB_USER",
			env:    "test_user",
			method: func(c *Configuration) interface{} { return c.DBUser() },
			want:   "test_user",
		},
		{
			name:   "DB_PASSWORD",
			env:    "test_password",
			method: func(c *Configuration) interface{} { return c.DBPassword() },
			want:   "test_password",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := os.Setenv(tt.name, tt.env)
			if err != nil {
				t.Fatal("failed to set environment variable: ", err)
			}
			defer func() {
				_ = os.Unsetenv(tt.name)
			}()
			c, err := NewConfiguration()
			if err != nil {
				t.Error("failed to init configuration: ", err)
			}
			got := tt.method(c)
			if got != tt.want {
				t.Errorf("not expected environment value: %v", err)
			}
		})
	}
}
