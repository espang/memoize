package memoize_test

import (
	"github.com/espang/memoize"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemoize(t *testing.T) {
	type testcase struct {
		label   string
		memoize func(func(int) int) func(int) int
	}

	testcases := []testcase{
		{
			label:   "simple map",
			memoize: memoize.M,
		},
		{
			label:   "map with FiFo",
			memoize: memoize.MFiFo,
		},
		{
			label:   "map with lru list",
			memoize: memoize.MLRU,
		},
	}

	tf := func(i int) int {
		time.Sleep(1 * time.Second)
		return 15 * i
	}

	for _, tc := range testcases {
		t.Run(tc.label, func(t *testing.T) {
			mf := tc.memoize(tf)
			assert.Equal(t, tf(12), mf(12))
		})
	}
}
