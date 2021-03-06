// Code generated by govpp binapi-generator DO NOT EDIT.
// Package stats represents the VPP binary API of the 'stats' VPP module.
// Generated from '/usr/share/vpp/api/stats.api.json'
package stats

import "git.fd.io/govpp.git/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xe8ea8b53

// IP4FibCounter represents the VPP binary API data type 'ip4_fib_counter'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 3:
//
//        ["ip4_fib_counter",
//            ["u32", "address"],
//            ["u8", "address_length"],
//            ["u64", "packets"],
//            ["u64", "bytes"],
//            {"crc" : "0xb2739495"}
//        ],
//
type IP4FibCounter struct {
	Address       uint32
	AddressLength uint8
	Packets       uint64
	Bytes         uint64
}

func (*IP4FibCounter) GetTypeName() string {
	return "ip4_fib_counter"
}
func (*IP4FibCounter) GetCrcString() string {
	return "b2739495"
}

// IP4NbrCounter represents the VPP binary API data type 'ip4_nbr_counter'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 10:
//
//        ["ip4_nbr_counter",
//            ["u32", "address"],
//            ["u8", "link_type"],
//            ["u64", "packets"],
//            ["u64", "bytes"],
//            {"crc" : "0x487e2e85"}
//        ],
//
type IP4NbrCounter struct {
	Address  uint32
	LinkType uint8
	Packets  uint64
	Bytes    uint64
}

func (*IP4NbrCounter) GetTypeName() string {
	return "ip4_nbr_counter"
}
func (*IP4NbrCounter) GetCrcString() string {
	return "487e2e85"
}

// IP6FibCounter represents the VPP binary API data type 'ip6_fib_counter'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 17:
//
//        ["ip6_fib_counter",
//            ["u64", "address", 2],
//            ["u8", "address_length"],
//            ["u64", "packets"],
//            ["u64", "bytes"],
//            {"crc" : "0xcf35769b"}
//        ],
//
type IP6FibCounter struct {
	Address       []uint64 `struc:"[2]uint64"`
	AddressLength uint8
	Packets       uint64
	Bytes         uint64
}

func (*IP6FibCounter) GetTypeName() string {
	return "ip6_fib_counter"
}
func (*IP6FibCounter) GetCrcString() string {
	return "cf35769b"
}

// IP6NbrCounter represents the VPP binary API data type 'ip6_nbr_counter'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 24:
//
//        ["ip6_nbr_counter",
//            ["u64", "address", 2],
//            ["u8", "link_type"],
//            ["u64", "packets"],
//            ["u64", "bytes"],
//            {"crc" : "0xefca741e"}
//        ]
//
type IP6NbrCounter struct {
	Address  []uint64 `struc:"[2]uint64"`
	LinkType uint8
	Packets  uint64
	Bytes    uint64
}

func (*IP6NbrCounter) GetTypeName() string {
	return "ip6_nbr_counter"
}
func (*IP6NbrCounter) GetCrcString() string {
	return "efca741e"
}

// WantStats represents the VPP binary API message 'want_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 33:
//
//        ["want_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0x4f2effb4"}
//        ],
//
type WantStats struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantStats) GetMessageName() string {
	return "want_stats"
}
func (*WantStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantStats) GetCrcString() string {
	return "4f2effb4"
}
func NewWantStats() api.Message {
	return &WantStats{}
}

// WantStatsReply represents the VPP binary API message 'want_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 41:
//
//        ["want_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xb36abf5f"}
//        ],
//
type WantStatsReply struct {
	Retval int32
}

func (*WantStatsReply) GetMessageName() string {
	return "want_stats_reply"
}
func (*WantStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantStatsReply) GetCrcString() string {
	return "b36abf5f"
}
func NewWantStatsReply() api.Message {
	return &WantStatsReply{}
}

// WantInterfaceSimpleStats represents the VPP binary API message 'want_interface_simple_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 47:
//
//        ["want_interface_simple_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0xbb4739ed"}
//        ],
//
type WantInterfaceSimpleStats struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantInterfaceSimpleStats) GetMessageName() string {
	return "want_interface_simple_stats"
}
func (*WantInterfaceSimpleStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantInterfaceSimpleStats) GetCrcString() string {
	return "bb4739ed"
}
func NewWantInterfaceSimpleStats() api.Message {
	return &WantInterfaceSimpleStats{}
}

// WantInterfaceSimpleStatsReply represents the VPP binary API message 'want_interface_simple_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 55:
//
//        ["want_interface_simple_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x5163615b"}
//        ],
//
type WantInterfaceSimpleStatsReply struct {
	Retval int32
}

func (*WantInterfaceSimpleStatsReply) GetMessageName() string {
	return "want_interface_simple_stats_reply"
}
func (*WantInterfaceSimpleStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantInterfaceSimpleStatsReply) GetCrcString() string {
	return "5163615b"
}
func NewWantInterfaceSimpleStatsReply() api.Message {
	return &WantInterfaceSimpleStatsReply{}
}

// WantPerInterfaceSimpleStats represents the VPP binary API message 'want_per_interface_simple_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 61:
//
//        ["want_per_interface_simple_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            ["u32", "num"],
//            ["u32", "sw_ifs", 0, "num"],
//            {"crc" : "0x0165e6f2"}
//        ],
//
type WantPerInterfaceSimpleStats struct {
	EnableDisable uint32
	Pid           uint32
	Num           uint32 `struc:"sizeof=SwIfs"`
	SwIfs         []uint32
}

func (*WantPerInterfaceSimpleStats) GetMessageName() string {
	return "want_per_interface_simple_stats"
}
func (*WantPerInterfaceSimpleStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantPerInterfaceSimpleStats) GetCrcString() string {
	return "0165e6f2"
}
func NewWantPerInterfaceSimpleStats() api.Message {
	return &WantPerInterfaceSimpleStats{}
}

// WantPerInterfaceSimpleStatsReply represents the VPP binary API message 'want_per_interface_simple_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 71:
//
//        ["want_per_interface_simple_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x720ee096"}
//        ],
//
type WantPerInterfaceSimpleStatsReply struct {
	Retval int32
}

func (*WantPerInterfaceSimpleStatsReply) GetMessageName() string {
	return "want_per_interface_simple_stats_reply"
}
func (*WantPerInterfaceSimpleStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantPerInterfaceSimpleStatsReply) GetCrcString() string {
	return "720ee096"
}
func NewWantPerInterfaceSimpleStatsReply() api.Message {
	return &WantPerInterfaceSimpleStatsReply{}
}

// WantInterfaceCombinedStats represents the VPP binary API message 'want_interface_combined_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 77:
//
//        ["want_interface_combined_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0xa98830b6"}
//        ],
//
type WantInterfaceCombinedStats struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantInterfaceCombinedStats) GetMessageName() string {
	return "want_interface_combined_stats"
}
func (*WantInterfaceCombinedStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantInterfaceCombinedStats) GetCrcString() string {
	return "a98830b6"
}
func NewWantInterfaceCombinedStats() api.Message {
	return &WantInterfaceCombinedStats{}
}

// WantInterfaceCombinedStatsReply represents the VPP binary API message 'want_interface_combined_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 85:
//
//        ["want_interface_combined_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xb260196d"}
//        ],
//
type WantInterfaceCombinedStatsReply struct {
	Retval int32
}

func (*WantInterfaceCombinedStatsReply) GetMessageName() string {
	return "want_interface_combined_stats_reply"
}
func (*WantInterfaceCombinedStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantInterfaceCombinedStatsReply) GetCrcString() string {
	return "b260196d"
}
func NewWantInterfaceCombinedStatsReply() api.Message {
	return &WantInterfaceCombinedStatsReply{}
}

// WantPerInterfaceCombinedStats represents the VPP binary API message 'want_per_interface_combined_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 91:
//
//        ["want_per_interface_combined_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            ["u32", "num"],
//            ["u32", "sw_ifs", 0, "num"],
//            {"crc" : "0xd914890f"}
//        ],
//
type WantPerInterfaceCombinedStats struct {
	EnableDisable uint32
	Pid           uint32
	Num           uint32 `struc:"sizeof=SwIfs"`
	SwIfs         []uint32
}

func (*WantPerInterfaceCombinedStats) GetMessageName() string {
	return "want_per_interface_combined_stats"
}
func (*WantPerInterfaceCombinedStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantPerInterfaceCombinedStats) GetCrcString() string {
	return "d914890f"
}
func NewWantPerInterfaceCombinedStats() api.Message {
	return &WantPerInterfaceCombinedStats{}
}

// WantPerInterfaceCombinedStatsReply represents the VPP binary API message 'want_per_interface_combined_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 101:
//
//        ["want_per_interface_combined_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x06158495"}
//        ],
//
type WantPerInterfaceCombinedStatsReply struct {
	Retval int32
}

func (*WantPerInterfaceCombinedStatsReply) GetMessageName() string {
	return "want_per_interface_combined_stats_reply"
}
func (*WantPerInterfaceCombinedStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantPerInterfaceCombinedStatsReply) GetCrcString() string {
	return "06158495"
}
func NewWantPerInterfaceCombinedStatsReply() api.Message {
	return &WantPerInterfaceCombinedStatsReply{}
}

// WantIP4FibStats represents the VPP binary API message 'want_ip4_fib_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 107:
//
//        ["want_ip4_fib_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0x6ce4937d"}
//        ],
//
type WantIP4FibStats struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantIP4FibStats) GetMessageName() string {
	return "want_ip4_fib_stats"
}
func (*WantIP4FibStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantIP4FibStats) GetCrcString() string {
	return "6ce4937d"
}
func NewWantIP4FibStats() api.Message {
	return &WantIP4FibStats{}
}

// WantIP4FibStatsReply represents the VPP binary API message 'want_ip4_fib_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 115:
//
//        ["want_ip4_fib_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xd337cb77"}
//        ],
//
type WantIP4FibStatsReply struct {
	Retval int32
}

func (*WantIP4FibStatsReply) GetMessageName() string {
	return "want_ip4_fib_stats_reply"
}
func (*WantIP4FibStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantIP4FibStatsReply) GetCrcString() string {
	return "d337cb77"
}
func NewWantIP4FibStatsReply() api.Message {
	return &WantIP4FibStatsReply{}
}

// WantIP6FibStats represents the VPP binary API message 'want_ip6_fib_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 121:
//
//        ["want_ip6_fib_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0x01697516"}
//        ],
//
type WantIP6FibStats struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantIP6FibStats) GetMessageName() string {
	return "want_ip6_fib_stats"
}
func (*WantIP6FibStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantIP6FibStats) GetCrcString() string {
	return "01697516"
}
func NewWantIP6FibStats() api.Message {
	return &WantIP6FibStats{}
}

// WantIP6FibStatsReply represents the VPP binary API message 'want_ip6_fib_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 129:
//
//        ["want_ip6_fib_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x531708ab"}
//        ],
//
type WantIP6FibStatsReply struct {
	Retval int32
}

func (*WantIP6FibStatsReply) GetMessageName() string {
	return "want_ip6_fib_stats_reply"
}
func (*WantIP6FibStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantIP6FibStatsReply) GetCrcString() string {
	return "531708ab"
}
func NewWantIP6FibStatsReply() api.Message {
	return &WantIP6FibStatsReply{}
}

// WantIP4NbrStats represents the VPP binary API message 'want_ip4_nbr_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 135:
//
//        ["want_ip4_nbr_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0x6bea26e1"}
//        ],
//
type WantIP4NbrStats struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantIP4NbrStats) GetMessageName() string {
	return "want_ip4_nbr_stats"
}
func (*WantIP4NbrStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantIP4NbrStats) GetCrcString() string {
	return "6bea26e1"
}
func NewWantIP4NbrStats() api.Message {
	return &WantIP4NbrStats{}
}

// WantIP4NbrStatsReply represents the VPP binary API message 'want_ip4_nbr_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 143:
//
//        ["want_ip4_nbr_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xe27d62cd"}
//        ],
//
type WantIP4NbrStatsReply struct {
	Retval int32
}

func (*WantIP4NbrStatsReply) GetMessageName() string {
	return "want_ip4_nbr_stats_reply"
}
func (*WantIP4NbrStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantIP4NbrStatsReply) GetCrcString() string {
	return "e27d62cd"
}
func NewWantIP4NbrStatsReply() api.Message {
	return &WantIP4NbrStatsReply{}
}

// WantIP6NbrStats represents the VPP binary API message 'want_ip6_nbr_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 149:
//
//        ["want_ip6_nbr_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0x0667c08a"}
//        ],
//
type WantIP6NbrStats struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantIP6NbrStats) GetMessageName() string {
	return "want_ip6_nbr_stats"
}
func (*WantIP6NbrStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantIP6NbrStats) GetCrcString() string {
	return "0667c08a"
}
func NewWantIP6NbrStats() api.Message {
	return &WantIP6NbrStats{}
}

// WantIP6NbrStatsReply represents the VPP binary API message 'want_ip6_nbr_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 157:
//
//        ["want_ip6_nbr_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x625da111"}
//        ],
//
type WantIP6NbrStatsReply struct {
	Retval int32
}

func (*WantIP6NbrStatsReply) GetMessageName() string {
	return "want_ip6_nbr_stats_reply"
}
func (*WantIP6NbrStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantIP6NbrStatsReply) GetCrcString() string {
	return "625da111"
}
func NewWantIP6NbrStatsReply() api.Message {
	return &WantIP6NbrStatsReply{}
}

// VnetIP4FibCounters represents the VPP binary API message 'vnet_ip4_fib_counters'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 163:
//
//        ["vnet_ip4_fib_counters",
//            ["u16", "_vl_msg_id"],
//            ["u32", "vrf_id"],
//            ["u32", "count"],
//            ["vl_api_ip4_fib_counter_t", "c", 0, "count"],
//            {"crc" : "0x1ab9d6c5"}
//        ],
//
type VnetIP4FibCounters struct {
	VrfID uint32
	Count uint32 `struc:"sizeof=C"`
	C     []IP4FibCounter
}

func (*VnetIP4FibCounters) GetMessageName() string {
	return "vnet_ip4_fib_counters"
}
func (*VnetIP4FibCounters) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*VnetIP4FibCounters) GetCrcString() string {
	return "1ab9d6c5"
}
func NewVnetIP4FibCounters() api.Message {
	return &VnetIP4FibCounters{}
}

// VnetIP4NbrCounters represents the VPP binary API message 'vnet_ip4_nbr_counters'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 170:
//
//        ["vnet_ip4_nbr_counters",
//            ["u16", "_vl_msg_id"],
//            ["u32", "count"],
//            ["u32", "sw_if_index"],
//            ["u8", "begin"],
//            ["vl_api_ip4_nbr_counter_t", "c", 0, "count"],
//            {"crc" : "0xfc2b5092"}
//        ],
//
type VnetIP4NbrCounters struct {
	Count     uint32 `struc:"sizeof=C"`
	SwIfIndex uint32
	Begin     uint8
	C         []IP4NbrCounter
}

func (*VnetIP4NbrCounters) GetMessageName() string {
	return "vnet_ip4_nbr_counters"
}
func (*VnetIP4NbrCounters) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*VnetIP4NbrCounters) GetCrcString() string {
	return "fc2b5092"
}
func NewVnetIP4NbrCounters() api.Message {
	return &VnetIP4NbrCounters{}
}

// VnetIP6FibCounters represents the VPP binary API message 'vnet_ip6_fib_counters'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 178:
//
//        ["vnet_ip6_fib_counters",
//            ["u16", "_vl_msg_id"],
//            ["u32", "vrf_id"],
//            ["u32", "count"],
//            ["vl_api_ip6_fib_counter_t", "c", 0, "count"],
//            {"crc" : "0x9ab453ae"}
//        ],
//
type VnetIP6FibCounters struct {
	VrfID uint32
	Count uint32 `struc:"sizeof=C"`
	C     []IP6FibCounter
}

func (*VnetIP6FibCounters) GetMessageName() string {
	return "vnet_ip6_fib_counters"
}
func (*VnetIP6FibCounters) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*VnetIP6FibCounters) GetCrcString() string {
	return "9ab453ae"
}
func NewVnetIP6FibCounters() api.Message {
	return &VnetIP6FibCounters{}
}

// VnetIP6NbrCounters represents the VPP binary API message 'vnet_ip6_nbr_counters'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 185:
//
//        ["vnet_ip6_nbr_counters",
//            ["u16", "_vl_msg_id"],
//            ["u32", "count"],
//            ["u32", "sw_if_index"],
//            ["u8", "begin"],
//            ["vl_api_ip6_nbr_counter_t", "c", 0, "count"],
//            {"crc" : "0x181b673f"}
//        ],
//
type VnetIP6NbrCounters struct {
	Count     uint32 `struc:"sizeof=C"`
	SwIfIndex uint32
	Begin     uint8
	C         []IP6NbrCounter
}

func (*VnetIP6NbrCounters) GetMessageName() string {
	return "vnet_ip6_nbr_counters"
}
func (*VnetIP6NbrCounters) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*VnetIP6NbrCounters) GetCrcString() string {
	return "181b673f"
}
func NewVnetIP6NbrCounters() api.Message {
	return &VnetIP6NbrCounters{}
}

// VnetGetSummaryStats represents the VPP binary API message 'vnet_get_summary_stats'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 193:
//
//        ["vnet_get_summary_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x16435c20"}
//        ],
//
type VnetGetSummaryStats struct {
}

func (*VnetGetSummaryStats) GetMessageName() string {
	return "vnet_get_summary_stats"
}
func (*VnetGetSummaryStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VnetGetSummaryStats) GetCrcString() string {
	return "16435c20"
}
func NewVnetGetSummaryStats() api.Message {
	return &VnetGetSummaryStats{}
}

// VnetGetSummaryStatsReply represents the VPP binary API message 'vnet_get_summary_stats_reply'.
// Generated from '/usr/share/vpp/api/stats.api.json', line 199:
//
//        ["vnet_get_summary_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u64", "total_pkts", 2],
//            ["u64", "total_bytes", 2],
//            ["f64", "vector_rate"],
//            {"crc" : "0x675ce280"}
//        ]
//
type VnetGetSummaryStatsReply struct {
	Retval     int32
	TotalPkts  []uint64 `struc:"[2]uint64"`
	TotalBytes []uint64 `struc:"[2]uint64"`
	VectorRate float64
}

func (*VnetGetSummaryStatsReply) GetMessageName() string {
	return "vnet_get_summary_stats_reply"
}
func (*VnetGetSummaryStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VnetGetSummaryStatsReply) GetCrcString() string {
	return "675ce280"
}
func NewVnetGetSummaryStatsReply() api.Message {
	return &VnetGetSummaryStatsReply{}
}
