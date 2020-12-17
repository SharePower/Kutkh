package datatransfer

func Transfer(source Transferable, target Transferable, rule TransferRuleExecute) {
	// TODO do something like log? or send?
	rule.execute(source, target)
	// TODO do something like log? or send?
}
