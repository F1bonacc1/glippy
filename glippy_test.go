package glippy

import (
	"context"
	"testing"
)

func TestSetGet(t *testing.T) {
	_, err := Get()
	if err != nil && err.Error() == "no X11 connection available" {
		t.Skip("no X11 connection available, skipping tests")
	}

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "abc",
			want:    "abc",
			wantErr: false,
		},
		{
			name:    "empty",
			want:    "",
			wantErr: false,
		},
		{
			name:    "emoji",
			want:    "🔥",
			wantErr: false,
		},
		{
			name:    "utf8",
			want:    "aš tave myliu",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.want)
			got, err := Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatch(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name: "abc",
			want: "abc",
		},
		{
			name: "empty",
			want: "",
		},
		{
			name: "emoji",
			want: "🔥",
		},
		{
			name: "utf8",
			want: "aš tave myliu",
		},
	}

	ch := Watch(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Set(tt.want)
			if err != nil {
				return
			}
			got := <-ch
			if got != tt.want {
				t.Errorf("Watch() = %v, want %v", got, tt.want)
			}
		})
	}
}
