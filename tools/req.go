package tools

type (
    CreateLangReq struct {
        Name     string `json:"name"`
        Paradigm string `json:"paradigm"`
    }
    CreateLangResp struct {
        s string `json:"success"`
    }
    GetLangReq struct {
        id string `json:"id"`
    }
    GetLangResp struct {
        Name string `json:"name"`
    }
)

func encodeResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeLangReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateLangReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeNameReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetLangReq
	vars := mux.Vars(r)

	req = GetLangReq {
		Id: vars["id"],
	}
	return req, nil
}
