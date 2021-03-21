package search

type MatchPhrase struct {
	Query               string `json:"query" bson:"query"`
	AnalyzerParam       `json:",inline" bson:",inline"`
	ZeroTermsQueryParam `json:",inline" bson:",inline"`
}

type MatchPhraseQuery struct {
	MatchPhrase map[string]MatchPhrase `json:"match_phrase,omitempty" bson:"match_phrase,omitempty"`
}

func (mp *MatchPhraseQuery) SetMatchPhrase(field string, matchPhrase MatchPhrase) {
	mp.MatchPhrase[field] = matchPhrase
}
