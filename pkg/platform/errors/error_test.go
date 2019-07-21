package errors

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var errSome = errors.New("some error")

func TestErrorCode(t *testing.T) {
	tt := []struct {
		tag  string
		err  error
		want string
	}{
		{
			tag:  "invalid_code-return_EINVALID",
			err:  New(EINVALID, "test", "testop", errors.New("test")),
			want: EINVALID,
		},
		{
			tag:  "empty_code-return_EINTERNAL",
			err:  New("", "test", "testop", errors.New("test")),
			want: EINTERNAL,
		},
		{
			tag:  "not_our_error-return_EINTERNAL",
			err:  errors.New("test"),
			want: EINTERNAL,
		},
		{
			tag:  "nil-return_empty_string",
			err:  nil,
			want: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			got := Code(tc.err)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestErrorMessage(t *testing.T) {
	tt := []struct {
		tag  string
		err  error
		want string
	}{
		{
			tag:  "have_message-return_message",
			err:  New(EINVALID, "test message", "testop", errSome),
			want: "test message",
		},
		{
			tag:  "no_message-return_default",
			err:  New(EINVALID, "", "testop", errSome),
			want: DefaultErrMessage,
		},
		{
			tag:  "not_our_error-return_default",
			err:  errSome,
			want: DefaultErrMessage,
		},
		{
			tag:  "nil-return_empty_string",
			err:  nil,
			want: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			got := Message(tc.err)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestError(t *testing.T) {
	tt := []struct {
		tag  string
		err  *Error
		want string
	}{
		{
			tag:  "full_form",
			err:  New(EINVALID, "test message", "testop", errSome),
			want: "testop: [invalid] test message. details: some error.",
		},
		{
			tag:  "missing_error",
			err:  New(EINVALID, "test message", "testop", nil),
			want: "testop: [invalid] test message. ",
		},
		{
			tag:  "nil_return-empty-string",
			err:  nil,
			want: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			got := tc.err.Error()

			assert.Equal(t, tc.want, got)
		})
	}
}
func TestWrap(t *testing.T) {
	tt := []struct {
		tag     string
		err     error
		message string
		want    error
	}{
		{
			tag:     "not_our_nil_error-return-nil",
			err:     nil,
			message: "",
			want:    nil,
		},
		{
			tag:     "nil_return-nil",
			err:     nil,
			message: "",
			want:    nil,
		},
		{
			tag:     "wrap_message",
			err:     New(EINVALID, "", "testop", errSome),
			message: "test message",
			want: &Error{
				Code:    "invalid",
				Message: "test message: ",
				Op:      "testop",
				Err:     errSome,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			got := Wrap(tc.err, tc.message)

			assert.Equal(t, tc.want, got)
		})
	}
}
