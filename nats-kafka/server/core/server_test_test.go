/*
 * Copyright 2019 The NATS Authors
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"testing"

	"github.com/nats-io/nats-kafka/server/conf"
	"github.com/stretchr/testify/require"
)

func TestFullTestEnvironment(t *testing.T) {
	connect := []conf.ConnectorConfig{}
	tbs, err := StartTestEnvironment(connect)
	require.NoError(t, err)
	tbs.Close()
}

func TestFullTLSTestEnvironment(t *testing.T) {
	connect := []conf.ConnectorConfig{}
	tbs, err := StartTLSTestEnvironment(connect)
	require.NoError(t, err)
	tbs.Close()
}
