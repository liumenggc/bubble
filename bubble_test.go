//bubble_test.go
package bubble

import "testing"

//测试
func TestBublleSort(t *testing.T) {
	values := []int{2, 4, 1, 45, 32, 65, 23}
	BublleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 4 || values[3] != 23 || values[4] != 32 ||
		values[5] != 45 || values[6] != 65{
			t.Error("Bubblesort() failed. Got", values, "Expected 1 2 4 23 32 45 65 ")
	}
}