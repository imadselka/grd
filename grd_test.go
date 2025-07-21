package grd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// Basic functionality tests
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

// Then method tests
func TestThenChainSuccess(t *testing.T) {
	result := Try(func() (int, error) {
		return 10, nil
	}).Then(func(val int) (int, error) {
		return val * 2, nil
	}).Then(func(val int) (int, error) {
		return val + 5, nil
	}).Catch(func(err error) int {
		return -1
	})

	expected := 25 // ((10 * 2) + 5)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestThenChainWithError(t *testing.T) {
	result := Try(func() (int, error) {
		return 10, nil
	}).Then(func(val int) (int, error) {
		return 0, errors.New("error in then")
	}).Then(func(val int) (int, error) {
		t.Error("This Then should not be called")
		return val + 5, nil
	}).Catch(func(err error) int {
		return -999
	})

	if result != -999 {
		t.Errorf("expected -999, got %d", result)
	}
}

func TestThenWithInitialError(t *testing.T) {
	thenCalled := false
	result := Try(func() (int, error) {
		return 0, errors.New("initial error")
	}).Then(func(val int) (int, error) {
		thenCalled = true
		return val * 2, nil
	}).Catch(func(err error) int {
		return -1
	})

	if thenCalled {
		t.Error("Then should not be called when there's an initial error")
	}
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
}

// Different types tests
func TestTryWithStringType(t *testing.T) {
	result := Try(func() (string, error) {
		return "hello", nil
	}).Then(func(s string) (string, error) {
		return s + " world", nil
	}).Catch(func(err error) string {
		return "error"
	})

	if result != "hello world" {
		t.Errorf("expected 'hello world', got '%s'", result)
	}
}

func TestTryWithCustomStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	result := Try(func() (User, error) {
		return User{Name: "Alice", Age: 30}, nil
	}).Then(func(u User) (User, error) {
		u.Age += 1
		return u, nil
	}).Catch(func(err error) User {
		return User{Name: "Error", Age: -1}
	})

	if result.Name != "Alice" || result.Age != 31 {
		t.Errorf("expected User{Name: 'Alice', Age: 31}, got %+v", result)
	}
}

func TestTryWithSliceType(t *testing.T) {
	result := Try(func() ([]int, error) {
		return []int{1, 2, 3}, nil
	}).Then(func(slice []int) ([]int, error) {
		return append(slice, 4), nil
	}).Catch(func(err error) []int {
		return []int{}
	})

	expected := []int{1, 2, 3, 4}
	if len(result) != len(expected) {
		t.Errorf("expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("expected %v, got %v", expected, result)
			break
		}
	}
}

// Finally tests
func TestFinallyCalledOnSuccess(t *testing.T) {
	finallyCalled := false
	result := Try(func() (int, error) {
		return 42, nil
	}).Finally(func() {
		finallyCalled = true
	}).Catch(func(err error) int {
		return -1
	})

	if !finallyCalled {
		t.Error("Finally should be called on success")
	}
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestFinallyCalledOnError(t *testing.T) {
	finallyCalled := false
	result := Try(func() (int, error) {
		return 0, errors.New("test error")
	}).Finally(func() {
		finallyCalled = true
	}).Catch(func(err error) int {
		return -1
	})

	if !finallyCalled {
		t.Error("Finally should be called on error")
	}
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
}

func TestMultipleFinally(t *testing.T) {
	count := 0
	Try(func() (int, error) {
		return 42, nil
	}).Finally(func() {
		count++
	}).Finally(func() {
		count++
	}).Catch(func(err error) int {
		return -1
	})

	if count != 2 {
		t.Errorf("expected count to be 2, got %d", count)
	}
}

func TestFinallyWithThenChain(t *testing.T) {
	var order []string

	result := Try(func() (int, error) {
		order = append(order, "try")
		return 10, nil
	}).Then(func(val int) (int, error) {
		order = append(order, "then1")
		return val * 2, nil
	}).Finally(func() {
		order = append(order, "finally1")
	}).Then(func(val int) (int, error) {
		order = append(order, "then2")
		return val + 5, nil
	}).Finally(func() {
		order = append(order, "finally2")
	}).Catch(func(err error) int {
		order = append(order, "catch")
		return -1
	})

	expectedOrder := []string{"try", "then1", "finally1", "then2", "finally2"}
	if len(order) != len(expectedOrder) {
		t.Errorf("expected order length %d, got %d", len(expectedOrder), len(order))
	}
	for i, expected := range expectedOrder {
		if i >= len(order) || order[i] != expected {
			t.Errorf("expected order %v, got %v", expectedOrder, order)
			break
		}
	}
	if result != 25 {
		t.Errorf("expected 25, got %d", result)
	}
}

// Error handling edge cases
func TestCatchWithNilError(t *testing.T) {
	// This shouldn't happen in practice, but let's test it
	tryResult := &TryResult[int]{result: 42, err: nil}
	result := tryResult.Catch(func(err error) int {
		t.Error("Catch should not be called when err is nil")
		return -1
	})

	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestErrorPropagation(t *testing.T) {
	originalErr := errors.New("original error")
	var caughtErr error

	Try(func() (int, error) {
		return 0, originalErr
	}).Catch(func(err error) int {
		caughtErr = err
		return -1
	})

	if caughtErr != originalErr {
		t.Error("Error should be propagated correctly")
	}
}

// Complex chaining scenarios
func TestComplexChaining(t *testing.T) {
	result := Try(func() (string, error) {
		return "5", nil
	}).Then(func(s string) (string, error) {
		// Convert string to int, then back to string with multiplication
		val, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(val * 10), nil
	}).Then(func(s string) (string, error) {
		return s + "0", nil
	}).Finally(func() {
		// Side effect
	}).Catch(func(err error) string {
		return "error: " + err.Error()
	})

	if result != "500" {
		t.Errorf("expected '500', got '%s'", result)
	}
}

func TestRealWorldFileOperation(t *testing.T) {
	// Simulate a file reading operation
	result := Try(func() (string, error) {
		// Simulate reading a file
		return "file content\nline 2\nline 3", nil
	}).Then(func(content string) (string, error) {
		// Process the content (count lines)
		lines := strings.Split(content, "\n")
		return fmt.Sprintf("Line count: %d", len(lines)), nil
	}).Then(func(processed string) (string, error) {
		// Add timestamp
		return processed + " (processed)", nil
	}).Catch(func(err error) string {
		return "Failed to process file: " + err.Error()
	})

	expected := "Line count: 3 (processed)"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestDatabaseOperationSimulation(t *testing.T) {
	type Record struct {
		ID   int
		Name string
	}

	// Simulate database operations
	result := Try(func() (Record, error) {
		// Simulate fetching from database
		return Record{ID: 1, Name: "John"}, nil
	}).Then(func(record Record) (Record, error) {
		// Validate record
		if record.Name == "" {
			return record, errors.New("name cannot be empty")
		}
		return record, nil
	}).Then(func(record Record) (Record, error) {
		// Transform record
		record.Name = strings.ToUpper(record.Name)
		return record, nil
	}).Catch(func(err error) Record {
		return Record{ID: -1, Name: "ERROR"}
	})

	if result.ID != 1 || result.Name != "JOHN" {
		t.Errorf("expected Record{ID: 1, Name: 'JOHN'}, got %+v", result)
	}
}

// Panic recovery test (if you want to add panic recovery to your library)
func TestWithPanicInTry(t *testing.T) {
	// Note: Current implementation doesn't handle panics
	// This test demonstrates what would happen
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic to be propagated")
		}
	}()

	Try(func() (int, error) {
		panic("something went wrong")
	}).Catch(func(err error) int {
		return -1
	})
}

// Performance test
func BenchmarkTryChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Try(func() (int, error) {
			return i, nil
		}).Then(func(val int) (int, error) {
			return val * 2, nil
		}).Then(func(val int) (int, error) {
			return val + 1, nil
		}).Catch(func(err error) int {
			return -1
		})
	}
}

func BenchmarkTryWithError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Try(func() (int, error) {
			return 0, errors.New("test error")
		}).Then(func(val int) (int, error) {
			return val * 2, nil
		}).Catch(func(err error) int {
			return -1
		})
	}
}
