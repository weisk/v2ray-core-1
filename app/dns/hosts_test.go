package dns_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	. "github.com/v2fly/v2ray-core/v4/app/dns"
	"github.com/v2fly/v2ray-core/v4/common"
	"github.com/v2fly/v2ray-core/v4/common/net"
)

func TestStaticHosts(t *testing.T) {
	pb := []*Config_HostMapping{
		{
			Type:   DomainMatchingType_Full,
			Domain: "v2fly.org",
			Ip: [][]byte{
				{1, 1, 1, 1},
			},
		},
		{
			Type:   DomainMatchingType_Subdomain,
			Domain: "v2ray.cn",
			Ip: [][]byte{
				{2, 2, 2, 2},
			},
		},
		{
			Type:   DomainMatchingType_Subdomain,
			Domain: "baidu.com",
			Ip: [][]byte{
				{127, 0, 0, 1},
			},
		},
	}

	hosts, err := NewStaticHosts(pb, nil)
	common.Must(err)

	{
		ips := hosts.Lookup("v2fly.org", IPOption{
			IPv4Enable: true,
			IPv6Enable: true,
		})
		if len(ips) != 1 {
			t.Error("expect 1 IP, but got ", len(ips))
		}
		if diff := cmp.Diff([]byte(ips[0].IP()), []byte{1, 1, 1, 1}); diff != "" {
			t.Error(diff)
		}
	}

	{
		ips := hosts.Lookup("www.v2ray.cn", IPOption{
			IPv4Enable: true,
			IPv6Enable: true,
		})
		if len(ips) != 1 {
			t.Error("expect 1 IP, but got ", len(ips))
		}
		if diff := cmp.Diff([]byte(ips[0].IP()), []byte{2, 2, 2, 2}); diff != "" {
			t.Error(diff)
		}
	}

	{
		ips := hosts.Lookup("baidu.com", IPOption{
			IPv4Enable: false,
			IPv6Enable: true,
		})
		if len(ips) != 1 {
			t.Error("expect 1 IP, but got ", len(ips))
		}
		if diff := cmp.Diff([]byte(ips[0].IP()), []byte(net.LocalHostIPv6.IP())); diff != "" {
			t.Error(diff)
		}
	}
}
