package factorymethod

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

type RuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type JsonRuleConfigParserFactory struct{}

func (j JsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &JsonRuleConfigParser{}
}

type XmlRuleConfigParserFactory struct{}

func (j XmlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &XmlRuleConfigParser{}
}

type YamlRuleConfigParserFactory struct{}

func (j YamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &YamlRuleConfigParser{}
}

type PropertiesRuleConfigParserFactory struct{}

func (j PropertiesRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &PropertiesRuleConfigParser{}
}
