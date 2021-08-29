package abstractfactory

import (
	"reflect"
	"testing"
)

func TestJsonConfigParserFactory_CreateRuleParser(t *testing.T) {
	tests := []struct {
		name string
		j    JsonConfigParserFactory
		want IRuleConfigParser
	}{
		{
			name: "JsonCreateRuleParser",
			j:    JsonConfigParserFactory{},
			want: &JsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonConfigParserFactory.CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonConfigParserFactory_CreateSystemParser(t *testing.T) {
	tests := []struct {
		name string
		j    JsonConfigParserFactory
		want ISystemConfigParser
	}{
		{
			name: "JsonCreateSystemParser",
			j:    JsonConfigParserFactory{},
			want: &JsonSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonConfigParserFactory{}
			if got := j.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonConfigParserFactory.CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlConfigParserFactory_CreateRuleParser(t *testing.T) {
	tests := []struct {
		name string
		j    XmlConfigParserFactory
		want IRuleConfigParser
	}{
		{
			name: "XmlCreateRuleParser",
			j:    XmlConfigParserFactory{},
			want: &XmlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := XmlConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlConfigParserFactory.CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlConfigParserFactory_CreateSystemParser(t *testing.T) {
	tests := []struct {
		name string
		j    XmlConfigParserFactory
		want ISystemConfigParser
	}{
		{
			name: "XmlCreateSystemParser",
			j:    XmlConfigParserFactory{},
			want: &XmlSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := XmlConfigParserFactory{}
			if got := j.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlConfigParserFactory.CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
