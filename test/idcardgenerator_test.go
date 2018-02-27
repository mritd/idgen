package test

import (
	"fmt"
	"github.com/mritd/idgen/generator"
	"testing"
)

func Test_GenerateIssueOrg(t *testing.T) {
	fmt.Println(generator.GenerateIssueOrg())
}

func Test_GenerateValidPeriod(t *testing.T) {
	fmt.Println(generator.GenerateValidPeriod())
}
