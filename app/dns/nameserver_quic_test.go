package dns_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	. "github.com/v2fly/v2ray-core/v4/app/dns"
	"github.com/v2fly/v2ray-core/v4/common"
	"github.com/v2fly/v2ray-core/v4/common/net"
)

func TestQUICNameServer(t *testing.T) {
	url, err := url.Parse("quic://dns.adguard.com")
	common.Must(err)
	s, err := NewQUICNameServer(url)
	common.Must(err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	ips, err := s.QueryIP(ctx, "google.com", net.IP(nil), IPOption{
		IPv4Enable: true,
		IPv6Enable: true,
	})
	cancel()
	common.Must(err)
	if len(ips) == 0 {
		t.Error("expect some ips, but got 0")
	}
}
