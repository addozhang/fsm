package debugger

import (
	"encoding/json"
	"fmt"
	"net/http"

	access "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha3"
	spec "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha4"
	split "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha4"

	"github.com/flomesh-io/fsm/pkg/identity"
)

type policies struct {
	TrafficSplits   []*split.TrafficSplit        `json:"traffic_splits"`
	ServiceAccounts []identity.K8sServiceAccount `json:"service_accounts"`
	RouteGroups     []*spec.HTTPRouteGroup       `json:"route_groups"`
	TrafficTargets  []*access.TrafficTarget      `json:"traffic_targets"`
}

func (ds DebugConfig) getFSMConfigHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		confJSON, err := ds.configurator.GetMeshConfigJSON()
		if err != nil {
			log.Error().Err(err).Msg("error getting MeshConfig JSON")
			return
		}
		_, _ = fmt.Fprint(w, confJSON)
	})
}

func (ds DebugConfig) getSMIPoliciesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p policies
		p.TrafficSplits, p.ServiceAccounts, p.RouteGroups, p.TrafficTargets = ds.meshCatalogDebugger.ListSMIPolicies()

		jsonPolicies, err := json.Marshal(p)
		if err != nil {
			log.Error().Err(err).Msgf("Error marshalling policy %+v", p)
		}

		_, _ = fmt.Fprint(w, string(jsonPolicies))
	})
}
