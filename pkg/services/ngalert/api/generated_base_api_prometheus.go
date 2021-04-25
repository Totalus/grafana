/*Package api contains base API implementation of unified alerting
 *
 *Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 *Do not manually edit these files, please find ngalert/api/swagger-codegen/ for commands on how to generate them.
 */

package api

import (
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/models"
)

type PrometheusApiService interface {
	RouteGetAlertStatuses(*models.ReqContext) response.Response
	RouteGetRuleStatuses(*models.ReqContext) response.Response
}

func (api *API) RegisterPrometheusApiEndpoints(srv PrometheusApiService) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Get(toMacaronPath("/api/prometheus/{Recipient}/api/v1/alerts"), routing.Wrap(srv.RouteGetAlertStatuses))
		group.Get(toMacaronPath("/api/prometheus/{Recipient}/api/v1/rules"), routing.Wrap(srv.RouteGetRuleStatuses))
	}, middleware.ReqSignedIn)
}