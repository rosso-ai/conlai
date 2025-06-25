package web

import "github.com/rosso-ai/conlai/conlpb"

type repoData struct {
	params *conlpb.ConLParams
	len    int
}

type Repository struct {
	data []repoData
}

func (r *Repository) isEmpty() bool {
	return len(r.data) == 0
}

func (r *Repository) Enqueue(p *conlpb.ConLParams) {
	var data repoData
	data.params = p
	data.len = len(p.Params)
	r.data = append(r.data, data)
}

func (r *Repository) Dequeue() *conlpb.ConLParams {
	var params conlpb.ConLParams
	params.Params = []byte("")
	if !r.isEmpty() {
		data := r.data[0]
		params.Params = data.params.GetParams()
		r.data = r.data[1:]
	}
	return &params
}
