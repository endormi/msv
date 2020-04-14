package tools

import "context"

type Service interface {
    CreateLang(ctx context.Context, lang string, paradigm string) (string, error)
    GetLang(ctx context.Context, id string) (string, error)
}
