package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InstanceMonitorDataInDescribeInstanceMonitorData is a nested struct in ecs response
type InstanceMonitorDataInDescribeInstanceMonitorData struct {
	InstanceId                   string  `json:"InstanceId" xml:"InstanceId"`
	CPU                          int     `json:"CPU" xml:"CPU"`
	IntranetRX                   int     `json:"IntranetRX" xml:"IntranetRX"`
	IntranetTX                   int     `json:"IntranetTX" xml:"IntranetTX"`
	IntranetBandwidth            int     `json:"IntranetBandwidth" xml:"IntranetBandwidth"`
	InternetRX                   int     `json:"InternetRX" xml:"InternetRX"`
	InternetTX                   int     `json:"InternetTX" xml:"InternetTX"`
	InternetBandwidth            int     `json:"InternetBandwidth" xml:"InternetBandwidth"`
	IOPSRead                     int     `json:"IOPSRead" xml:"IOPSRead"`
	IOPSWrite                    int     `json:"IOPSWrite" xml:"IOPSWrite"`
	BPSRead                      int     `json:"BPSRead" xml:"BPSRead"`
	BPSWrite                     int     `json:"BPSWrite" xml:"BPSWrite"`
	CPUCreditUsage               float64 `json:"CPUCreditUsage" xml:"CPUCreditUsage"`
	CPUCreditBalance             float64 `json:"CPUCreditBalance" xml:"CPUCreditBalance"`
	CPUAdvanceCreditBalance      float64 `json:"CPUAdvanceCreditBalance" xml:"CPUAdvanceCreditBalance"`
	CPUNotpaidSurplusCreditUsage float64 `json:"CPUNotpaidSurplusCreditUsage" xml:"CPUNotpaidSurplusCreditUsage"`
	TimeStamp                    string  `json:"TimeStamp" xml:"TimeStamp"`
}
