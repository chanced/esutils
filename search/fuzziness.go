package search

import "github.com/tidwall/gjson"

const DefaultFuzziness = "0"

// WithFuzziness is an interface for queries with fuzziness the parameter
//
// Maximum edit distance allowed for matching. See Fuzziness for valid values and more information.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
type WithFuzziness interface {
	// Fuzziness is the maximum edit distance allowed for matching. See
	// Fuzziness for valid values and more information. See Fuzziness in the
	// match query for an example.
	Fuzziness() string
	// SetFuzziness sets the FuzzinessValue to v
	SetFuzziness(str string)
	// FuzzyRewrite is the method used to rewrite the query. See the rewrite
	// parameter for valid values and more information.
	//
	// If the fuzziness parameter is not 0, the match query uses a fuzzy_rewrite
	// method of top_terms_blended_freqs_${max_expansions} by default.
	FuzzyRewrite() Rewrite
	// SetFuzzyRewrite sets the value of FuzzyRewrite to v
	SetFuzzyRewrite(v Rewrite)
	DefaultFuzzyRewrite() Rewrite
}

// FuzzinessParam is a mixin that adds the fuzziness parameter to queries
//
// Maximum edit distance allowed for matching. See Fuzziness for valid values and more information.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
type FuzzinessParam struct {
	FuzzinessValue    string  `json:"fuzziness,omitempty" bson:"fuzziness,omitempty"`
	FuzzyRewriteValue Rewrite `json:"fuzzy_rewrite,omitempty" bson:"fuzzy_rewrite,omitempty"`
}

// Fuzziness is the maximum edit distance allowed for matching. See Fuzziness
// for valid values and more information.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
func (f FuzzinessParam) Fuzziness() string {
	return f.FuzzinessValue
}

// SetFuzziness sets the fuzzinessValue to v
func (f *FuzzinessParam) SetFuzziness(v string) {
	if f.Fuzziness() != v {
		f.FuzzinessValue = v
	}
}

// FuzzyRewrite is the method used to rewrite the query. See the rewrite
// parameter for valid values and more information.
//
// If the fuzziness parameter is not 0, the match query uses a fuzzy_rewrite
// method of top_terms_blended_freqs_${max_expansions} by default.
func (f FuzzinessParam) FuzzyRewrite() Rewrite {
	if f.FuzzyRewriteValue != "" {
		return f.FuzzyRewriteValue
	}
	return f.DefaultFuzzyRewrite()
}

func (f FuzzinessParam) DefaultFuzzyRewrite() Rewrite {
	if f.Fuzziness() != DefaultFuzziness {
		return RewriteTopTermsBlendedFreqsN
	}
	return RewriteConstantScore
}

// SetFuzzyRewrite sets the value of FuzzyRewrite to v
func (f *FuzzinessParam) SetFuzzyRewrite(v Rewrite) error {
	if !v.IsValid() {
		return ErrInvalidRewrite
	}
	if f.FuzzyRewrite() != v {
		f.FuzzyRewriteValue = v
	}
	return nil
}

func unmarshalFuzzinessParam(data gjson.Result, target interface{}) error {
	if r, ok := target.(WithFuzziness); ok {
		r.SetFuzziness(data.Str)
	}
	return nil
}

func marshalFuzzinessParam(data map[string]interface{}, source interface{}) (map[string]interface{}, error) {
	if a, ok := source.(WithFuzziness); ok {
		if a.Fuzziness() != DefaultFuzziness {
			data[paramFuzziness] = a.Fuzziness()
		}
	}
	return data, nil
}
func unmarshalFuzzyRewriteParam(data gjson.Result, target interface{}) error {
	if r, ok := target.(WithFuzziness); ok {
		r.SetFuzzyRewrite(Rewrite(data.Str))
	}
	return nil
}
func marshalFuzzyRewriteParam(data map[string]interface{}, source interface{}) (map[string]interface{}, error) {
	if a, ok := source.(WithFuzziness); ok {
		if a.FuzzyRewrite() != a.DefaultFuzzyRewrite() {
			data[paramFuzzyRewrite] = a.FuzzyRewrite()
		}
	}
	return data, nil
}
