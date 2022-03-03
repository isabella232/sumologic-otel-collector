// Copyright 2021, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package telegrafreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
)

func createTestConfig() *Config {
	config := createDefaultConfig().(*Config)
	config.AgentConfig = `
[agent]
	interval = "2s"
	flush_interval = "3s"
[[inputs.mem]]
	`
	return config
}

func TestStartShutdown(t *testing.T) {
	ctx := context.Background()
	cfg := createTestConfig()
	receiver, err := createMetricsReceiver(ctx, componenttest.NewNopReceiverCreateSettings(), cfg, consumertest.NewNop())
	require.NoError(t, err)
	require.NoError(t, receiver.Start(ctx, componenttest.NewNopHost()))
	require.NoError(t, receiver.Shutdown(ctx))
}