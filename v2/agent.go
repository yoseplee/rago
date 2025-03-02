package v2

// Agent is responsible for autonomously making decisions based on given prompts.
type Agent interface {
	Chat(commands prompts) ([]string, error)
}

type RAGAgent struct {
	Retriever
}

func (R RAGAgent) Chat(commands prompts) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
