package intutils

import (
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans == -2 {

		//t.Error* will report test failures but continue executing the test.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)

		//t.Fatal* will report test failures and stop the test immediately.
		t.Fatalf("IntMin(2, -2) = %d; want -2", ans)
	}
}
