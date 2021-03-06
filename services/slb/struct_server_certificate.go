package slb

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

// ServerCertificate is a nested struct in slb response
type ServerCertificate struct {
	ServerCertificateId     string             `json:"ServerCertificateId" xml:"ServerCertificateId"`
	Fingerprint             string             `json:"Fingerprint" xml:"Fingerprint"`
	ServerCertificateName   string             `json:"ServerCertificateName" xml:"ServerCertificateName"`
	RegionId                string             `json:"RegionId" xml:"RegionId"`
	RegionIdAlias           string             `json:"RegionIdAlias" xml:"RegionIdAlias"`
	AliCloudCertificateId   string             `json:"AliCloudCertificateId" xml:"AliCloudCertificateId"`
	AliCloudCertificateName string             `json:"AliCloudCertificateName" xml:"AliCloudCertificateName"`
	IsAliCloudCertificate   int                `json:"IsAliCloudCertificate" xml:"IsAliCloudCertificate"`
	ExpireTime              string             `json:"ExpireTime" xml:"ExpireTime"`
	ExpireTimeStamp         int64              `json:"ExpireTimeStamp" xml:"ExpireTimeStamp"`
	Domain                  string             `json:"Domain" xml:"Domain"`
	AlternativeDomains      AlternativeDomains `json:"AlternativeDomains" xml:"AlternativeDomains"`
}
