// Copyright 2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prolly

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTreeChunker(t *testing.T) {
	t.Run("round trip tree items", func(t *testing.T) {
		roundTripTreeItems(t)
	})
}

func roundTripTreeItems(t *testing.T) {
	root, items, nrw := randomTree(t, 1000)
	assert.NotNil(t, root)
	assert.True(t, root.nodeCount() > 0)
	assert.True(t, root.level() > 0)
	assert.Equal(t, uint64(1000), root.cumulativeCount())
	assert.Equal(t, countTree(t, nrw, root), 1000)
	validateTreeItems(t, nrw, root, items)

	root, items, nrw = randomTree(t, 10_000)
	assert.NotNil(t, root)
	assert.True(t, root.nodeCount() > 0)
	assert.True(t, root.level() > 0)
	assert.Equal(t, uint64(10_000), root.cumulativeCount())
	assert.Equal(t, countTree(t, nrw, root), 10_000)
	validateTreeItems(t, nrw, root, items)

	root, items, nrw = randomTree(t, 100_000)
	assert.NotNil(t, root)
	assert.True(t, root.nodeCount() > 0)
	assert.True(t, root.level() > 0)
	assert.Equal(t, uint64(100_000), root.cumulativeCount())
	assert.Equal(t, countTree(t, nrw, root), 100_000)
	validateTreeItems(t, nrw, root, items)
}

func countTree(t *testing.T, nrw NodeReadWriter, nd Node) (count int) {
	ctx := context.Background()
	err := iterTree(ctx, nrw, nd, func(_ nodeItem) (err error) {
		count++
		return
	})
	require.NoError(t, err)
	return
}

func validateTreeItems(t *testing.T, nrw NodeReadWriter, nd Node, expected [][2]nodeItem) {
	i := 0
	ctx := context.Background()
	err := iterTree(ctx, nrw, nd, func(actual nodeItem) (err error) {
		if !assert.Equal(t, expected[i/2][i%2], actual) {
			panic("here")
		}
		i++
		return
	})
	require.NoError(t, err)
	return
}

func iterTree(ctx context.Context, nrw NodeReadWriter, nd Node, cb func(item nodeItem) error) error {
	if nd.empty() {
		return nil
	}

	cur, err := newCursor(ctx, nrw, nd)
	if err != nil {
		return err
	}

	ok := true
	for ok {
		curr := cur.current()

		err = cb(curr)
		if err != nil {
			return err
		}

		ok, err = cur.advance(ctx)
		if err != nil {
			return err
		}
	}
	return err
}