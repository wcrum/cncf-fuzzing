// Copyright 2022 ADA Logics Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package sensor

import (
	fuzz "github.com/AdaLogics/go-fuzz-headers"
	eventbusv1alpha1 "github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1"
	"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	"sigs.k8s.io/yaml"
)

func FuzzValidateSensor(data []byte) int {
	f := fuzz.NewConsumer(data)
	eventBus := &eventbusv1alpha1.EventBus{}
	err := f.GenerateStruct(eventBus)
	if err != nil {
		return 0
	}
	content, err := f.GetBytes()
	if err != nil {
		return 0
	}
	sensor := &v1alpha1.Sensor{}
	err = yaml.Unmarshal(content, &sensor)
	if err != nil {
		return 0
	}
	if sensor == nil {
		return 0
	}
	ValidateSensor(sensor, eventBus)
	return 1
}
