package types

import "github.com/insomniacslk/dhcp/dhcpv4"

type OptionTagName string

const (
	TagNameSubnetMask = "subnet_mask"
	TagNameRouter     = "router"
	TagNameTimeServer = "time_server"
	TagNameNameServer = "name_server"
	TagNameDomainName = "domain_name"
	TagNameBootfile   = "bootfile"
	TagNameTFTPserver = "tftp_server"
)

// https://datatracker.ietf.org/doc/html/rfc2131
// https://datatracker.ietf.org/doc/html/rfc2132
var TagMap map[OptionTagName]uint8 = map[OptionTagName]uint8{
	TagNameSubnetMask: dhcpv4.OptionSubnetMask.Code(),
	TagNameRouter:     dhcpv4.OptionRouter.Code(),
	TagNameTimeServer: dhcpv4.OptionTimeServer.Code(),
	TagNameNameServer: dhcpv4.OptionNameServer.Code(),
	TagNameDomainName: dhcpv4.OptionDomainName.Code(),
	TagNameBootfile:   dhcpv4.OptionBootfileName.Code(),
	TagNameTFTPserver: dhcpv4.OptionTFTPServerName.Code(),
}

var IPTags = map[uint8]bool{
	dhcpv4.OptionRouter.Code():           true,
	dhcpv4.OptionSubnetMask.Code():       true,
	dhcpv4.OptionNameServer.Code():       true,
	dhcpv4.OptionDomainNameServer.Code(): true,
	dhcpv4.OptionTimeServer.Code():       true,
}

type DHCPOption struct {
	Tag     *uint8   `json:"tag"`
	TagName string   `json:"tagName"`
	Value   *string  `json:"value"`
	Value64 []string `json:"value64"`
}

func OptionValue(input string) *string {
	return &input
}
