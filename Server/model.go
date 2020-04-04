package Server

type GatewayStruct struct {
	Name         string   `json:"name"`
	Ip_addresses []string `json:"ip_addresses"`
	Prefix       string   `json:"prefix"`
}

type CreateGatewayStruct struct {
	Name         string   `json:"name"`
	IP_addresses []string `json:"ip_addressess"`
}

type GatewayRouteStruct struct {
	Prefix     string `json:"prefix"`
	Gateway_id string `json:"gateway_id"`
}
