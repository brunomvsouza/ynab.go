// Copyright (c) 2019, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package api

import "fmt"

// Filter represents the optional version filter while
// fetching a budget
type Filter struct {
	// LastKnowledgeOfServer The starting server knowledge. If provided,
	// only entities that have changed since last_knowledge_of_server
	// will be included
	LastKnowledgeOfServer uint64
}

// ToQuery returns the filters as a HTTP query string
func (f *Filter) ToQuery() string {
	return fmt.Sprintf("last_knowledge_of_server=%d", f.LastKnowledgeOfServer)
}
