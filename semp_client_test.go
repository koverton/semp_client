package semp_client

import (
	"fmt"
	"testing"
)

const testHost  = "192.168.56.103"
const basePath  = "http://" + testHost + ":8080/SEMP/v2/config"
const adminUser = "admin"
const adminPass = "admin"

func TestSempClient(t *testing.T) {
	cfg := NewConfiguration()
	//cfg.Host = testHost
	cfg.Username = adminUser
	cfg.Password = adminPass
	cfg.BasePath = basePath

	vpnApi := MsgVpnApi{cfg}
	vpn, _, err := vpnApi.GetMsgVpn("default", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Here's my VPN max-spool: %d\n", vpn.Data.MaxMsgSpoolUsage)

	 queue := MsgVpnQueue{ QueueName:"go_client_test", MsgVpnName:"default"}
	 queueApi := QueueApi{cfg}
	 q, _, err := queueApi.CreateMsgVpnQueue("default", queue, nil)
	 if err != nil {
		panic(err)
	 }
	 fmt.Printf("Here's my Queue max-spool: %d\n", q.Data.MaxMsgSpoolUsage)

	 m, _, err := queueApi.DeleteMsgVpnQueue("default", "go_client_test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleting queue, response code: %d\n", m.Meta.ResponseCode)
}

