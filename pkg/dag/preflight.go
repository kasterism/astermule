package dag

type preflightChain func() (*DAG, error)

func (d *DAG) Preflight() error {
	// Check node data and fill nodeRef
	err := d.preflightChainStart().
		findEntryNode().
		findExitNode().
		preflightChainEnd()
	return err
}

func (d *DAG) preflightChainStart() preflightChain {
	return func() (*DAG, error) {
		return d, nil
	}
}

func (process preflightChain) findEntryNode() preflightChain {
	return func() (*DAG, error) {
		d, err := process()
		if err != nil {
			return d, err
		}
		isEntry := []string{}
		for _, node := range d.Nodes {
			if len(node.Dependencies) == 0 {
				isEntry = append(isEntry, node.Name)
			}
		}
		d.entry = isEntry
		return d, nil
	}
}

func (process preflightChain) findExitNode() preflightChain {
	return func() (*DAG, error) {
		d, err := process()
		if err != nil {
			return d, err
		}
		isExit := make(map[string]struct{})
		for _, node := range d.Nodes {
			for _, dep := range node.Dependencies {
				isExit[dep] = struct{}{}
			}
		}
		results := []string{}
		for _, node := range d.Nodes {
			if _, ok := isExit[node.Name]; !ok {
				results = append(results, node.Name)
			}
		}
		d.exit = results
		return d, nil
	}
}

func (process preflightChain) preflightChainEnd() error {
	_, err := process()
	if err != nil {
		return err
	}
	return nil
}
