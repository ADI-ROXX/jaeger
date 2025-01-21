// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package features

import (
	"fmt"

	"go.opentelemetry.io/collector/featuregate"
)

type flagValue struct {
	reg *featuregate.Registry
}

func DisplayFeatures() {
	reg := featuregate.GlobalRegistry()
	f := &flagValue{reg: reg}
	if f.reg == nil {
		return
	}
	f.reg.VisitAll(func(g *featuregate.Gate) {
		id := g.ID()
		desc := g.Description()
		fmt.Println("Feature:\t" + id)
		if !g.IsEnabled() {
			fmt.Println("Default state:\t" + "On")
		} else {
			fmt.Println("Default state:\t" + "Off")
		}
		fmt.Println("Description:\t" + desc)
		ref := g.ReferenceURL()
		if ref != "" {
			fmt.Println("ReferenceURL:\t" + ref)
		}
		fmt.Println()
	})
}
