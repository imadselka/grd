package grd

import (
	"errors"
	"testing"
)

func TestTrySuccess(t *testing.T) {
	val := Try(func() (int, error) {
		return 42, nil
	}).Catch(func(err error) int {
		return -1
	})

	if val != 42 {
		t.Errorf("expected 42, got %d", val)
	}
}

func TestTryFailure(t *testing.T) {
	val := Try(func() (int, error) {
		return 0, errors.New("boom")
	}).Catch(func(err error) int {
		return -1
	})

	if val != -1 {
		t.Errorf("expected -1, got %d", val)
	}
}

func TestTryFinally(t *testing.T) {
	var called bool
	Try(func() (int, error) {
		return 42, nil
	}).Finally(func() {
		called = true
	})

	if !called {
		t.Error("Finally block was not called")
	}
}