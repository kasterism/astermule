package dag

type DAG struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Name         string   `json:"name"`
	Action       string   `json:"action" default:"GET"`
	URL          string   `json:"url"`
	Dependencies []string `json:"dependencies,omitempty"`
}

func NewDAG() *DAG {
	return &DAG{
		Nodes: []Node{},
	}
}
