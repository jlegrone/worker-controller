// Unless explicitly stated otherwise all files in this repository are licensed under the MIT License.
//
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2024 Datadog, Inc.

package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"

	temporaliov1alpha1 "github.com/DataDog/temporal-worker-controller/api/v1alpha1"
)

func TestFindHighestPriorityStatus(t *testing.T) {
	type testCase struct {
		statuses []temporaliov1alpha1.VersionStatus
		expected temporaliov1alpha1.VersionStatus
	}

	for name, tc := range map[string]testCase{
		"empty": {},
		"single status": {
			statuses: []temporaliov1alpha1.VersionStatus{temporaliov1alpha1.VersionStatusCurrent},
			expected: temporaliov1alpha1.VersionStatusCurrent,
		},
		"multiple statuses": {
			statuses: []temporaliov1alpha1.VersionStatus{
				temporaliov1alpha1.VersionStatusNotRegistered,
				temporaliov1alpha1.VersionStatusDraining,
				temporaliov1alpha1.VersionStatusCurrent,
				temporaliov1alpha1.VersionStatusDrained,
				temporaliov1alpha1.VersionStatusRamping,
			},
			expected: temporaliov1alpha1.VersionStatusCurrent,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := findHighestPriorityStatus(tc.statuses)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
