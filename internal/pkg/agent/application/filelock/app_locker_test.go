// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package filelock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testLockFile = "test.lock"

func TestAppLocker(t *testing.T) {
	tmp := t.TempDir()

	locker1 := NewAppLocker(tmp, testLockFile)
	locker2 := NewAppLocker(tmp, testLockFile)

	require.NoError(t, locker1.TryLock())
	assert.ErrorIs(t, locker2.TryLock(), ErrAppAlreadyRunning)
	require.NoError(t, locker1.Unlock())
	require.NoError(t, locker2.TryLock())
	assert.ErrorIs(t, locker1.TryLock(), ErrAppAlreadyRunning)
	require.NoError(t, locker2.Unlock())
}
