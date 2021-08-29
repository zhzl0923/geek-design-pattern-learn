package builder

import (
	"reflect"
	"testing"
)

func TestBuilder_Build(t *testing.T) {
	type fields struct {
		name     string
		maxTotal int
		maxIdle  int
		minIdle  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *ResourcePoolConfig
	}{
		{
			name: "test Builder builder name empty",
			fields: fields{
				name:     "",
				maxTotal: 8,
			},
		},
		{
			name: "test Builder builder name not empty",
			fields: fields{
				name:     "test",
				maxTotal: 8,
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: defaultMaxTotal,
				maxIdle:  defaultMaxIdle,
				minIdle:  defaultMinIdle,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if e := recover(); e != nil {
					t.Errorf("Panicing %s\r\n", e) // Panicing bad end
				}
			}()
			b := &Builder{}
			b = b.SetName(tt.fields.name).SetMaxIdle(tt.fields.maxIdle).SetMaxTotal(tt.fields.maxTotal).SetMinIdle(tt.fields.minIdle)
			if got := b.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
