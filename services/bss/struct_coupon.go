package bss

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

// Coupon is a nested struct in bss response
type Coupon struct {
	CouponTemplateId int64                            `json:"CouponTemplateId" xml:"CouponTemplateId"`
	TotalAmount      string                           `json:"TotalAmount" xml:"TotalAmount"`
	BalanceAmount    string                           `json:"BalanceAmount" xml:"BalanceAmount"`
	FrozenAmount     string                           `json:"FrozenAmount" xml:"FrozenAmount"`
	ExpiredAmount    string                           `json:"ExpiredAmount" xml:"ExpiredAmount"`
	DeliveryTime     string                           `json:"DeliveryTime" xml:"DeliveryTime"`
	ExpiredTime      string                           `json:"ExpiredTime" xml:"ExpiredTime"`
	CouponNumber     string                           `json:"CouponNumber" xml:"CouponNumber"`
	Status           string                           `json:"Status" xml:"Status"`
	Description      string                           `json:"Description" xml:"Description"`
	CreationTime     string                           `json:"CreationTime" xml:"CreationTime"`
	ModificationTime string                           `json:"ModificationTime" xml:"ModificationTime"`
	PriceLimit       string                           `json:"PriceLimit" xml:"PriceLimit"`
	Application      string                           `json:"Application" xml:"Application"`
	ProductCodes     ProductCodesInDescribeCouponList `json:"ProductCodes" xml:"ProductCodes"`
	TradeTypes       TradeTypesInDescribeCouponList   `json:"TradeTypes" xml:"TradeTypes"`
}
