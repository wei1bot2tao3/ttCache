package OSCache

import (
	"context"
	"fmt"
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
			name:       "cacheOne 插入一次",
			key:        "oneKey",
			value:      "oneValue",
			wantRes:    "oneValue",
			wantErr:    errs.NewErrNotfound("oneKey"),
			expiration: 3 * time.Second,
			Cache: func() *CacheOneGo {
				build := NewBuildInMapCache(10)
				cache := NewBuildInMapCacheOneGo(build, 5*time.Second)
				err := cache.Set(ctx, "oneKey", "oneValue", 3*time.Second)
				require.NoError(t, err)
				return cache
			},
		},
		{
			name:       "cacheOne 插入两次",
			key:        "oneKey",
			value:      "oneValue",
			wantRes:    "oneValue",
			wantErr:    errs.NewErrNotfound("oneKey"),
			expiration: 3 * time.Second,
			Cache: func() *CacheOneGo {
				build := NewBuildInMapCache(10)
				cache := NewBuildInMapCacheOneGo(build, 5*time.Second)
				err := cache.Set(ctx, "oneKey", "oneValue", 3*time.Second)
				require.NoError(t, err)
				err = cache.Set(ctx, "oneKey", "oneValue", 3*time.Second)
				assert.Equal(t, errs.ErrKeyExists, err)

				return cache
			},
		},
	}
	for _, tc := range testCass {
		t.Run(tc.name, func(t *testing.T) {
			cache := tc.Cache()

			res, err := cache.Get(ctx, tc.key)
			require.NoError(t, err)
			itme, ok := res.(*item)
			assert.True(t, ok)
			assert.Equal(t, tc.wantRes, itme.val)
			time.Sleep(6 * time.Second)
			_, err = cache.Get(ctx, tc.key)
			assert.Equal(t, tc.wantErr, err)
			err = cache.Set(ctx, "oneKey", "oneValue", 3*time.Second)
			val, err := cache.Delete(ctx, "oneKey")
			fmt.Println(err, val)
			err = cache.Set(ctx, "oneKey", "oneValue", 3*time.Second)
			fmt.Println(err)

		})
	}

}

func TestNewBuildInMapCacheGos(t *testing.T) {
	ctx := context.Background()
	testCass := []struct {
		name       string
		Cache      func() *CacheGos
		key        string
		value      string
		wantErr    error
		wantRes    string
		expiration time.Duration
	}{
		{
			name: "插入一次",
			Cache: func() *CacheGos {
				build := NewBuildInMapCache(10)
				cache := NewBuildInMapCacheGos(build)
				err := cache.Set(ctx, "key1", "value1", 5*time.Second)
				assert.NoError(t, err)
				return cache
			},
			key:        "key1",
			value:      "value1",
			wantRes:    "value1",
			expiration: 2 * time.Second,
			wantErr:    errs.NewErrNotfound("key1"),
		},
	}
	for _, tc := range testCass {
		t.Run(tc.name, func(t *testing.T) {
			cache := tc.Cache()
			Map, err := cache.Get(ctx, tc.key)
			assert.NoError(t, err)
			node, ok := Map.(*item)
			assert.True(t, ok)
			assert.Equal(t, tc.wantRes, node.val)
			fmt.Println(node.val)
			time.Sleep(4 * time.Second)
			node1, err := cache.Get(ctx, tc.key)
			fmt.Println(node1)

		})
	}
}

func TestNewBuildInMapCacheNoGo(t *testing.T) {
	testCass := []struct {
		name string

		wantErr error
		wantRes map[string]any
	}{
		{},
	}
	for _, tc := range testCass {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}