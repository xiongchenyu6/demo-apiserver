/*


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

package main

import (
	"k8s.io/klog"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"

	// +kubebuilder:scaffold:resource-imports
	demov1beta1 "github.com/xiongchenyu6/demo-apiserver/pkg/apis/demo/v1beta1"
	openapi "github.com/xiongchenyu6/demo-apiserver/pkg/openapi"
)

func main() {
	err := builder.APIServer.
		WithOpenAPIDefinitions("api", "v1beta1", openapi.GetOpenAPIDefinitions).
		WithLocalDebugExtension().
		// +kubebuilder:scaffold:resource-register
		WithResource(&demov1beta1.Foo{}).
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
