/*
 * Copyright 2018-2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package contrast

import (
	"fmt"

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak/bard"
	"github.com/paketo-buildpacks/libpak/bindings"
)

type Detect struct {
	Logger bard.Logger
}

func (d Detect) Detect(context libcnb.DetectContext) (libcnb.DetectResult, error) {
	if _, ok, err := bindings.ResolveOne(context.Platform.Bindings, bindings.OfType("ContrastSecurity")); err != nil {
		return libcnb.DetectResult{}, fmt.Errorf("unable to resolve binding ContrastSecurity\n%w", err)
	} else if !ok {
		d.Logger.Info("SKIPPED: No binding of type 'ContrastSecurity' found")
		return libcnb.DetectResult{Pass: false}, nil
	}

	return libcnb.DetectResult{
		Pass: true,
		Plans: []libcnb.BuildPlan{
			{
				Provides: []libcnb.BuildPlanProvide{
					{Name: "contrast-java"},
				},
				Requires: []libcnb.BuildPlanRequire{
					{Name: "contrast-java"},
					{Name: "jvm-application"},
				},
			},
		},
	}, nil
}
