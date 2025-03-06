package slimit_test

import (
	"testing"

	"github.com/thinkgos/TDengine-gorm/clause/slimit"
	"github.com/thinkgos/TDengine-gorm/clause/tests"

	"gorm.io/gorm/clause"
)

func Test_SLimit(t *testing.T) {
	testCases := []struct {
		Name    string
		Clauses []clause.Interface
		Result  string
		Vars    []any
	}{
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SLimit{
				Limit:  10,
				Offset: 20,
			}},
			"SELECT * FROM `users` SLIMIT 10 SOFFSET 20", nil,
		},
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SLimit{Offset: 20}},
			"SELECT * FROM `users` SOFFSET 20", nil,
		},
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SLimit{Offset: 20}, slimit.SLimit{Offset: 30}},
			"SELECT * FROM `users` SOFFSET 30", nil,
		},
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SLimit{Offset: 20}, slimit.SLimit{Limit: 10}},
			"SELECT * FROM `users` SLIMIT 10 SOFFSET 20", nil,
		},
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SLimit{Limit: 10, Offset: 20}, slimit.SLimit{Offset: 30}},
			"SELECT * FROM `users` SLIMIT 10 SOFFSET 30", nil,
		},
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SLimit{Limit: 10, Offset: 20}, slimit.SLimit{Offset: 30}, slimit.SLimit{Offset: -10}},
			"SELECT * FROM `users` SLIMIT 10", nil,
		},
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SLimit{Limit: 10, Offset: 20}, slimit.SLimit{Offset: 30}, slimit.SLimit{Limit: -10}},
			"SELECT * FROM `users` SOFFSET 30", nil,
		},
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SLimit{Limit: 10, Offset: 20}, slimit.SLimit{Offset: 30}, slimit.SLimit{Limit: 50}},
			"SELECT * FROM `users` SLIMIT 50 SOFFSET 30", nil,
		},
		{
			"",
			[]clause.Interface{clause.Select{}, clause.From{}, slimit.SetSLimit(10, 20), slimit.SetSLimit(0, 30), slimit.SetSLimit(50, 0)},
			"SELECT * FROM `users` SLIMIT 50 SOFFSET 30", nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tests.CheckBuildClauses(t, tc.Clauses, []string{tc.Result}, [][][]any{{tc.Vars}})
		})
	}
}
