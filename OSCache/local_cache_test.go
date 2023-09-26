package OSCache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"ttCache/internal/errs"
)

func TestNewBuildInMapCache(t *testing.T) {
	ctx := context.Background()
	testCass := []struct {
		name       string
		Cache      func() *CacheOneGo
		key        string
		value      string
		wantErr    error
		wantRes    string
		expiration time.Duration
	}{
		{
			name:       "cacheOne",
			key:        "oneKey",
			value:      "oneValue",
			wantRes:    "oneValue",
			wantErr:    errs.NewErrNotfound("oneKey"),
			expiration: 3 * time.Second,
		},
		{
			name:       "cacheOne",
			key:        "oneKey",
			value:      "oneValue",
			wantRes:    "oneValue",
			wantErr:    errs.NewErrNotfound("oneKey"),
			expiration: 3 * time.Second,
		},
	}
	for _, tc := range testCass {
		t.Run(tc.name, func(t *testing.T) {
			build := NewBuildInMapCache(10)
			cache := NewBuildInMapCacheOneGo(build, 5*time.Second)
			err := cache.Set(ctx, tc.key, tc.value, tc.expiration)
			require.NoError(t, err)
			res, err := cache.Get(ctx, tc.key)
			require.NoError(t, err)
			itme, ok := res.(*item)
			assert.True(t, ok)
			assert.Equal(t, tc.wantRes, itme.val)
			time.Sleep(6 * time.Second)
			_, err = cache.Get(ctx, tc.key)
			assert.Equal(t, tc.wantErr, err)

		})
	}

}
