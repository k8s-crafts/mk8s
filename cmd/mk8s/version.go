// Copyright 2024 k8s-crafts Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mk8s

import (
	"github.com/k8s-crafts/mk8s/pkg/mk8s/out"
	"github.com/k8s-crafts/mk8s/pkg/mk8s/version"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version of mk8s",
	Long:  "Print the version of mk8s",
	Run: func(cmd *cobra.Command, args []string) {
		out.PrintLn(version.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
