package movie

import "testing"

func TestParam_Valid(t *testing.T) {
	type fields struct {
		Keyword string
		Page    int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"should be valid",
			fields{
				Keyword: "batman",
				Page:    1,
			},
			true,
		},
		{
			"should be invalid",
			fields{
				Keyword: "batman",
				Page:    0,
			},
			false,
		},
		{
			"should be invalid",
			fields{
				Keyword: "",
				Page:    1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Param{
				Keyword: tt.fields.Keyword,
				Page:    tt.fields.Page,
			}
			if got := p.Valid(); got != tt.want {
				t.Errorf("Param.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
