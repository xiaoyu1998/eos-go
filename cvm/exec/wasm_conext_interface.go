package exec

import (
	"github.com/eosspark/eos-go/chain/types"
	"github.com/eosspark/eos-go/common"
)

type WasmContextInterface interface {

	//action api
	GetActionData() []byte
	GetReceiver() common.AccountName
	GetCode() common.AccountName
	GetAct() common.ActionName

	//context authorization api
	RequireAuthorization(account common.AccountName)
	HasAuthorization(account common.AccountName) bool
	RequireAuthorization2(account common.AccountName, permission common.PermissionName)
	//RequireAuthorizations(account common.AccountName)
	RequireRecipient(recipient common.AccountName)
	IsAccount(n common.AccountName) bool
	HasReciptient(code common.AccountName) bool

	//contet console text
	ResetConsole()
	ContextAppend(str string)

	//context database api
	DbStoreI64(scope int64, table int64, payer int64, id int64, buffer []byte) int
	DbUpdateI64(iterator int, payer int64, buffer []byte)
	DbRemoveI64(iterator int)
	DbGetI64(iterator int, buffer []byte, bufferSize int) int
	DbNextI64(iterator int, primary *uint64) int
	DbPreviousI64(iterator int, primary *uint64) int
	DbFindI64(code int64, scope int64, table int64, id int64) int
	DbLowerBoundI64(code int64, scope int64, table int64, id int64) int
	DbUpperBoundI64(code int64, scope int64, table int64, id int64) int
	DbEndI64(code int64, scope int64, table int64) int

	IdxI64Store(scope int64, table int64, payer int64, id int64, value *types.Uint64_t) int
	IdxI64Remove(iterator int)
	IdxI64Update(iterator int, payer int64, value *types.Uint64_t)
	IdxI64FindSecondary(code int64, scope int64, table int64, secondary *types.Uint64_t, primary *uint64) int
	IdxI64LowerBound(code int64, scope int64, table int64, secondary *types.Uint64_t, primary *uint64) int
	IdxI64UpperBound(code int64, scope int64, table int64, secondary *types.Uint64_t, primary *uint64) int
	IdxI64End(code int64, scope int64, table int64) int
	IdxI64Next(iterator int, primary *uint64) int
	IdxI64Previous(iterator int, primary *uint64) int
	IdxI64FindPrimary(code int64, scope int64, table int64, secondary *types.Uint64_t, primary *uint64) int

	UpdateDbUsage(pager common.AccountName, delta int64)
	//FindTable(code int64, scope int64, table int64) types.TableIDObject
	//FindOrCreateTable(code common.Name, scope common.Name, table common.Name, payer *common.AccountName) types.TableIDObject
	RemoveTable(tid types.TableIdObject)

	//context permission api
	GetPermissionLastUsed(account common.AccountName, permission common.PermissionName) int64
	GetAccountCreateTime(account common.AccountName) int64

	//context privileged api
	SetResourceLimits(account common.AccountName, ramBytes uint64, netWeight uint64, cpuWeigth uint64)
	GetResourceLimits(account common.AccountName, ramBytes *uint64, netWeight *uint64, cpuWeigth *uint64)
	SetBlockchainParametersPacked(parameters []byte)
	GetBlockchainParametersPacked() []byte
	IsPrivileged(n common.AccountName) bool
	SetPrivileged(n common.AccountName, isPriv bool)

	//context producer api
	SetProposedProducers(producers []byte)
	GetActiveProducersInBytes() []byte
	//GetActiveProducers() []common.AccountName

	//context system api
	CheckTime()
	CurrentTime() int64
	PublicationTime() int64

	//context transaction api
	ExecuteInline(action []byte)
	ExecuteContextFreeInline(action []byte)
	ScheduleDeferredTransaction(sendId common.TransactionIdType, payer common.AccountName, trx []byte, replaceExisting bool)
	CancelDeferredTransaction(sendId common.TransactionIdType) bool
	GetPackedTransaction() []byte
	Expiration() int
	TaposBlockNum() int
	TaposBlockPrefix() int
	GetAction(typ uint32, index int, bufferSize int) (int, []byte)
	GetContextFreeData(intdex int, bufferSize int) (int, []byte)
}
