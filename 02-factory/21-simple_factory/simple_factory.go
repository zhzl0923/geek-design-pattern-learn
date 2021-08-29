package simplefactory

type IRuleConfigParser interface {
	Parse()
}

type JsonRuleConfigParser struct{}

func (p JsonRuleConfigParser) Parse() {}

type XmlRuleConfigParser struct{}

func (p XmlRuleConfigParser) Parse() {}

type YamlRuleConfigParser struct{}

func (p YamlRuleConfigParser) Parse() {}

type PropertiesRuleConfigParser struct{}

func (p PropertiesRuleConfigParser) Parse() {}

type RuleConfigParserFactory struct{}

func (f RuleConfigParserFactory) CreateParser(configFormat string) IRuleConfigParser {
	var parser IRuleConfigParser
	if configFormat == "json" {
		parser = &JsonRuleConfigParser{}
	} else if configFormat == "xml" {
		parser = &XmlRuleConfigParser{}
	} else if configFormat == "yaml" {
		parser = &YamlRuleConfigParser{}
	} else if configFormat == "properties" {
		parser = &PropertiesRuleConfigParser{}
	}
	return parser
}
