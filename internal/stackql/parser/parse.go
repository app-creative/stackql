package parser

import (
	"fmt"

	"github.com/stackql/stackql-parser/go/vt/sqlparser"
)

func specialiseParserError(err error, cmd string) error {
	if err != nil {
		return fmt.Errorf(
			`error:  You have an error in your stackql syntax; parser error: %s`,
			err.Error(),
		)
	}
	return err
}

type Parser interface {
	ParseQuery(cmd string) (sqlparser.Statement, error)
}

func NewParser() (Parser, error) {
	return &basicParser{}, nil
}

type basicParser struct{}

func (p *basicParser) ParseQuery(cmd string) (sqlparser.Statement, error) {
	statement, err := sqlparser.Parse(cmd)
	return statement, specialiseParserError(err, cmd)
}