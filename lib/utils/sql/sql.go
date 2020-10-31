package sql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func RepeatQuestionMarks(n int) string {
	questionMarks := []string{}
	for i := 0; i < n; i++ {
		questionMarks = append(questionMarks, "?")
	}
	return strings.Join(questionMarks, ",")
}

func MakeSqlIn(col string, vals interface{}) (string, []interface{}, error) {
	switch reflect.TypeOf(vals).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(vals)
		if s.Len() == 0 {
			return "", []interface{}{}, errors.New("'vals' has no value")
		}
		args := []interface{}{}
		for i := 0; i < s.Len(); i++ {
			args = append(args, s.Index(i).Interface())
		}
		return fmt.Sprintf("%s IN (%s)", col, RepeatQuestionMarks(s.Len())), args, nil
	}

	return "", []interface{}{}, errors.New("'vals' must be a slice")
}

func MakeSqlLimitOffset(limit uint32, offset uint64) string {
	var l uint32 = 10
	var o uint64 = 0

	if limit > 0 {
		l = limit
	}

	if offset > 0 {
		o = offset
	}

	return fmt.Sprintf("limit %d offset %d", l, o)
}