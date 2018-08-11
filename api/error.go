// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package api

import (
	"fmt"
)

// Error represents an API Error
type Error struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

// Error returns the string version of the error
func (e Error) Error() string {
	return fmt.Sprintf("api: error id=%s name=%s detail=%s",
		e.ID, e.Name, e.Detail)
}
