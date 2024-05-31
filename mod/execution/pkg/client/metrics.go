// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package client

import (
	"time"

	"github.com/berachain/beacon-kit/mod/log"
)

// clientMetrics is a struct that contains metrics for the engine.
type clientMetrics struct {
	// TelemetrySink is the sink for the metrics.
	sink TelemetrySink
	// logger is the logger for the engineMetrics.
	logger log.Logger[any]
}

// newClientMetrics creates a new engineMetrics.
func newClientMetrics(
	sink TelemetrySink,
	logger log.Logger[any],
) *clientMetrics {
	return &clientMetrics{
		sink:   sink,
		logger: logger,
	}
}

// measureForkchoiceUpdateDuration measures the duration of the forkchoice
// update.
func (cm *clientMetrics) measureForkchoiceUpdateDuration(startTime time.Time) {
	// TODO: Add Labels.
	cm.sink.MeasureSince(
		"beacon_kit.execution.client.forkchoice_update_duration",
		startTime,
	)
}

// measureNewPayloadDuration measures the duration of the new payload.
func (cm *clientMetrics) measureNewPayloadDuration(startTime time.Time) {
	// TODO: Add Labels.
	cm.sink.MeasureSince(
		"beacon_kit.execution.client.new_payload_duration",
		startTime,
	)
}

// measureGetPayloadDuration measures the duration of the get payload.
func (cm *clientMetrics) measureGetPayloadDuration(startTime time.Time) {
	// TODO: Add Labels.
	cm.sink.MeasureSince(
		"beacon_kit.execution.client.get_payload_duration",
		startTime,
	)
}

// incrementForkchoiceUpdateTimeout increments the timeout counter
// for forkchoice update.
func (cm *clientMetrics) incrementForkchoiceUpdateTimeout() {
	cm.incrementTimeoutCounter(
		"beacon_kit.execution.client.forkchoice_update_duration")
}

// incrementNewPayloadTimeout increments the timeout counter for
// new payload.
func (cm *clientMetrics) incrementNewPayloadTimeout() {
	cm.incrementTimeoutCounter(
		"beacon_kit.execution.client.new_payload_duration")
}

// incrementGetPayloadTimeout increments the timeout counter for
// get payload.
func (cm *clientMetrics) incrementGetPayloadTimeout() {
	cm.incrementTimeoutCounter(
		"beacon_kit.execution.client.get_payload_duration")
}

// incrementTimeoutCounter increments the timeout counter for
// the given metric.
func (cm *clientMetrics) incrementTimeoutCounter(metricName string) {
	cm.sink.IncrementCounter(metricName + "_timeout")
}
