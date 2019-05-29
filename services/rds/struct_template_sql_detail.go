
package rds

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

// TemplateSqlDetail is a nested struct in rds response
type TemplateSqlDetail struct {
            DBName     string `json:"DBName" xml:"DBName"`
            AccountName     string `json:"AccountName" xml:"AccountName"`
            HostAddress     string `json:"HostAddress" xml:"HostAddress"`
            SQLText     string `json:"SQLText" xml:"SQLText"`
            Consume     int64 `json:"Consume" xml:"Consume"`
            ScanRows     int64 `json:"ScanRows" xml:"ScanRows"`
            UpdateRows     int64 `json:"UpdateRows" xml:"UpdateRows"`
            State     string `json:"State" xml:"State"`
            ExecuteTime     string `json:"ExecuteTime" xml:"ExecuteTime"`
            ThreadID     int `json:"ThreadID" xml:"ThreadID"`
            SqlType     string `json:"SqlType" xml:"SqlType"`
}
