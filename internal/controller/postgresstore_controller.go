/*
Copyright 2024.

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

package controller

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	postgressgroupv1 "github.com/venkateshsredhat/postgress-gorm-operator/api/v1"
	dbconnect "github.com/venkateshsredhat/postgress-gorm-operator/postgress"
)

// PostgresStoreReconciler reconciles a PostgresStore object
type PostgresStoreReconciler struct {
	client.Client
	Scheme       *runtime.Scheme
	DbConnection *gorm.DB
}

// +kubebuilder:rbac:groups=postgressgroup.venkateshredhat.com,resources=postgresstores,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=postgressgroup.venkateshredhat.com,resources=postgresstores/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=postgressgroup.venkateshredhat.com,resources=postgresstores/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PostgresStore object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *PostgresStoreReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	//Creates the table
	r.DbConnection.AutoMigrate(&dbconnect.Quest{})
	store := &postgressgroupv1.PostgresStore{}
	err := r.Client.Get(ctx, req.NamespacedName, store)
	if err != nil {
		return ctrl.Result{}, err
	}
	quest := &dbconnect.Quest{
		ID:    uint(store.Spec.ID),
		Title: store.Spec.Title,
	}
	var find dbconnect.Quest

	if err := r.DbConnection.Where("id = ?", quest.ID).First(&find).Error; err != nil {
		fmt.Println("quest not found : ", err)
		r.DbConnection.Create(quest)
		fmt.Println("Quest created successfully")
		return ctrl.Result{}, err
	}

	fmt.Println("Already present in database", find)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PostgresStoreReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&postgressgroupv1.PostgresStore{}).
		Complete(r)
}
