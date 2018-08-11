// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package api

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	errInvalidRateLimit = errors.New(`api: invalid rate limit string`)

	rateLimitRegex = regexp.MustCompile(`^(?P<used>[0-9]+)/(?P<total>[0-9]+)$`)
)

// RateLimit represents an API rate limit
type RateLimit struct {
	used  uint64
	total uint64
}

// Used represents the used rate limit
func (r *RateLimit) Used() uint64 {
	return r.used
}

// Total represents the total rate limit
func (r *RateLimit) Total() uint64 {
	return r.total
}

// ParseRateLimit returns a *RateLimit for a given rate limit string
func ParseRateLimit(rateLimit string) (*RateLimit, error) {
	m := rateLimitRegex.FindStringSubmatch(rateLimit)
	if len(m) != 3 {
		return nil, errInvalidRateLimit
	}

	used, err := strconv.ParseUint(m[1], 10, 64)
	if err != nil {
		return nil, err
	}

	total, err := strconv.ParseUint(m[2], 10, 64)
	if err != nil {
		return nil, err
	}

	return &RateLimit{
		used:  used,
		total: total,
	}, nil
}
