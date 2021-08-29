package abstractfactory

type IRuleConfigParser interface {
	ParseRuleConfig()
}

type JsonRuleConfigParser struct {
}

func (j JsonRuleConfigParser) ParseRuleConfig() {}

type XmlRuleConfigParser struct{}

func (x XmlRuleConfigParser) ParseRuleConfig() {}

type ISystemConfigParser interface {
	ParseSystemConfig()
}

type JsonSystemConfigParser struct{}

func (j JsonSystemConfigParser) ParseSystemConfig() {}

type XmlSystemConfigParser struct{}

func (x XmlSystemConfigParser) ParseSystemConfig() {}

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type JsonConfigParserFactory struct{}

func (j JsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return &JsonRuleConfigParser{}
}

func (j JsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return &JsonSystemConfigParser{}
}

type XmlConfigParserFactory struct{}

func (j XmlConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return &XmlRuleConfigParser{}
}

func (j XmlConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return &XmlSystemConfigParser{}
}
