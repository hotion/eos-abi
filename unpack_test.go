package eosabi_test

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Jeiwan/eos-abi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		abi      string
		t        string
		hexData  string
		expected string
	}{
		{
			"eosio",
			"newaccount",
			"0000000000ea30550000735802ea305501000000010003350529efca8c607421e95846cc2a3d2efaa8454018deb75a204e27acf29ee5cc0100000001000000010003350529efca8c607421e95846cc2a3d2efaa8454018deb75a204e27acf29ee5cc01000000",
			`
			{
				"creator" : "eosio",
				"name" : "eosio.msig",
				"owner" : {
					"threshold" : 1,
					"keys" : [
						{
							"key" : "EOS7EarnUhcyYqmdnPon8rm7mBCTnBoot6o7fE2WzjvEX2TdggbL3",
							"weight" : 1
						}
					],
					"accounts" : [],
					"waits" : []
				},
				"active" : {
					"threshold" : 1,
					"keys" : [
						{
							"key" : "EOS7EarnUhcyYqmdnPon8rm7mBCTnBoot6o7fE2WzjvEX2TdggbL3",
							"weight" : 1
						}
					],
					"accounts" : [],
					"waits" : []
				}
			}
			`,
		},
		{
			"eosio",
			"setcode",
			"0000735802ea305500000a0061736d010000000198",
			`
			{
				"account" : "eosio.msig",
				"vmtype" : 0,
				"vmversion" : 0,
				"code" : "AGFzbQEAAAABmA=="
			}
			`,
		},
		{
			"eosio",
			"setabi",
			"9098ba5303ea3055d5010e656f73696f3a3a6162692f312e300110657468657265756d5f6164647265737306737472696e6702076164647265737300030269640675696e74363410657468657265756d5f6164647265737310657468657265756d5f616464726573730762616c616e636505617373657403616464000210657468657265756d5f6164647265737310657468657265756d5f616464726573730762616c616e63650561737365740100000000000052320361646400010000c00a637553320369363401026964010675696e7436340761646472657373000000",
			`
			{
				"account" : "eosio.unregd",
				"abi" : "DmVvc2lvOjphYmkvMS4wARBldGhlcmV1bV9hZGRyZXNzBnN0cmluZwIHYWRkcmVzcwADAmlkBnVpbnQ2NBBldGhlcmV1bV9hZGRyZXNzEGV0aGVyZXVtX2FkZHJlc3MHYmFsYW5jZQVhc3NldANhZGQAAhBldGhlcmV1bV9hZGRyZXNzEGV0aGVyZXVtX2FkZHJlc3MHYmFsYW5jZQVhc3NldAEAAAAAAABSMgNhZGQAAQAAwApjdVMyA2k2NAECaWQBBnVpbnQ2NAdhZGRyZXNzAAAA"
			}
			`,
		},
		{
			"eosio",
			"buyrambytes",
			"0000000000ea3055000000000000403800200000",
			`
			{
				"payer" : "eosio",
				"receiver" : "b1",
				"bytes" : 8192
			}
			`,
		},
		{
			"eosio",
			"buyram",
			"90d8a4914db5b48b90d8a4914db5b48b5cd542000000000004454f5300000000",
			`
			{
				"payer": "liufenglongd",
				"receiver": "liufenglongd",
				"quant": "437.9996 EOS"
			}`,
		},
		{
			"eosio",
			"delegatebw",
			"0000000000ea30550000000000004038e2c4516a7400000004454f5300000000e2c4516a7400000004454f530000000001",
			`
			{
				"from" : "eosio",
				"receiver" : "b1",
				"stake_net_quantity" : "49999995.0050 EOS",
				"stake_cpu_quantity" : "49999995.0050 EOS",
				"transfer" : 1
			}
			`,
		},
		{
			"eosio",
			"setparams",
			"0000100000000000e8030000000008000c000000f4010000140000006400000000e1f505e80300009be0f50564000000100e00005802000080533b000010000004000600",
			`
			{
				"params" : {
					"max_block_net_usage" : 1048576,
					"target_block_net_usage_pct" : 1000,
					"max_transaction_net_usage" : 524288,
					"base_per_transaction_net_usage" : 12,
					"net_usage_leeway" : 500,
					"context_free_discount_net_usage_num" : 20,
					"context_free_discount_net_usage_den" : 100,
					"max_block_cpu_usage" : 100000000,
					"target_block_cpu_usage_pct" : 1000,
					"max_transaction_cpu_usage" : 99999899,
					"min_transaction_cpu_usage" : 100,
					"max_transaction_lifetime" : 3600,
					"deferred_trx_expiration_window" : 600,
					"max_transaction_delay" : 3888000,
					"max_inline_action_size" : 4096,
					"max_inline_action_depth" : 4,
					"max_authority_depth" : 6
				}
			}
			`,
		},
		{
			"eosio.token",
			"transfer",
			"0000000000ea30550000000000004038a08601000000000004454f530000000098014e6576657220646f7562742074686174206120736d616c6c2067726f7570206f662074686f7567687466756c2c20636f6d6d697474656420636974697a656e732063616e206368616e67652074686520776f726c643b20696e646565642c206974277320746865206f6e6c79207468696e672074686174206576657220686173202d20656f7361636b6e6f776c6564676d656e74732e696f",
			`
			{
				"from" : "eosio",
				"to" : "b1",
				"quantity" : "10.0000 EOS",
				"memo" : "Never doubt that a small group of thoughtful, committed citizens can change the world; indeed, it's the only thing that ever has - eosacknowledgments.io"
			}
			`,
		},
		{
			"eosio.token",
			"transfer",
			"0000000000ea3055a0d492e602ea3055050000000000000004454f53000000000772616d20666565",
			`
			{
				"from" : "eosio",
				"to" : "eosio.ramfee",
				"quantity" : "0.0005 EOS",
				"memo" : "ram fee"
			}
			`,
		},
		{
			"eosio",
			"voteproducer",
			"a09864ff4d94bd62000000000000000001e0b3bbb4656d3055",
			`
			{
				"voter": "geytcnjzgmge",
				"proxy": "",
				"producers": [
					"eosauthority"
				]
			}
			`,
		},
		{
			"eosio.token",
			"issue",
			"0000000000ea305500a0724e1809000004454f530000000010696e697469616c2069737375616e6365",
			`
			{
				"to": "eosio",
				"quantity": "1000000000.0000 EOS",
				"memo": "initial issuance"
			}
			`,
		},
		{
			"eosio.token",
			"create",
			"0000000000ea305500407a10f35a000004454f5300000000",
			`
			{
				"issuer": "eosio",
				"maximum_supply": "10000000000.0000 EOS"
			}
			`,
		},
		{
			"eosio",
			"setprods",
			"0100118d073baca6620003e1e431dead6379a8b9ac4e711ed959e0c224fefe2528a5dd01b6a74576ed33a6",
			`
			{
				"schedule":[
					{
						"block_signing_key": "EOS8Yid3mE5bwWMvGGKYEDxFRGHostu5xCzFanyJP1UdgZ5mpPdwZ",
						"producer_name": "genesisblock"
					}
				]
			}
			`,
		},
		{
			"eosio.msig",
			"propose",
			"1098c35d4db7b23b000000008090b1ca031098c35d4db7b23b00000000a8ed32322098c35d4db7b23b00000000a8ed32325098c35d4db7b23b00000000a8ed3232fd34295b000000000000000000000100a6823403ea3055000000572d3ccdcd011098c35d4db7b23b00000000a8ed32322c1002475d4db7b23b2002475d4db7b23b102700000000000004454f53000000000b53696d706c65207465737400",
			`
			{
				"proposer": "bitfinexsig1",
				"proposal_name": "test1",
				"requested": [
				  {
					"actor": "bitfinexsig1",
					"permission": "active"
				  },
				  {
					"actor": "bitfinexsig2",
					"permission": "active"
				  },
				  {
					"actor": "bitfinexsig5",
					"permission": "active"
				  }
				],
				"trx": {
				  "expiration": "2018-06-19T16:53:17Z",
				  "ref_block_num": 0,
				  "ref_block_prefix": 0,
				  "max_net_usage_words": 0,
				  "max_cpu_usage_ms": 0,
				  "delay_sec": 0,
				  "context_free_actions": [],
				  "actions": [
					{
					  "account": "eosio.token",
					  "name": "transfer",
					  "authorization": [
						{
						  "actor": "bitfinexsig1",
						  "permission": "active"
						}
					  ],
					  "data": "EAJHXU23sjsgAkddTbeyOxAnAAAAAAAABEVPUwAAAAALU2ltcGxlIHRlc3Q="
					}
				  ],
				  "transaction_extensions": []
				}
			  }
			`,
		},
	}

	abis := make(map[string][]byte)
	eosio, _ := ioutil.ReadFile("fixtures/eosio.json")
	abis["eosio"] = eosio
	eostoken, _ := ioutil.ReadFile("fixtures/eosiotoken.json")
	abis["eosio.token"] = eostoken
	msig, _ := ioutil.ReadFile("fixtures/eosio.msig.json")
	abis["eosio.msig"] = msig

	for _, test := range tests {
		t.Run(test.t, func(tt *testing.T) {
			data, err := hex.DecodeString(test.hexData)
			require.Nil(tt, err)

			abi := abis[test.abi]
			r, err := eosabi.UnpackAction(abi, test.t, data)
			require.Nil(tt, err)
			unpacked := r.(map[string]interface{})

			actual, err := json.Marshal(unpacked)
			assert.Nil(tt, err)

			assert.JSONEq(tt, test.expected, string(actual))
		})
	}
}
