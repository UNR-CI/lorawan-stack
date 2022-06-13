// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

package templates

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "temporary_password",
		email.FSTemplate{
			SubjectTemplate:      "Your temporary password for {{ .Network.Name }}",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "temporary_password.html.tmpl",
			TextTemplateFile:     "temporary_password.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
}

// TemporaryPasswordData is the data for the temporary_password email.
type TemporaryPasswordData struct {
	email.TemplateData
	TemporaryPassword string
	TTL               time.Duration
}