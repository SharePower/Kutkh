package datatransfer

type TransferRule struct {
	TransferRuleExecute
}

type TransferRuleExecute interface {
	execute(source Transferable, target Transferable)
}

func (self *TransferRule) execute(source Transferable, target Transferable) {
}
