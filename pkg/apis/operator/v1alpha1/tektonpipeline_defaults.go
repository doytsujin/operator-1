/*
Copyright 2021 The Tekton Authors

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

package v1alpha1

import (
	"context"

	"github.com/tektoncd/pipeline/pkg/apis/config"
	"knative.dev/pkg/ptr"
)

func (tp *TektonPipeline) SetDefaults(ctx context.Context) {
	tp.Spec.PipelineProperties.setDefaults()
}

func (p *PipelineProperties) setDefaults() {
	if p.DisableCredsInit == nil {
		p.DisableCredsInit = ptr.Bool(config.DefaultDisableCredsInit)
	}
	if p.AwaitSidecarReadiness == nil {
		p.AwaitSidecarReadiness = ptr.Bool(config.DefaultAwaitSidecarReadiness)
	}
	if p.RunningInEnvironmentWithInjectedSidecars == nil {
		p.RunningInEnvironmentWithInjectedSidecars = ptr.Bool(config.DefaultRunningInEnvWithInjectedSidecars)
	}
	if p.RequireGitSshSecretKnownHosts == nil {
		p.RequireGitSshSecretKnownHosts = ptr.Bool(config.DefaultRequireGitSSHSecretKnownHosts)
	}
	if p.EnableTektonOciBundles == nil {
		p.EnableTektonOciBundles = ptr.Bool(config.DefaultEnableTektonOciBundles)
	}
	if p.EnableCustomTasks == nil {
		p.EnableCustomTasks = ptr.Bool(config.DefaultEnableCustomTasks)
	}
	if p.SendCloudEventsForRuns == nil {
		p.SendCloudEventsForRuns = ptr.Bool(config.DefaultSendCloudEventsForRuns)
	}
	if p.EnableApiFields == "" {
		p.EnableApiFields = config.DefaultEnableAPIFields
	}
	if p.EmbeddedStatus == "" {
		p.EmbeddedStatus = config.DefaultEmbeddedStatus
	}

	// Deprecated: set to nil, remove in further release
	p.ScopeWhenExpressionsToTask = nil

	if p.MetricsPipelinerunDurationType == "" {
		p.MetricsPipelinerunDurationType = config.DefaultDurationPipelinerunType
	}
	if p.MetricsPipelinerunLevel == "" {
		p.MetricsPipelinerunLevel = config.DefaultPipelinerunLevel
	}
	if p.MetricsTaskrunDurationType == "" {
		p.MetricsTaskrunDurationType = config.DefaultDurationTaskrunType
	}
	if p.MetricsTaskrunLevel == "" {
		p.MetricsTaskrunLevel = config.DefaultTaskrunLevel
	}
}
