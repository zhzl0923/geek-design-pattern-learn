package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAlertRule_Interpret(t *testing.T) {
	stats := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	tests := []struct {
		name  string
		stats map[string]int
		rule  string
		want  bool
	}{
		{
			name:  "case1",
			stats: stats,
			rule:  "a > 1 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case2",
			stats: stats,
			rule:  "a < 2 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case3",
			stats: stats,
			rule:  "a < 5 && b > 1 && c < 10",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewAlertRuleInterpreter(tt.rule)
			require.NoError(t, err)
			assert.Equal(t, tt.want, r.Interpret(tt.stats))
		})
	}
}
