/*
Copyright © 2021-2022 Infinite Devices GmbH

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
package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpretHeaderFlags(t *testing.T) {
	input := byte(11)
	hdr, err := interpretPublishHeaderFlags(input)
	assert.NoError(t, err)
	assert.True(t, hdr.Dup)
	assert.True(t, hdr.Retain)
	assert.Equal(t, QoSLevelAtLeastOnce, hdr.QoS, "Expected at least once")
}
