package model

import (
	"testing"
)

func Test_intDB(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "name",wantErr: false},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := intDB(); (err != nil) != tt.wantErr {
				t.Errorf("intDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}