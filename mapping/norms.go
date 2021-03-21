package mapping

// WithNorms is a mapping with the Norms parameter
//
// Norms store various normalization factors that are later used at query time
// in order to compute the score of a document relatively to a query.
//
// Although useful for scoring, norms also require quite a lot of disk
// (typically in the order of one byte per document per field in your index,
// even for documents that don’t have this specific field). As a consequence, if
// you don’t need scoring on a specific field, you should disable norms on that
// field. In particular, this is the case for fields that are used solely for
// filtering or aggregations.
//
// Norms can be disabled (but not reenabled after the fact)
//
// If updating the norms via  the REST API, they will not be removed instantly,
// but will be removed as old segments are merged into new segments as you
// continue indexing new documents. Any score computation on a field that has
// had norms removed might return inconsistent results since some documents
// won’t have norms anymore while other documents might still have norms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/norms.html
type WithNorms interface {
	// Whether field-length should be taken into account when scoring queries.
	// Accepts true (default) or false.
	Norms() bool
	// SetNorms sets the Norms value to v
	SetIndexPhrases(v bool)
}

// FieldWithNorms is a Field with the norms parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/norms.html
type FieldWithNorms interface {
	Field
	WithNorms
}

// NormsParam is a mixin that adds the norms parameter
//
// Norms store various normalization factors that are later used at query time
// in order to compute the score of a document relatively to a query.
//
// Although useful for scoring, norms also require quite a lot of disk
// (typically in the order of one byte per document per field in your index,
// even for documents that don’t have this specific field). As a consequence, if
// you don’t need scoring on a specific field, you should disable norms on that
// field. In particular, this is the case for fields that are used solely for
// filtering or aggregations.
//
// Norms can be disabled (but not reenabled after the fact)
//
// If updating the norms via  the REST API, they will not be removed instantly,
// but will be removed as old segments are merged into new segments as you
// continue indexing new documents. Any score computation on a field that has
// had norms removed might return inconsistent results since some documents
// won’t have norms anymore while other documents might still have norms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/norms.html
type NormsParam struct {
	NormsValue *bool `bson:"norms,omitempty" json:"norms,omitempty"`
}

// Norms determines whether field-length should be taken into account when
// scoring queries. Accepts true (default) or false.
func (n NormsParam) Norms() bool {
	if n.NormsValue != nil {
		return *n.NormsValue
	}
	return true
}

// SetNorms sets the Norms value to v
func (n *NormsParam) SetNorms(v bool) {
	if n.Norms() != v {
		n.NormsValue = &v
	}
}
