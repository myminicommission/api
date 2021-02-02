package calc

import (
	"testing"
	"time"

	"github.com/myminicommission/api/graph/model"
)

func TestEstimateTotal(t *testing.T) {
	type args struct {
		estimate *model.Estimate
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "estimate with options",
			want: 111,
			args: args{
				estimate: &model.Estimate{
					ID:        "70bfc0e1-7927-49da-83e8-be8157b4a00e",
					CreatedAt: time.Now(),
					Minis: []*model.MiniQuantity{
						{
							ID:       "f96f9fe1-c6fa-463a-9f24-2a86666642c5",
							Quantity: 1,
							Mini: &model.Mini{
								ID:   "c84b3f7a-b402-4bb4-a5eb-a08194e98aa3",
								Name: "Silent King",
								Size: model.MiniSizeLarge,
								Cost: 100,
							},
							Options: []*model.MiniOption{
								{
									ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad8",
									Name: "Basing",
									Cost: 1,
								},
								{
									ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad9",
									Name: "Assembly",
									Cost: 10,
								},
							},
						},
					},
					User: &model.User{
						Name:      "Extreme Moderation",
						ID:        "25d816e3-31bd-4cdf-86ab-e2bdb406b907",
						CreatedAt: time.Now().Add(-5 * time.Hour),
					},
				},
			},
		},
		{
			name: "estimate with no minis",
			want: 0,
			args: args{
				estimate: &model.Estimate{
					CreatedAt: time.Now(),
					Minis:     []*model.MiniQuantity{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EstimateTotal(tt.args.estimate); got != tt.want {
				t.Errorf("EstimateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuoteTotal(t *testing.T) {
	type args struct {
		quote *model.Quote
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "quote with no minis",
			want: 0,
			args: args{
				quote: &model.Quote{
					CreatedAt: time.Now(),
					Minis:     []*model.MiniQuantity{},
				},
			},
		},
		{
			name: "quote with options",
			want: 111,
			args: args{
				quote: &model.Quote{
					ID:        "70bfc0e1-7927-49da-83e8-be8157b4a00e",
					CreatedAt: time.Now(),
					Minis: []*model.MiniQuantity{
						{
							ID:       "f96f9fe1-c6fa-463a-9f24-2a86666642c5",
							Quantity: 1,
							Mini: &model.Mini{
								ID:   "c84b3f7a-b402-4bb4-a5eb-a08194e98aa3",
								Name: "Silent King",
								Size: model.MiniSizeLarge,
								Cost: 100,
							},
							Options: []*model.MiniOption{
								{
									ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad8",
									Name: "Basing",
									Cost: 1,
								},
								{
									ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad9",
									Name: "Assembly",
									Cost: 10,
								},
							},
						},
					},
					User: &model.User{
						Name:      "Extreme Moderation",
						ID:        "25d816e3-31bd-4cdf-86ab-e2bdb406b907",
						CreatedAt: time.Now().Add(-5 * time.Hour),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuoteTotal(tt.args.quote); got != tt.want {
				t.Errorf("QuoteTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcMiniTotals(t *testing.T) {
	type args struct {
		minis []*model.MiniQuantity
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test without options",
			want: 100,
			args: args{
				minis: []*model.MiniQuantity{
					{
						ID:       "f96f9fe1-c6fa-463a-9f24-2a86666642c5",
						Quantity: 1,
						Mini: &model.Mini{
							ID:   "c84b3f7a-b402-4bb4-a5eb-a08194e98aa3",
							Name: "Silent King",
							Size: model.MiniSizeLarge,
							Cost: 100,
						},
					},
				},
			},
		},
		{
			name: "test with options",
			want: 111,
			args: args{
				minis: []*model.MiniQuantity{
					{
						ID:       "f96f9fe1-c6fa-463a-9f24-2a86666642c5",
						Quantity: 1,
						Mini: &model.Mini{
							ID:   "c84b3f7a-b402-4bb4-a5eb-a08194e98aa3",
							Name: "Silent King",
							Size: model.MiniSizeLarge,
							Cost: 100,
						},
						Options: []*model.MiniOption{
							{
								ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad8",
								Name: "Basing",
								Cost: 1,
							},
							{
								ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad9",
								Name: "Assembly",
								Cost: 10,
							},
						},
					},
				},
			},
		},
		{
			name: "test with options and many minis",
			want: 177,
			args: args{
				minis: []*model.MiniQuantity{
					{
						Quantity: 1,
						Mini: &model.Mini{
							Name: "Silent King",
							Size: model.MiniSizeLarge,
							Cost: 100,
						},
						Options: []*model.MiniOption{
							{
								Name: "Basing",
								Cost: 1,
							},
							{
								Name: "Assembly",
								Cost: 10,
							},
						},
					},
					{
						Quantity: 1,
						Mini: &model.Mini{
							Name: "Drahzar",
							Size: model.MiniSizeRegular,
							Cost: 9,
						},
						Options: []*model.MiniOption{
							{
								Name: "Basing",
								Cost: 1,
							},
							{
								Name: "Assembly",
								Cost: 1,
							},
						},
					},
					{
						Quantity: 5,
						Mini: &model.Mini{
							Name: "Incubi",
							Size: model.MiniSizeRegular,
							Cost: 9,
						},
						Options: []*model.MiniOption{
							{
								Name: "Basing",
								Cost: 1,
							},
							{
								Name: "Assembly",
								Cost: 1,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMiniTotals(tt.args.minis); got != tt.want {
				t.Errorf("calcMiniTotals() = %v, want %v", got, tt.want)
			}
		})
	}
}
