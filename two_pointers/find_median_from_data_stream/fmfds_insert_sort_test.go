package find_median_from_data_stream

import "testing"

func TestMFInsertSort(t *testing.T) {
	mf := ConstructorIS()
	mf.AddNum(6)
	t.Log(mf.FindMedian())
	mf.AddNum(10)
	t.Log(mf.FindMedian())
}
