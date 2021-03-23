package celltype

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
	"time"
)

/**
 * Copyright (C), 2019-2020
 * FileName: data_test
 * Author:   LinGuanHong
 * Date:     2020/12/20 2:57 下午
 * Description:
 */

func Test_ExpiredAtFromOutputData(t *testing.T) {
	dataHex := "c4de24c38f1a22e65b9a1a24aaae7d4db37e7ae138e9d44651d76f1d179f95e8ee06f79afc0af40e7198faf1611a8fa5324263b3f2dd3b620000000062616161616161612e626974"
	dataBys, _ := hex.DecodeString(dataHex)
	expired, e := ExpiredAtFromOutputData(dataBys)
	if e != nil {
		panic(e)
	}
	t.Log(expired)
}

func Test_DefaultAccountCellDataBytes(t *testing.T) {
	id := DasAccountIdFromBytes([]byte("123"))
	nextId := DasAccountIdFromBytes([]byte("456"))
	bys := DefaultAccountCellDataBytes(id, nextId)
	t.Log(bys)
}

func Test_DasAccountIdFromBytes(t *testing.T) {
	id := DasAccountIdFromBytes([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1})
	fmt.Println(id.HexStr())
}

func Test_EchoTypeId(t *testing.T) {
	t.Log(hexToArgsBytes("0x"))
	bys, _ := json.Marshal(DasTimeCellScript.Out.Script())
	t.Log(string(bys))
	// t.Log(DasAccountCellScript.Out)
	// account_cell : 0x274775e475c1252b5333c20e1512b7b1296c4c5b52a25aa2ebd6e41f5894c41f
	// // 0x9878b226df9465c215fd3c94dc9f9bf6648d5bea48a24579cf83274fe13801d2
	// t.Log(DasWalletCellScript.Out)
	// t.Log(DasTimeCellScript.Out.TypeId())
}

func Test_InitSystemScript(t *testing.T) {
	fmt.Println(DasProposeCellScript.Dep.ToDepCell().OutPoint.TxHash.String())
	SetSystemCodeScriptOutPoint(DasProposeCellScript.Out.CodeHash, types.OutPoint{
		TxHash: types.HexToHash("111"),
	})
	fmt.Println(SystemCodeScriptMap[DasProposeCellScript.Out.CodeHash].Dep.TxHash.String())
	fmt.Println(DasProposeCellScript.Dep.ToDepCell().OutPoint.TxHash.String())
	bys, err := SystemCodeScriptBytes()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hex.EncodeToString(bys))
	}
}

// func Test_CodeScriptFromBys(t *testing.T) {
// 	hexStr := "7b226163636f756e745f63656c6c223a7b22646570223a7b2274785f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c2274785f696e646578223a302c226465705f74797065223a22636f6465227d2c226f7574223a7b22636f64655f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c22636f64655f686173685f74797065223a2274797065222c2261726773223a6e756c6c7d7d2c226170706c795f72656769737465725f63656c6c223a7b22646570223a7b2274785f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c2274785f696e646578223a302c226465705f74797065223a22636f6465227d2c226f7574223a7b22636f64655f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c22636f64655f686173685f74797065223a2274797065222c2261726773223a6e756c6c7d7d2c2262696464696e675f63656c6c223a7b22646570223a7b2274785f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c2274785f696e646578223a302c226465705f74797065223a22636f6465227d2c226f7574223a7b22636f64655f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c22636f64655f686173685f74797065223a2274797065222c2261726773223a6e756c6c7d7d2c226f6e5f73616c655f63656c6c223a7b22646570223a7b2274785f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c2274785f696e646578223a302c226465705f74797065223a22636f6465227d2c226f7574223a7b22636f64655f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c22636f64655f686173685f74797065223a2274797065222c2261726773223a6e756c6c7d7d2c227072656163636f756e745f63656c6c223a7b22646570223a7b2274785f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c2274785f696e646578223a302c226465705f74797065223a22636f6465227d2c226f7574223a7b22636f64655f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c22636f64655f686173685f74797065223a2274797065222c2261726773223a6e756c6c7d7d2c2270726f706f73655f63656c6c223a7b22646570223a7b2274785f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c2274785f696e646578223a302c226465705f74797065223a22227d2c226f7574223a7b22636f64655f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303061626363222c22636f64655f686173685f74797065223a22222c2261726773223a6e756c6c7d7d2c227265665f63656c6c223a7b22646570223a7b2274785f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c2274785f696e646578223a302c226465705f74797065223a22636f6465227d2c226f7574223a7b22636f64655f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c22636f64655f686173685f74797065223a2274797065222c2261726773223a6e756c6c7d7d2c2277616c6c65745f63656c6c223a7b22646570223a7b2274785f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c2274785f696e646578223a302c226465705f74797065223a22636f6465227d2c226f7574223a7b22636f64655f68617368223a22307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c22636f64655f686173685f74797065223a2274797065222c2261726773223a6e756c6c7d7d7d"
// 	bys, err := hex.DecodeString(hexStr)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if err = SystemCodeScriptFromBytes(bys); err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println(string(bys))
// }

func Test_Ticker(t *testing.T) {
	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-ticker.C:
				fmt.Println("1111")
			default:
				time.Sleep(time.Second)
				fmt.Println("2")
			}
		}
	}()
	select {}
}

func Test_AccountCharLen(t *testing.T) {
	// accountId 包含 bit
	// 取价格，不需要
	//
	fmt.Println(len([]rune("xx🌹你")))
	fmt.Println([]byte("🌹"))
	/**
	[
		{
			emoji
			[]byte("🌹")
		},
		{
			en
			[]byte("a")
		},
		{
			zh
			[]byte("你")
		}
	]
	*/
}

func Test_PriceConfigs(t *testing.T) {
	getItem := func() *PriceConfig {
		p1 := NewPriceConfigBuilder().Length(GoUint8ToMoleculeU8(1)).Build()
		p2 := NewPriceConfigBuilder().Length(GoUint8ToMoleculeU8(2)).Build()
		p3 := NewPriceConfigBuilder().Length(GoUint8ToMoleculeU8(3)).Build()
		list := NewPriceConfigListBuilder().Push(p1).Push(p2).Push(p3).Build()
		fmt.Println(list.ItemCount())
		priceIndex := uint(0)
		accountBytesLen := uint8(len([]rune("123"))) // 字符长度
		for ; priceIndex < list.ItemCount(); priceIndex++ {
			item := list.Get(priceIndex)
			accountLen, err := MoleculeU8ToGo(item.Length().AsSlice())
			if err != nil {
				panic(err)
			} else if accountLen == accountBytesLen {
				return item
			}
			fmt.Println(accountLen, accountBytesLen, priceIndex)
		}
		return nil
	}
	i := getItem()
	fmt.Println(i.Length().RawData())
}

func Test_U64Bytes(t *testing.T) {
	d, _ := blake2b.Blake256([]byte("0"))
	t.Log(len(d))
	t.Log(len(GoUint64ToBytes(0)))
}

func Test_AccountChar(t *testing.T) {
	t.Log(len([]byte("account")))
}

func Test_CalAccountCellExpiredAt(t *testing.T) {
	// registerAt:=
	// 2021-01-28 18:02:50, 1611828171
	accountCellCap, err := AccountCellCap("11111111.bit")
	if err != nil {
		panic(err)
	}
	// CalAccountCellExpiredAt ====>
	// {"quote":1000,"account_cell_cap":14600000000,"price_config_new":5000000,"pre_account_cell_cap":536800000000,"ref_cell_cap":21000000000}
	param := CalAccountCellExpiredAtParam{
		Quote:             1000, // 1000 ckb = 1 usd
		AccountCellCap:    accountCellCap,
		PriceConfigNew:    5000000, // 10 usd
		PreAccountCellCap: 5368 * OneCkb,
		RefCellCap:        2 * RefCellBaseCap,
	}
	timeSec, err := CalAccountCellExpiredAt(param, time.Now().Unix())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(timeSec)
		fmt.Println(time.Unix(int64(timeSec), 0).String())
		bys, _ := json.Marshal(param)
		t.Log(string(bys))
	}
}

func Test_Blake2b_256(t *testing.T) {
	// 0xc9804583fc51c64512c0153264a707c254ae81ff
	bys, _ := blake2b.Blake160([]byte("das00007.bit"))
	t.Log(hex.EncodeToString(bys))
	t.Log(len(bys), bys)
}

func Test_ParseActionCell(t *testing.T) {
	hexStr := "646173000000001a0000000c0000001600000006000000636f6e66696700000000"
	bys, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatal(err)
	}
	if witness, err := NewDasWitnessDataFromSlice(bys); err != nil {
		t.Fatal(err)
	} else {
		actionData, _ := ActionDataFromSlice(witness.TableBys, false)
		t.Log(witness.Tag, witness.TableType, string(actionData.Action().RawData()))
	}
}

// func Test_StateCellData(t *testing.T) {
// 	stateCell := NewStateCellDataBuilder()
// 	rootHash := HashFromSliceUnchecked([]byte("hello world!h"))
// 	stateCell.ReservedAccountRoot(*rootHash)
// 	// dataBytes := stateCell.Build()
// 	raw := string(stateCell.reserved_account_root.AsSlice())
// 	t.Log("raw ===> ", raw)
// 	t.Log("rawHex ===> ", hex.EncodeToString(stateCell.reserved_account_root.RawData()))
//
// }
