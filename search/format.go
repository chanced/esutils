package search

import "github.com/tidwall/gjson"

// WithFormat is a query with a format parameter
//
// In JSON documents, dates are represented as strings. Elasticsearch uses a set
// of preconfigured formats to recognize and parse these strings into a long
// value representing milliseconds-since-the-epoch in UTC.
//
// Besides the built-in formats, custom formats can be specified using the
// familiar yyyy/MM/dd syntax
//
//  "epoch_millis"
// A formatter for the number of milliseconds since the epoch. Note, that this
// timestamp is subject to the limits of a Java Long.MIN_VALUE and
// Long.MAX_VALUE.
//
//  "epoch_second"
// A formatter for the number of seconds since the epoch. Note, that this
// timestamp is subject to the limits of a Java Long.MIN_VALUE and Long.
// MAX_VALUE divided by 1000 (the number of milliseconds in a second).
//
//  "date_optional_time" || "strict_date_optional_time"
// A generic ISO datetime parser, where the date must include the year at a
// minimum, and the time (separated by T), is optional. Examples:
// "yyyy-MM-dd'T'HH:mm:ss.SSSZ" || "yyyy-MM-dd"
//
//  "strict_date_optional_time_nanos"
// A generic ISO datetime parser, where the date must include the year at a
// minimum, and the time (separated by T), is optional. The fraction of a second
// part has a nanosecond resolution. Examples: yyyy-MM-dd'T'HH:mm:ss.SSSSSSZ or
// yyyy-MM-dd.
//
//  "basic_date"
// A basic formatter for a full date as four digit year, two digit month of
// year, and two digit day of month: yyyyMMdd.
//
//  "basic_date_time"
// A basic formatter that combines a basic date and time, separated by a T:
// yyyyMMdd'T'HHmmss.SSSZ.
//
//  "basic_date_time_no_millis"
// A basic formatter that combines a basic date and time without millis,
// separated by a T: yyyyMMdd'T'HHmmssZ.
//
//  "basic_ordinal_date"
// A formatter for a full ordinal date, using a four digit year and three digit
// dayOfYear: yyyyDDD.
//
//  "basic_ordinal_date_time"
// A formatter for a full ordinal date and time, using a four digit year and
// three digit dayOfYear: yyyyDDD'T'HHmmss.SSSZ.
//
//  "basic_ordinal_date_time_no_millis"
// A formatter for a full ordinal date and time without millis, using a four
// digit year and three digit dayOfYear: yyyyDDD'T'HHmmssZ.
//
//  "basic_time"
// A basic formatter for a two digit hour of day, two digit minute of hour, two
// digit second of minute, three digit millis, and time zone offset:
// HHmmss.SSSZ.
//
//  "basic_time_no_millis"
// A basic formatter for a two digit hour of day, two digit minute of hour, two
// digit second of minute, and time zone offset: HHmmssZ.
//
//  "basic_t_time"
// A basic formatter for a two digit hour of day, two digit minute of hour, two
// digit second of minute, three digit millis, and time zone off set prefixed by
// T: 'T'HHmmss.SSSZ.
//
//  "basic_t_time_no_millis"
// A basic formatter for a two digit hour of day, two digit minute of hour, two
// digit second of minute, and time zone offset prefixed by T: 'T'HHmmssZ.
//
//  "basic_week_date" || "strict_basic_week_date"
// A basic formatter for a full date as four digit weekyear, two digit week of
// weekyear, and one digit day of week: xxxx'W'wwe.
//
//  "basic_week_date_time" || "strict_basic_week_date_time"
// A basic formatter that combines a basic weekyear date and time, separated by
// a T: xxxx'W'wwe'T'HHmmss.SSSZ.
//
//  "basic_week_date_time_no_millis" || "strict_basic_week_date_time_no_millis"
// A basic formatter that combines a basic weekyear date and time without
// millis, separated by a T: xxxx'W'wwe'T'HHmmssZ.
//
//  "date" || "strict_date"
// A formatter for a full date as four digit year, two digit month of year, and
// two digit day of month: yyyy-MM-dd.
//
//  "date_hour" || "strict_date_hour"
// A formatter that combines a full date and two digit hour of day:
// yyyy-MM-dd'T'HH.
//
//  "date_hour_minute" || "strict_date_hour_minute"
// A formatter that combines a full date, two digit hour of day, and two digit
// minute of hour: yyyy-MM-dd'T'HH:mm.
//
//  "date_hour_minute_second" || "strict_date_hour_minute_second"
// A formatter that combines a full date, two digit hour of day, two digit
// minute of hour, and two digit second of minute: yyyy-MM-dd'T'HH:mm:ss.
//
//  "date_hour_minute_second_fraction" || "strict_date_hour_minute_second_fraction"
// A formatter that combines a full date, two digit hour of day, two digit
// minute of hour, two digit second of minute, and three digit fraction of
// second: yyyy-MM-dd'T'HH:mm:ss.SSS.
//
// "date_hour_minute_second_millis" || "strict_date_hour_minute_second_millis" A
// formatter that combines a full date, two digit hour of day, two digit minute
// of hour, two digit second of minute, and three digit fraction of second:
// yyyy-MM-dd'T'HH:mm:ss.SSS.
//
//  "date_time" || "strict_date_time"
// A formatter that combines a full date and time, separated by a T:
// yyyy-MM-dd'T'HH:mm:ss.SSSZZ.
//
// "date_time_no_millis" || "strict_date_time_no_millis" A formatter that
// combines a full date and time without millis, separated by a T:
// yyyy-MM-dd'T'HH:mm:ssZZ.
//
// "hour" || "strict_hour" A formatter for a two digit hour of day: HH
//
//  "hour_minute" || "strict_hour_minute"
// A formatter for a two digit hour of day and two digit minute of hour: HH:mm.
//
//  "hour_minute_second" || "strict_hour_minute_second"
// A formatter for a two digit hour of day, two digit minute of hour, and two
// digit second of minute: HH:mm:ss.
//
//  "hour_minute_second_fraction" || "strict_hour_minute_second_fraction"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, and three digit fraction of second: HH:mm:ss.SSS.
//
//  "hour_minute_second_millis" || "strict_hour_minute_second_millis"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, and three digit fraction of second: HH:mm:ss.SSS.
//
//  "ordinal_date" || "strict_ordinal_date"
// A formatter for a full ordinal date, using a four digit year and three digit
// dayOfYear: yyyy-DDD.
//
//  "ordinal_date_time" || "strict_ordinal_date_time"
// A formatter for a full ordinal date and time, using a four digit year and
// three digit dayOfYear: yyyy-DDD'T'HH:mm:ss.SSSZZ.
//
//  "ordinal_date_time_no_millis" || "strict_ordinal_date_time_no_millis"
// A formatter for a full ordinal date and time without millis, using a four
// digit year and three digit dayOfYear: yyyy-DDD'T'HH:mm:ssZZ.
//
//  "time" || "strict_time"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, three digit fraction of second, and time zone offset:
// HH:mm:ss.SSSZZ.
//
//  "time_no_millis" || "strict_time_no_millis"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, and time zone offset: HH:mm:ssZZ.
//
//  "t_time" || "strict_t_time"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, three digit fraction of second, and time zone offset
// prefixed by T: 'T'HH:mm:ss.SSSZZ.
//
//  "t_time_no_millis" || "strict_t_time_no_millis"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, and time zone offset prefixed by T: 'T'HH:mm:ssZZ.
//
//  "week_date" || "strict_week_date"
// A formatter for a full date as four digit weekyear, two digit week of
// weekyear, and one digit day of week: xxxx-'W'ww-e.
//
//  "week_date_time" || "strict_week_date_time"
// A formatter that combines a full weekyear date and time, separated by a T:
// xxxx-'W'ww-e'T'HH:mm:ss.SSSZZ.
//
//  "week_date_time_no_millis" || "strict_week_date_time_no_millis"
// A formatter that combines a full weekyear date and time without millis,
// separated by a T: xxxx-'W'ww-e'T'HH:mm:ssZZ.
//
//  "weekyear" || "strict_weekyear"
// A formatter for a four digit weekyear: xxxx.
//
//  "weekyear_week" || "strict_weekyear_week"
// A formatter for a four digit weekyear and two digit week of weekyear:
// xxxx-'W'ww.
//
//  "weekyear_week_day" || "strict_weekyear_week_day"
// A formatter for a four digit weekyear, two digit week of weekyear, and one
// digit day of week: xxxx-'W'ww-e.
//
//  "year" || "strict_year"
// A formatter for a four digit year: yyyy.
//
//  "year_month" || "strict_year_month"
// A formatter for a four digit year and two digit month of year: yyyy-MM.
//
//  "year_month_day" || "strict_year_month_day"
// A formatter for a four digit year, two digit month of year, and two digit day
// of month: yyyy-MM-dd.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-date-format.html
type WithFormat interface {
	//The date format(s) that can be parsed. Defaults to
	//strict_date_optional_time||epoch_millis.
	Format() string
	// SetFormat sets the Format Value to v
	SetFormat(v string)
}

// FormatParam is a mixin that adds the format parameter to mappings
//
// In JSON documents, dates are represented as strings. Elasticsearch uses a set
// of preconfigured formats to recognize and parse these strings into a long
// value representing milliseconds-since-the-epoch in UTC.
//
// Besides the built-in formats, custom formats can be specified using the
// familiar yyyy/MM/dd syntax
//
//  "epoch_millis"
// A formatter for the number of milliseconds since the epoch. Note, that this
// timestamp is subject to the limits of a Java Long.MIN_VALUE and
// Long.MAX_VALUE.
//
//  "epoch_second"
// A formatter for the number of seconds since the epoch. Note, that this
// timestamp is subject to the limits of a Java Long.MIN_VALUE and Long.
// MAX_VALUE divided by 1000 (the number of milliseconds in a second).
//
//  "date_optional_time" || "strict_date_optional_time"
// A generic ISO datetime parser, where the date must include the year at a
// minimum, and the time (separated by T), is optional. Examples:
// "yyyy-MM-dd'T'HH:mm:ss.SSSZ" || "yyyy-MM-dd"
//
//  "strict_date_optional_time_nanos"
// A generic ISO datetime parser, where the date must include the year at a
// minimum, and the time (separated by T), is optional. The fraction of a second
// part has a nanosecond resolution. Examples: yyyy-MM-dd'T'HH:mm:ss.SSSSSSZ or
// yyyy-MM-dd.
//
//  "basic_date"
// A basic formatter for a full date as four digit year, two digit month of
// year, and two digit day of month: yyyyMMdd.
//
//  "basic_date_time"
// A basic formatter that combines a basic date and time, separated by a T:
// yyyyMMdd'T'HHmmss.SSSZ.
//
//  "basic_date_time_no_millis"
// A basic formatter that combines a basic date and time without millis,
// separated by a T: yyyyMMdd'T'HHmmssZ.
//
//  "basic_ordinal_date"
// A formatter for a full ordinal date, using a four digit year and three digit
// dayOfYear: yyyyDDD.
//
//  "basic_ordinal_date_time"
// A formatter for a full ordinal date and time, using a four digit year and
// three digit dayOfYear: yyyyDDD'T'HHmmss.SSSZ.
//
//  "basic_ordinal_date_time_no_millis"
// A formatter for a full ordinal date and time without millis, using a four
// digit year and three digit dayOfYear: yyyyDDD'T'HHmmssZ.
//
//  "basic_time"
// A basic formatter for a two digit hour of day, two digit minute of hour, two
// digit second of minute, three digit millis, and time zone offset:
// HHmmss.SSSZ.
//
//  "basic_time_no_millis"
// A basic formatter for a two digit hour of day, two digit minute of hour, two
// digit second of minute, and time zone offset: HHmmssZ.
//
//  "basic_t_time"
// A basic formatter for a two digit hour of day, two digit minute of hour, two
// digit second of minute, three digit millis, and time zone off set prefixed by
// T: 'T'HHmmss.SSSZ.
//
//  "basic_t_time_no_millis"
// A basic formatter for a two digit hour of day, two digit minute of hour, two
// digit second of minute, and time zone offset prefixed by T: 'T'HHmmssZ.
//
//  "basic_week_date" || "strict_basic_week_date"
// A basic formatter for a full date as four digit weekyear, two digit week of
// weekyear, and one digit day of week: xxxx'W'wwe.
//
//  "basic_week_date_time" || "strict_basic_week_date_time"
// A basic formatter that combines a basic weekyear date and time, separated by
// a T: xxxx'W'wwe'T'HHmmss.SSSZ.
//
//  "basic_week_date_time_no_millis" || "strict_basic_week_date_time_no_millis"
// A basic formatter that combines a basic weekyear date and time without
// millis, separated by a T: xxxx'W'wwe'T'HHmmssZ.
//
//  "date" || "strict_date"
// A formatter for a full date as four digit year, two digit month of year, and
// two digit day of month: yyyy-MM-dd.
//
//  "date_hour" || "strict_date_hour"
// A formatter that combines a full date and two digit hour of day:
// yyyy-MM-dd'T'HH.
//
//  "date_hour_minute" || "strict_date_hour_minute"
// A formatter that combines a full date, two digit hour of day, and two digit
// minute of hour: yyyy-MM-dd'T'HH:mm.
//
//  "date_hour_minute_second" || "strict_date_hour_minute_second"
// A formatter that combines a full date, two digit hour of day, two digit
// minute of hour, and two digit second of minute: yyyy-MM-dd'T'HH:mm:ss.
//
//  "date_hour_minute_second_fraction" || "strict_date_hour_minute_second_fraction"
// A formatter that combines a full date, two digit hour of day, two digit
// minute of hour, two digit second of minute, and three digit fraction of
// second: yyyy-MM-dd'T'HH:mm:ss.SSS.
//
// "date_hour_minute_second_millis" || "strict_date_hour_minute_second_millis" A
// formatter that combines a full date, two digit hour of day, two digit minute
// of hour, two digit second of minute, and three digit fraction of second:
// yyyy-MM-dd'T'HH:mm:ss.SSS.
//
//  "date_time" || "strict_date_time"
// A formatter that combines a full date and time, separated by a T:
// yyyy-MM-dd'T'HH:mm:ss.SSSZZ.
//
// "date_time_no_millis" || "strict_date_time_no_millis" A formatter that
// combines a full date and time without millis, separated by a T:
// yyyy-MM-dd'T'HH:mm:ssZZ.
//
// "hour" || "strict_hour" A formatter for a two digit hour of day: HH
//
//  "hour_minute" || "strict_hour_minute"
// A formatter for a two digit hour of day and two digit minute of hour: HH:mm.
//
//  "hour_minute_second" || "strict_hour_minute_second"
// A formatter for a two digit hour of day, two digit minute of hour, and two
// digit second of minute: HH:mm:ss.
//
//  "hour_minute_second_fraction" || "strict_hour_minute_second_fraction"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, and three digit fraction of second: HH:mm:ss.SSS.
//
//  "hour_minute_second_millis" || "strict_hour_minute_second_millis"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, and three digit fraction of second: HH:mm:ss.SSS.
//
//  "ordinal_date" || "strict_ordinal_date"
// A formatter for a full ordinal date, using a four digit year and three digit
// dayOfYear: yyyy-DDD.
//
//  "ordinal_date_time" || "strict_ordinal_date_time"
// A formatter for a full ordinal date and time, using a four digit year and
// three digit dayOfYear: yyyy-DDD'T'HH:mm:ss.SSSZZ.
//
//  "ordinal_date_time_no_millis" || "strict_ordinal_date_time_no_millis"
// A formatter for a full ordinal date and time without millis, using a four
// digit year and three digit dayOfYear: yyyy-DDD'T'HH:mm:ssZZ.
//
//  "time" || "strict_time"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, three digit fraction of second, and time zone offset:
// HH:mm:ss.SSSZZ.
//
//  "time_no_millis" || "strict_time_no_millis"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, and time zone offset: HH:mm:ssZZ.
//
//  "t_time" || "strict_t_time"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, three digit fraction of second, and time zone offset
// prefixed by T: 'T'HH:mm:ss.SSSZZ.
//
//  "t_time_no_millis" || "strict_t_time_no_millis"
// A formatter for a two digit hour of day, two digit minute of hour, two digit
// second of minute, and time zone offset prefixed by T: 'T'HH:mm:ssZZ.
//
//  "week_date" || "strict_week_date"
// A formatter for a full date as four digit weekyear, two digit week of
// weekyear, and one digit day of week: xxxx-'W'ww-e.
//
//  "week_date_time" || "strict_week_date_time"
// A formatter that combines a full weekyear date and time, separated by a T:
// xxxx-'W'ww-e'T'HH:mm:ss.SSSZZ.
//
//  "week_date_time_no_millis" || "strict_week_date_time_no_millis"
// A formatter that combines a full weekyear date and time without millis,
// separated by a T: xxxx-'W'ww-e'T'HH:mm:ssZZ.
//
//  "weekyear" || "strict_weekyear"
// A formatter for a four digit weekyear: xxxx.
//
//  "weekyear_week" || "strict_weekyear_week"
// A formatter for a four digit weekyear and two digit week of weekyear:
// xxxx-'W'ww.
//
//  "weekyear_week_day" || "strict_weekyear_week_day"
// A formatter for a four digit weekyear, two digit week of weekyear, and one
// digit day of week: xxxx-'W'ww-e.
//
//  "year" || "strict_year"
// A formatter for a four digit year: yyyy.
//
//  "year_month" || "strict_year_month"
// A formatter for a four digit year and two digit month of year: yyyy-MM.
//
//  "year_month_day" || "strict_year_month_day"
// A formatter for a four digit year, two digit month of year, and two digit day
// of month: yyyy-MM-dd.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-date-format.html
type FormatParam struct {
	FormatValue string `bson:"format,omitempty" json:"format,omitempty"`
}

//Format is the format(s) that the that can be parsed. Defaults to strict_date_optional_time||epoch_millis.
//
// Multiple formats can be seperated by ||
func (f FormatParam) Format() string {
	if f.FormatValue == "" {
		return "strict_date_optional_time||epoch_millis"
	}
	return f.FormatValue
}

func (f *FormatParam) SetFormat(v string) {
	if v != f.Format() {
		f.FormatValue = v
	}
}

func unmarshalFormatParam(value gjson.Result, target interface{}) error {
	if r, ok := target.(WithFormat); ok {
		r.SetFormat(value.Str)
	}
	return nil
}
