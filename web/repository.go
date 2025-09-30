package web

import (
	"github.com/rosso-ai/conlai/conlpb"
	"sync"
)

type repoData struct {
	params *conlpb.ConLParams
}

type Repository struct {
	data    repoData
	mu      sync.Mutex
	isEmpty bool
}

func (r *Repository) Enqueue(p *conlpb.ConLParams) {
	var data repoData
	data.params = p

	r.mu.Lock()
	r.data = data
	r.mu.Unlock()

	r.isEmpty = false
}

func (r *Repository) Dequeue() *conlpb.ConLParams {
	var params conlpb.ConLParams
	params.Params = []byte("")

	r.mu.Lock()
	if !r.isEmpty {
		params.Params = r.data.params.GetParams()
	}
	r.mu.Unlock()

	return &params
}
