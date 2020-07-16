package ingressmonitor

import (
	"fmt"
	ingressmonitorv1alpha1 "github.com/stakater/IngressMonitorController/pkg/apis/ingressmonitor/v1alpha1"
	"github.com/stakater/IngressMonitorController/pkg/models"
	"github.com/stakater/IngressMonitorController/pkg/monitors"
	"github.com/stakater/IngressMonitorController/pkg/kube/util"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func (r *ReconcileIngressMonitor) handleCreate(request reconcile.Request, instance *ingressmonitorv1alpha1.IngressMonitor, monitorName string, monitorService monitors.MonitorServiceProxy) (reconcile.Result, error) {
	log.Info("Creating Monitor: " + monitorName)

	// TODO: REMOVE THIS
	fmt.Printf("%+v\n", instance.Spec)

 	// TODO:
 	// Replace annotations with provider specific configuration IF ANY
	url, err := util.GetMonitorURL(r.client, instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	monitor := models.NewMonitor(monitorName, url)

	// TODO: Generate error and handle it
	// Add monitor
	monitorService.Add(monitor)

	return reconcile.Result{RequeueAfter: defaultRequeueTime}, nil
}
