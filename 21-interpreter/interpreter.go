package interpreter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//接口告警规则
type AlertRuleInterpreter struct {
	expression Expression
}

// NewAlertRule NewAlertRule
func NewAlertRuleInterpreter(rule string) (*AlertRuleInterpreter, error) {
	exp, err := NewAndExpression(rule)
	return &AlertRuleInterpreter{expression: exp}, err
}

func (a *AlertRuleInterpreter) Interpret(stats map[string]int) bool {
	return a.expression.Interpret(stats)
}

type Expression interface {
	Interpret(stats map[string]int) bool
}

type GreaterExpression struct {
	key   string
	value int
}

func NewGreaterExpression(exp string) (*GreaterExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != ">" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}
	v, err := strconv.Atoi(data[2])
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}
	return &GreaterExpression{
		key:   data[0],
		value: v,
	}, nil
}

func (g GreaterExpression) Interpret(stats map[string]int) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v > g.value
}

type LessExpression struct {
	key   string
	value int
}

func NewLessExpression(exp string) (*LessExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != "<" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}
	v, err := strconv.Atoi(data[2])
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}
	return &LessExpression{
		key:   data[0],
		value: v,
	}, nil
}

func (l LessExpression) Interpret(stats map[string]int) bool {
	v, ok := stats[l.key]
	if !ok {
		return false
	}
	return v < l.value
}

type AndExpression struct {
	expressions []Expression
}

func NewAndExpression(exp string) (*AndExpression, error) {
	exps := strings.Split(exp, "&&")
	expressions := make([]Expression, len(exps))
	for k, e := range exps {
		var expression Expression
		var err error
		switch {
		case strings.Contains(e, ">"):
			expression, err = NewGreaterExpression(e)
		case strings.Contains(e, "<"):
			expression, err = NewLessExpression(e)
		default:
			err = fmt.Errorf("exp is invalid: %s", exp)
		}
		if err != nil {
			return nil, err
		}

		expressions[k] = expression
	}
	return &AndExpression{expressions: expressions}, nil
}

func (a *AndExpression) Interpret(stats map[string]int) bool {
	for _, expression := range a.expressions {
		if !expression.Interpret(stats) {
			return false
		}
	}
	return true
}
