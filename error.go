package main

import (
	"errors"
	"fmt"
)

// Sentinel error: a predefined, reusable error variable
var ErrZeroInput = errors.New("input is zero")

// f1 returns the reciprocal of a float.
// It demonstrates both direct and wrapped sentinel error usage.
func f1(num float64, wrap bool) (float64, error) {
	if num == 0 {
		if wrap {
			// Wrapping the sentinel error using fmt.Errorf + %w allows it to be unwrapped later
			return -1, fmt.Errorf("f1 failed: %w", ErrZeroInput)
		}
		// Returning the sentinel error directly (no wrapping)
		return -1, ErrZeroInput
	}
	return 1.0 / num, nil
}

// Custom error type with fields for richer error information
type myError struct {
	num  int
	prob string
}

// Implement the error interface for custom error type
func (e myError) Error() string {
	return fmt.Sprintf("%d - %s", e.num, e.prob)
}

// f2 returns a reciprocal of int and demonstrates a custom error type
func f2(num int) (int, error) {
	if num == 0 {
		return -1, myError{num, "it is zero"}
	}
	return int(1.0 / float64(num)), nil
}

func main() {
	// --- Using sentinel error without wrapping ---
	res1, err1 := f1(0, false)
	fmt.Println("f1 (no wrap):", res1, err1)

	// You can compare directly using == when not wrapped
	if err1 == ErrZeroInput {
		fmt.Println("✅ err1 matched using == : sentinel error matched directly")
	}

	// You can still use errors.Is even if the error is not wrapped — it falls back to ==
	if errors.Is(err1, ErrZeroInput) {
		fmt.Println("✅ err1 matched using errors.Is : fallback to == worked")
	}

	// --- Using sentinel error WITH wrapping ---
	res2, err2 := f1(0, true)
	fmt.Println("f1 (with wrap):", res2, err2)

	// This won't work: err2 != ErrZeroInput because it’s wrapped
	if err2 == ErrZeroInput {
		fmt.Println("❌ err2 == ErrZeroInput (won’t match because it's wrapped)")
	}

	// Use errors.Is to match the underlying sentinel error inside a wrapped error
	if errors.Is(err2, ErrZeroInput) {
		fmt.Println("✅ err2 matched using errors.Is : correctly detected wrapped sentinel error")
	}

	// --- Wrong way: wrapping using %v (not detectable by errors.Is) ---
	errWrong := fmt.Errorf("bad wrap: %v", ErrZeroInput)
	if errors.Is(errWrong, ErrZeroInput) {
		fmt.Println("❌ matched incorrectly")
	} else {
		fmt.Println("🚫 errWrong: errors.Is failed as expected (because %v doesn’t wrap)")
	}

	// --- f2 custom error type ---
	res3, err3 := f2(10)
	fmt.Println("f2 (valid input):", res3, err3)

	res4, err4 := f2(0)
	fmt.Println("f2 (custom error):", res4, err4)

	// Access fields using type assertion (custom error)
	if ae, ok := err4.(myError); ok {
		fmt.Println("🛠️ Custom error details:")
		fmt.Println("- num:", ae.num)
		fmt.Println("- prob:", ae.prob)
	}

	// ⚠️ NOTE: errors.Is works by comparing known sentinel values or wrapped chains.
	// But here we're creating a *new instance* of myError on the fly,
	// so even if the content is same, the pointer/reference is different.
	// Hence, `errors.Is` returns false — it's not meant for custom struct matching.
	if errors.Is(err4, myError{0, "it is zero"}) {
		fmt.Println("✅ matched custom error? (unexpected!)")
	} else {
		// ✅ Correct expectation: errors.Is fails to match two different struct instances.
		fmt.Println("🚫 errors.Is cannot match custom error types by value — use type assertion instead")
	}

	// ✅ Best practice:
	// Use type assertions (as shown above) to match custom error types
	// and access their fields for detailed logic or logging.

}
