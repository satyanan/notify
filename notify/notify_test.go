package notify

import (
	"errors"
	"reflect"
	"testing"
)

func TestNotify(t *testing.T) {
	testCases := []struct {
		cmd string
		err error
	}{
		{"./art", errors.New("fork/exec ./art: no such file or directory")},
		{"ls", nil},
	}

	for _, tc := range testCases {
		err := Notify(tc.cmd, "")
		// TODO: There should be a better way to test this
		if err == nil && tc.err == nil {
			continue
		}
		if (err == nil && tc.err != nil) || (err != nil && tc.err == nil) {
			t.Errorf("Expected %s, Got %s", tc.err, err)
			continue
		}
		if reflect.DeepEqual(err.Error(), tc.err.Error()) == false {
			t.Errorf("Expected %s, Got %s", tc.err, err)
		}
	}
}
