package celltype

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2020
 * FileName: publishaccountcell
 * Author:   LinGuanHong
 * Date:     2020/12/25 5:51 下午
 * Description:
 */

/**
table DataEntity {
    index: Uint32, // 表明此数据项属于 inputs/outputs 中的第几个 cell
    version: Uint32, // 表明 entity 数据结构的版本号
    entity: Bytes, // 代表具体的数据结构
}
*/
var TestNetAccountCell = func(depIndex, oldIndex, newIndex uint32, dep, old *AccountCellData, new *AccountCellFullData) *AccountCellParam {
	acp := &AccountCellParam{
		Version: 1,
		Data: *buildDasCommonMoleculeDataObj(
			depIndex, oldIndex, newIndex, dep, old, &new.AccountInfo),
		CellCodeInfo: DasAccountCellScript,
		AccountCellDatas: AccountCellDatas{
			DepAccountCellData: dep,
			OldAccountCellData: old,
			NewAccountCellData: new,
		},
		AlwaysSpendableScriptInfo: DASCellBaseInfo{
			Dep: DASCellBaseInfoDep{
				TxHash:  types.HexToHash("0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6"),
				TxIndex: 0,
				DepType: types.DepTypeDepGroup,
			},
			Out: DasAnyOneCanSendCellInfo,
		},
	}
	return acp
}

/**
lock: <always_success>
type:
  code_hash: <nft_script>
  type: type
  args: []
data:
  // 20 20 32, account = 72...
  id // 自己的 ID，生成算法为 hash(account)，然后取前 20 bytes
  next // 下一个 AccountCell 的 ID
  hash(data: AccountCellData)
  account // AccountCell 为了避免数据丢失导致用户无法找回自己用户所以额外储存了 account 的明文信息

witness:
  table Data {
    old: table DataEntityOpt {
    	index: Uint32,
    	version: Uint32,
    	entity: AccountCellData
    },
    new: table DataEntityOpt {
      index: Uint32,
      version: Uint32,
      entity: AccountCellData
    },
  }

======
table AccountCellData {
    // The first 160 bits of the hash of account.
    id: AccountId,
    // The lock script of owner.
    owner: Script,
    // The lock script of manager.
    manager: Script,
    account: Bytes,
    // The status of the account, 0 means normal, 1 means being sold, 2 means being auctioned.
    status: Uint8,
    registered_at: Timestamp,
    expired_at: Timestamp,
    records: Records,
}

array AccountId [byte; 20];

option AccountIdOpt (AccountId);

table Record {
    record_type: Bytes,
    record_label: Bytes,
    record_value: Bytes,
    record_ttl: Uint32,
}

vector Records <Record>;
*/

type AccountCell struct {
	p *AccountCellParam
}

func NewAccountCell(p *AccountCellParam) *AccountCell {
	return &AccountCell{p: p}
}

func (c *AccountCell) LockDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.AlwaysSpendableScriptInfo.Dep.TxHash,
			Index:  c.p.AlwaysSpendableScriptInfo.Dep.TxIndex,
		},
		DepType: c.p.AlwaysSpendableScriptInfo.Dep.DepType,
	}
}
func (c *AccountCell) TypeDepCell() *types.CellDep {
	return &types.CellDep{ // state_cell
		OutPoint: &types.OutPoint{
			TxHash: c.p.CellCodeInfo.Dep.TxHash,
			Index:  c.p.CellCodeInfo.Dep.TxIndex, // state_script_tx_index
		},
		DepType: c.p.CellCodeInfo.Dep.DepType,
	}
}
func (c *AccountCell) LockScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.AlwaysSpendableScriptInfo.Out.CodeHash,
		HashType: c.p.AlwaysSpendableScriptInfo.Out.CodeHashType,
		Args:     c.p.AlwaysSpendableScriptInfo.Out.Args,
	}
}
func (c *AccountCell) TypeScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.CellCodeInfo.Out.CodeHash,
		HashType: c.p.CellCodeInfo.Out.CodeHashType,
		Args:     c.p.CellCodeInfo.Out.Args,
	}
}

/**
data:
  // 20 20 32, account = 72...
  id // 自己的 ID，生成算法为 hash(account)，然后取前 20 bytes
  next // 下一个 AccountCell 的 ID
  hash(data: AccountCellData)
  account // AccountCell 为了避免数据丢失导致用户无法找回自己用户所以额外储存了 account 的明文信息
*/

func AccountIdFromOutputData(data []byte) (DasAccountId, error) {
	if size := len(data); size < 52 {
		return nil, fmt.Errorf("AccountIdFromOutputData invalid data, len not enough: %d", size)
	}
	return data[32:52], nil
}

func DefaultAccountCellDataBytes(accountId, nextAccountId []byte) []byte {
	if accountId == nil || len(accountId) != 20 {
		accountId = EmptyAccountId
	}
	if nextAccountId == nil || len(nextAccountId) != 20 {
		nextAccountId = EmptyAccountId
	}
	holder, _ := blake2b.Blake256([]byte(""))
	return append(append(holder, accountId...), nextAccountId...)
}

func accountCellOutputData(newData *AccountCellFullData) ([]byte, error) {
	dataBytes := []byte{}
	accountInfoDataBytes, _ := blake2b.Blake256(newData.AccountInfo.AsSlice())
	dataBytes = append(dataBytes, accountInfoDataBytes...)
	accountBytes := newData.AccountInfo.Account().AsSlice()
	accountIdBytes, _ := blake2b.Blake160(accountBytes)
	dataBytes = append(dataBytes, accountIdBytes...) // id
	if len(newData.NextAccountId) > 0 {
		nextBytes, _ := blake2b.Blake160(newData.NextAccountId)
		dataBytes = append(dataBytes, nextBytes...) // next
	} else {
		dataBytes = append(dataBytes, EmptyAccountId...)
	}
	dataBytes = append(dataBytes, newData.AccountInfo.ExpiredAt().AsSlice()...) // expired_at
	dataBytes = append(dataBytes, accountBytes...)                              // account
	return dataBytes, nil
}

func AccountCellCap(account string) (uint64, error) {
	output := types.CellOutput{
		Lock: &types.Script{
			CodeHash: DasAnyOneCanSendCellInfo.CodeHash,
			HashType: DasAnyOneCanSendCellInfo.CodeHashType,
			Args:     DasAnyOneCanSendCellInfo.Args,
		},
		Type: &types.Script{
			CodeHash: DasAccountCellScript.Out.CodeHash,
			HashType: DasAccountCellScript.Out.CodeHashType,
			Args:     DasAccountCellScript.Out.Args,
		},
	}
	dataBytes := []byte{}
	dataHash, _ := blake2b.Blake256([]byte("0"))
	idBytes, _ := blake2b.Blake160([]byte("0"))
	nextBytes, _ := blake2b.Blake160([]byte("0"))
	expiredAtBytes := GoUint64ToBytes(0)
	accountBytes := []byte(account)

	dataBytes = append(dataBytes, dataHash...)
	dataBytes = append(dataBytes, idBytes...)
	dataBytes = append(dataBytes, nextBytes...)
	dataBytes = append(dataBytes, expiredAtBytes...)
	dataBytes = append(dataBytes, accountBytes...)

	return output.OccupiedCapacity(dataBytes), nil
}

func (c *AccountCell) Data() ([]byte, error) {
	return accountCellOutputData(c.p.AccountCellDatas.NewAccountCellData)
}

func (c *AccountCell) TableType() TableType {
	return TableType_ACCOUNT_CELL
}

func (c *AccountCell) TableData() []byte {
	return c.p.Data.AsSlice()
}
