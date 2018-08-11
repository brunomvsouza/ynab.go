// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRateLimit(t *testing.T) {
	table := []struct {
		In  string
		Out *RateLimit
		Err error
	}{
		{"1/10", &RateLimit{used: uint64(1), total: uint64(10)}, nil},
		{"10/10", &RateLimit{used: uint64(10), total: uint64(10)}, nil},
		{"13/10", &RateLimit{used: uint64(13), total: uint64(10)}, nil},
		{"/10", nil, errInvalidRateLimit},
		{"1/", nil, errInvalidRateLimit},
		{"1", nil, errInvalidRateLimit},
		{"", nil, errInvalidRateLimit},
		{"a/a", nil, errInvalidRateLimit},
		{"/a", nil, errInvalidRateLimit},
		{"a/", nil, errInvalidRateLimit},
		{"a", nil, errInvalidRateLimit},
	}

	for _, test := range table {
		rl, err := ParseRateLimit(test.In)
		assert.Equal(t, test.Err, err)
		assert.Equal(t, test.Out, rl)
	}
}
