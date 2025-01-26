package log

import (
	"errors"
	"os"
	"testing"
)

func TestNewWithNilOpts(t *testing.T) {
	_, err := New(os.Stdin, nil)
	wantError := errors.New("opts is nil")
	if err == nil {
		t.Fatalf("Get err: %v, Want: %v", err, wantError)
	}
}

func TestEmptyStr(t *testing.T) {
	get := Str("", "")
	want := Msg(strColor + " : " + resetColor)
	if get != want {
		t.Fatalf("Get: %v, Want: %v", get, want)
	}
}

func TestNilErr(t *testing.T) {
	get := Error(nil)
	want := Msg(errColor + "Error: " + resetColor)
	if get != want {
		t.Fatalf("Get: %v, Want: %v", get, want)
	}
}

func TestEmptyWarn(t *testing.T) {
	get := Warning("")
	want := Msg(warnColor + "Warning: " + resetColor)
	if get != want {
		t.Fatalf("Get: %v, Want: %v", get, want)
	}
}

func TestEmptyInfo(t *testing.T) {
	get := Info("")
	want := Msg(infoColor + "Info: " + resetColor)
	if get != want {
		t.Fatalf("Get: %v, Want: %v", get, want)
	}
}
