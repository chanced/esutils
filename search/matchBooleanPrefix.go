package search

type MatchBooleanPrefix struct {
	Query                    string `json:"query" bson:"query"`
	MinimumShouldMatchParam  `json:",inline" bson:",inline"`
	OperatorParam            `json:",inline" bson:",inline"`
	AnalyzerParam            `json:",inline" bson:",inline"`
	FuzzinessParam           `json:",inline" bson:",inline"`
	PrefixLengthParam        `json:",inline" bson:",inline"`
	FuzzyTranspositionsParam `json:",inline" bson:",inline"`
}

type MatchBooleanPrefixQuery struct {
	MatchBooleanPrefix map[string]MatchBooleanPrefix `json:"match_boolean_prefix,omitempty" bson:"match_boolean_prefix,omitempty"`
}

func NewMatchBooleanPrefixQuery() MatchBooleanPrefixQuery {
	return MatchBooleanPrefixQuery{
		MatchBooleanPrefix: map[string]MatchBooleanPrefix{},
	}
}
