package tools

type Lang struct {
    ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
    Paradigm string `json:"paradigm"`
}

type Rep interface {
    CreateLang(ctx context.Context, lang Lang) error
    GetLang(ctx context.Context, id string) (string, error)
}
