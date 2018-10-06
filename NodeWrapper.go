package lun

type NodeWrapper struct {
	node   Node
	logger Logger
}

func NewNodeWrapper(node Node, logger Logger) *NodeWrapper {
	return &NodeWrapper{node: node, logger: logger}
}

func (w *NodeWrapper) Node() Node {
	return w.node
}

func (w *NodeWrapper) Logger() Logger {
	return w.logger
}

func (w *NodeWrapper) Run(cmd string) {
	w.run(false, cmd)
}

func (w *NodeWrapper) SafeRun(cmd string) error {
	return w.run(true, cmd)
}

func (w *NodeWrapper) run(ignore bool, cmd string) error {
	if ignore {
		w.logger.Log(LogLevelInfo, "🚙 "+cmd)
	} else {
		w.logger.Log(LogLevelInfo, "🚗 "+cmd)
	}
	output, err := w.node.SafeRun(cmd)
	if err != nil {
		if len(output) > 0 {
			w.logger.Log(LogLevelError, string(output))
		}
		w.logger.Log(LogLevelError, err.Error())
		if !ignore {
			panic(err)
		}
	} else {
		if len(output) > 0 {
			w.logger.Log(LogLevelVerbose, string(output))
		}
	}
	return err
}
