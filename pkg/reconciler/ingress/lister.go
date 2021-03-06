/*
Copyright 2020 The Knative Authors.

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

package ingress

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/sets"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"knative.dev/net-kourier/pkg/config"
	"knative.dev/net-kourier/pkg/knative"
	"knative.dev/pkg/network"
	"knative.dev/pkg/system"
	"knative.dev/serving/pkg/apis/networking/v1alpha1"
	"knative.dev/serving/pkg/network/status"
)

func NewProbeTargetLister(logger *zap.SugaredLogger, endpointsLister corev1listers.EndpointsLister) status.ProbeTargetLister {
	return &gatewayPodTargetLister{
		logger:          logger,
		endpointsLister: endpointsLister,
	}
}

type gatewayPodTargetLister struct {
	logger          *zap.SugaredLogger
	endpointsLister corev1listers.EndpointsLister
}

func (l *gatewayPodTargetLister) ListProbeTargets(ctx context.Context, ing *v1alpha1.Ingress) ([]status.ProbeTarget, error) {

	eps, err := l.endpointsLister.Endpoints(system.Namespace()).Get(config.InternalServiceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get internal service: %w", err)
	}

	var readyIPs []string
	for _, sub := range eps.Subsets {
		for _, address := range sub.Addresses {
			readyIPs = append(readyIPs, address.IP)
		}
	}
	if len(readyIPs) == 0 {
		return nil, fmt.Errorf("no gateway pods available")
	}
	return l.getIngressUrls(ing, readyIPs)
}

func (l *gatewayPodTargetLister) getIngressUrls(ing *v1alpha1.Ingress, gatewayIps []string) ([]status.ProbeTarget, error) {
	var targets []status.ProbeTarget
	localDomainName := network.GetClusterDomainName()
	ips := sets.NewString()

	for _, ip := range gatewayIps {
		ips.Insert(ip)
	}

	for _, rule := range ing.Spec.Rules {
		var target status.ProbeTarget

		externalDomains := getExternalDomains(rule, localDomainName)
		internalDomains := getInternalDomains(rule, localDomainName)
		scheme := "http"

		if knative.RuleIsExternal(rule, ing.Spec.Visibility) {
			target = status.ProbeTarget{
				PodIPs: ips,
			}
			if len(ing.Spec.TLS) != 0 {
				target.PodPort = strconv.Itoa(int(config.HTTPSPortExternal))
				target.URLs = domainsToURL(externalDomains, "https")
			} else {
				target.PodPort = strconv.Itoa(int(config.HTTPPortExternal))
				target.URLs = domainsToURL(externalDomains, scheme)
			}
			targets = append(targets, target)
		}

		target = status.ProbeTarget{
			PodIPs:  ips,
			PodPort: strconv.Itoa(int(config.HTTPPortInternal)),
			URLs:    domainsToURL(internalDomains, scheme),
		}

		targets = append(targets, target)

	}
	return targets, nil
}

func domainsToURL(domains []string, scheme string) []*url.URL {
	var urls []*url.URL

	for _, domain := range domains {
		url := &url.URL{
			Scheme: scheme,
			Host:   domain,
			Path:   "/",
		}
		urls = append(urls, url)
	}
	return urls
}

func getInternalDomains(rule v1alpha1.IngressRule, localDomainName string) []string {
	var res []string

	for _, host := range rule.Hosts {
		if strings.Contains(host, localDomainName) {
			res = append(res, host)
		}
	}

	return res
}

func getExternalDomains(rule v1alpha1.IngressRule, localDomainName string) []string {
	var res []string
	for _, host := range rule.Hosts {
		if !strings.Contains(host, localDomainName) {
			res = append(res, host)
		}
	}
	return res
}
