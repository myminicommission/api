package resolver

import (
	"reflect"
	"testing"

	"github.com/myminicommission/api/graph/model"
)

func TestGetCommission(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Commission
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommission(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommission() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommissions(t *testing.T) {
	tests := []struct {
		name    string
		want    []*model.Commission
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommissions()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommissions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
