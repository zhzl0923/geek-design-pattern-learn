package simplefactory

import (
	"reflect"
	"testing"
)

func TestCreateParser(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &JsonRuleConfigParser{},
		},
		{
			name: "xml",
			args: args{t: "xml"},
			want: &XmlRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &YamlRuleConfigParser{},
		},
		{
			name: "properties",
			args: args{t: "properties"},
			want: &PropertiesRuleConfigParser{},
		},
	}
	factory := RuleConfigParserFactory{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory.CreateParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factory.CreateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
