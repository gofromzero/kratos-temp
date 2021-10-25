// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2021/10/25

package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase)
