/*
Copyright 2017 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package tables

import (
	"encoding/json"
	"io/ioutil"
)

type Definitions struct {
	IsLabelArg        map[string]bool
	LabelBlacklist    map[string]bool
	IsSortableListArg map[string]bool
	SortableBlacklist map[string]bool
	SortableWhitelist map[string]bool
	NamePriority      map[string]int
}

func ParseJsonDefinitions(file string) (Definitions, error) {
	var definitions Definitions

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return definitions, err
	}

	err = json.Unmarshal(data, &definitions)
	return definitions, err
}

func ParseAndUpdateJsonDefinitions(file string, merge bool) error {
	definitions, err := ParseJsonDefinitions(file)
	if err != nil {
		return err
	}

	if merge {
		MergeTables(definitions.IsLabelArg, definitions.LabelBlacklist, definitions.IsSortableListArg, definitions.SortableBlacklist, definitions.SortableWhitelist, definitions.NamePriority)
	} else {
		OverrideTables(definitions.IsLabelArg, definitions.LabelBlacklist, definitions.IsSortableListArg, definitions.SortableBlacklist, definitions.SortableWhitelist, definitions.NamePriority)
	}
	return nil
}
