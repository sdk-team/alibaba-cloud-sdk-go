package bss

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "bss.aliyuncs.com",
			"cn-beijing":     "bss.aliyuncs.com",
			"ap-south-1":     "bss.aliyuncs.com",
			"eu-west-1":      "bss.aliyuncs.com",
			"ap-northeast-1": "bss.aliyuncs.com",
			"me-east-1":      "bss.aliyuncs.com",
			"cn-chengdu":     "bss.aliyuncs.com",
			"cn-qingdao":     "bss.aliyuncs.com",
			"cn-shanghai":    "bss.aliyuncs.com",
			"cn-hongkong":    "bss.aliyuncs.com",
			"ap-southeast-1": "bss.aliyuncs.com",
			"ap-southeast-2": "bss.aliyuncs.com",
			"ap-southeast-3": "bss.aliyuncs.com",
			"eu-central-1":   "bss.aliyuncs.com",
			"cn-huhehaote":   "bss.aliyuncs.com",
			"ap-southeast-5": "bss.aliyuncs.com",
			"us-east-1":      "bss.aliyuncs.com",
			"cn-zhangjiakou": "bss.aliyuncs.com",
			"us-west-1":      "bss.aliyuncs.com",
			"cn-hangzhou":    "bss.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
