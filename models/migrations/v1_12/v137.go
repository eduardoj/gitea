// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1_12 //nolint

import (
	"xorm.io/xorm"
)

func AddBlockOnOutdatedBranch(x *xorm.Engine) error {
	type ProtectedBranch struct {
		BlockOnOutdatedBranch bool `xorm:"NOT NULL DEFAULT false"`
	}
	return x.Sync2(new(ProtectedBranch))
}
