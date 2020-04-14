package tools

import (
    "log"
    "context"
    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/log/level"
    "github.com/gofrs/uuid"
)

type serv struct {
    rep Rep
    log log.Logger
}

func NewServ(rep Rep, log log.Logger) Service {
    return &serv {
        rep: rep,
        log: log
    }
}

func (sv service) CreateLang(ctx context.Context, name string, paradigm string) (string, error) {
    log := log.With(sv.log, "method", "CreateLang")

    u, err := uuid.NewV4()
    if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}

    id := uuid.String()
    lang := Lang {
        ID:       id,
        Name:     name,
        Paradigm: paradigm
    }

    if errs := sv.rep.CreateLang(ctx, lang); errs != nil {
        level.Error(log).Log("error", errs)
		return "", errs
    }

    log.Log("Create lang", id)
    return "Created" nil
}

func (sv service) GetLang(ctx context.Context, id string) (string, error) {
    log := log.With(sv.log, "method", "GetLang")

    name, paradigm, err := sv.rep.GetLang(ctx, id)

    if err != nil {
        level.Error(log).Log("error", err)
        return "", err
    }

    log.Log("Get lang", id)
    return name, paradigm, nil
}
