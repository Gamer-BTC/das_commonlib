package celltype

/**
 * Copyright (C), 2019-2020
 * FileName: value
 * Author:   LinGuanHong
 * Date:     2020/12/20 3:12 下午
 * Description:
 */

const witnessDas = "das"
const oneDaySec = float64(24 * 2600)
const oneYearDays = float64(365)
const CellVersionByteLen = 4
const MoleculeBytesHeaderSize = 4
const OneCkb = uint64(1e8)
const DasAccountSuffix = ".bit"
const CkbTxMinOutputCKBValue = 61 * OneCkb
const AccountCellDataAccountIdStartIndex = 72
const RefCellCap = 114 * OneCkb

var EmptyAccountId = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TableType uint32
type AccountCharType uint32
type AccountCellStatus uint8
type DataEntityChangeType uint

func (a AccountCellStatus) Str() string {
	switch a {
	case AccountWitnessStatus_Exist:
		return "exist"
	case AccountWitnessStatus_New:
		return "new"
	case AccountWitnessStatus_Proposed:
		return "proposed"
	}
	return "unknown"
}

type LockScriptType uint

const (
	ScriptType_User LockScriptType = 0
	ScriptType_Any  LockScriptType = 1
)

type DasAccountSearchStatus int

const (
	SearchStatus_NotOpenRegister  DasAccountSearchStatus = -1
	SearchStatus_Registerable     DasAccountSearchStatus = 0
	SearchStatus_ReservedAccount  DasAccountSearchStatus = 1 // 候选
	SearchStatus_OnePriceSell     DasAccountSearchStatus = 2
	SearchStatus_AuctionSell      DasAccountSearchStatus = 3 // 竞拍，候选 -> 竞拍
	SearchStatus_CandidateAccount DasAccountSearchStatus = 4
	SearchStatus_Registered       DasAccountSearchStatus = 5
)

type MarketType int

// 0x01 for primary，0x02 for secondary
const (
	MarketType_Primary   = 1
	MarketType_Secondary = 2
)

const (
	AccountChar_Emoji AccountCharType = 0
	AccountChar_En    AccountCharType = 1
	AccountChar_Zh_Cn AccountCharType = 2
)

const (
	TableType_ACTION       TableType = 0
	TableType_CONFIG_CELL  TableType = 1
	TableType_ACCOUNT_CELL TableType = 2
	// TableType_REGISTER_CELL TableType = 3
	TableType_ON_SALE_CELL        TableType = 3
	TableType_BIDDING_CELL        TableType = 4
	TableType_PROPOSE_CELL        TableType = 5
	TableType_PRE_ACCOUNT_CELL    TableType = 6
	TableTyte_APPLY_REGISTER_CELL TableType = 7 // todo change it
)

const (
	/**
	- status ，状态字段：
	  - 0 ，正常；
	  - 1 ，出售中；
	  - 2 ，拍卖中；
	*/
	AccountCellStatus_Normal AccountCellStatus = 0
	AccountCellStatus_OnSale AccountCellStatus = 1
	AccountCellStatus_OnBid  AccountCellStatus = 2
)

const (
	AccountWitnessStatus_Exist    AccountCellStatus = 0
	AccountWitnessStatus_Proposed AccountCellStatus = 1
	AccountWitnessStatus_New      AccountCellStatus = 2
)

const (
	NewToDep   DataEntityChangeType = 0
	NewToInput DataEntityChangeType = 1
	DepToInput DataEntityChangeType = 2
	depToDep   DataEntityChangeType = 3
)

const (
	CkbSize_AccountCell = 147 * OneCkb
)

const (
	SystemScript_ApplyRegisterCell = "apply_register_cell"
	SystemScript_PreAccoutnCell    = "preaccount_cell"
	SystemScript_AccoutnCell       = "account_cell"
	SystemScript_BiddingCell       = "bidding_cell"
	SystemScript_OnSaleCell        = "on_sale_cell"
	SystemScript_ProposeCell       = "propose_cell"
	SystemScript_WalletCell        = "wallet_cell"
	SystemScript_RefCell           = "ref_cell"
)

const (
	Action_Deploy                = "deploy"
	Action_Config                = "config"
	Action_InitLink              = "init_linked_list"
	Action_ApplyRegister         = "apply_register"
	Action_PreRegister           = "pre_register"
	Action_CreateWallet          = "create_wallet"
	Action_DeleteWallet          = "delete_wallet"
	Action_Propose               = "propose"
	Action_TransferAccount       = "transfer_account"
	Action_RenewAccount          = "renew_account"
	Action_ExtendPropose         = "extend_propose"
	Action_ConfirmProposal       = "confirm_proposal"
	Action_RecyclePropose        = "recycle_propose"
	Action_WithdrawFromWallet    = "withdraw_from_wallet"
	Action_Register              = "register"
	Action_VoteBiddingList       = "vote_bidding_list"
	Action_PublishAccount        = "publish_account"
	Action_RejectRegister        = "reject_register"
	Action_PublishBiddingList    = "publish_bidding_list"
	Action_BidAccount            = "bid_account"
	Action_EditManager           = "edit_manager"
	Action_EditRecords           = "edit_records"
	Action_CancelBidding         = "cancel_bidding"
	Action_CloseBidding          = "close_bidding"
	Action_RefundRegister        = "refund_register"
	Action_QuotePriceForCkb      = "quote_price_for_ckb"
	Action_StartAccountSale      = "start_account_sale"
	Action_CancelAccountSale     = "cancel_account_sale"
	Action_StartAccountAuction   = "start_account_auction"
	Action_CancelAccountAuction  = "cancel_account_auction"
	Action_AccuseAccountRepeat   = "accuse_account_repeat"
	Action_AccuseAccountIllegal  = "accuse_account_illegal"
	Action_RecycleExpiredAccount = "recycle_expired_account_by_keeper"
	Action_CancelSaleByKeeper    = "cancel_sale_by_keeper"
)
