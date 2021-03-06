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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	foos "github.com/xiongchenyu6/demo-apiserver/pkg/apis/demo/v1beta1"
)

// FooReconciler reconciles a Foo object
type FooReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=,resources=foos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=,resources=foos/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=,resources=foos/finalizers,verbs=update

func (r *FooReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("foo", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *FooReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&foos.Foo{}).
		Complete(r)
}
