package resolver

import (
	"reflect"
	"testing"

	"github.com/myminicommission/api/graph/model"
)

func TestGetEstimate(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Estimate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEstimate(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEstimate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEstimate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEstimates(t *testing.T) {
	tests := []struct {
		name    string
		want    []*model.Estimate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEstimates()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEstimates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEstimates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMinis(t *testing.T) {
	tests := []struct {
		name    string
		want    []*model.Mini
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMinis()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMinis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMinis() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEstimates(t *testing.T) {
	tests := []struct {
		name string
		want []*model.Estimate
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEstimates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEstimates() = %v, want %v", got, tt.want)
			}
		})
	}
}
