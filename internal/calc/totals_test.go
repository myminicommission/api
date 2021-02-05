package calc

import (
	"testing"
	"time"

	"github.com/myminicommission/api/graph/model"
)

func TestCommissionTotal(t *testing.T) {
	type args struct {
		commission *model.Commission
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "commission with no minis",
			want: 0,
			args: args{
				commission: &model.Commission{
					CreatedAt: time.Now(),
					Minis:     []*model.CommissionedMini{},
				},
			},
		},
		{
			name: "commission with one mini",
			want: 100,
			args: args{
				commission: &model.Commission{
					ID:        "70bfc0e1-7927-49da-83e8-be8157b4a00e",
					CreatedAt: time.Now(),
					Minis: []*model.CommissionedMini{
						{
							ID:       "f96f9fe1-c6fa-463a-9f24-2a86666642c5",
							Quantity: 1,
							Name:     "Silent King",
							Size:     model.MiniSizeLarge,
							Price:    100.00,
						},
					},
				},
			},
		},
		// {
		// 	name: "quote with options",
		// 	want: 111,
		// 	args: args{
		// 		commission: &model.Commission{
		// 			ID:        "70bfc0e1-7927-49da-83e8-be8157b4a00e",
		// 			CreatedAt: time.Now(),
		// 			Minis: []*model.CommissionedMini{
		// 				{
		// 					ID:       "f96f9fe1-c6fa-463a-9f24-2a86666642c5",
		// 					Quantity: 1,
		// 					Name:     "Silent King",
		// 					Size:     model.MiniSizeLarge,
		// 					Price:    100.00,
		// 					Options: []*model.MiniOption{
		// 						{
		// 							ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad8",
		// 							Name: "Basing",
		// 							Cost: 1,
		// 						},
		// 						{
		// 							ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad9",
		// 							Name: "Assembly",
		// 							Cost: 10,
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommissionTotal(tt.args.commission); got != tt.want {
				t.Errorf("CommissionTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
