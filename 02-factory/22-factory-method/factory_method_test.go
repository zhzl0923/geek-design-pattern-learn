package factorymethod

import (
	"reflect"
	"testing"
)

func TestJsonRuleConfigParserFactory_CreateParser(t *testing.T) {
	tests := []struct {
		name string
		j    JsonRuleConfigParserFactory
		want IRuleConfigParser
	}{
		{
			name: "JsonCreateParser",
			j:    JsonRuleConfigParserFactory{},
			want: &JsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonRuleConfigParserFactory{}
			if got := j.CreateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonRuleConfigParserFactory.CreateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlRuleConfigParserFactory_CreateParser(t *testing.T) {
	tests := []struct {
		name string
		j    XmlRuleConfigParserFactory
		want IRuleConfigParser
	}{
		{
			name: "XmlCreateParser",
			j:    XmlRuleConfigParserFactory{},
			want: &XmlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := XmlRuleConfigParserFactory{}
			if got := j.CreateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XmlRuleConfigParserFactory.CreateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYamlRuleConfigParserFactory_CreateParser(t *testing.T) {
	tests := []struct {
		name string
		j    YamlRuleConfigParserFactory
		want IRuleConfigParser
	}{
		{
			name: "YamlCreateParser",
			j:    YamlRuleConfigParserFactory{},
			want: &YamlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := YamlRuleConfigParserFactory{}
			if got := j.CreateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YamlRuleConfigParserFactory.CreateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertiesRuleConfigParserFactory_CreateParser(t *testing.T) {
	tests := []struct {
		name string
		j    PropertiesRuleConfigParserFactory
		want IRuleConfigParser
	}{
		{
			name: "PropertiesCreateParser",
			j:    PropertiesRuleConfigParserFactory{},
			want: &PropertiesRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := PropertiesRuleConfigParserFactory{}
			if got := j.CreateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PropertiesRuleConfigParserFactory.CreateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
