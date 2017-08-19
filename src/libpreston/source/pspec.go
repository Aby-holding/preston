//
// Copyright © 2017 Ikey Doherty <ikey@solus-project.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package source

// pspecPackage is the simplest possible representation of a pspec.xml package
// whilst providing access to the things we care about: licenses and name
type pspecPackage struct {
	Source struct {
		Name    string
		License []string
	}
}

// NewEopkgPackageLegacy will return a PSPEC parsed source.Package
func NewEopkgPackageLegacy(path string) (*Package, error) {
	return nil, ErrNotYetImplemented
}
