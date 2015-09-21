package api

type CloudflareListResponse struct {
  Request struct {
    Act string `json:"act"`
  } `json:"request"`
  Response struct {
    Zones   ZoneList   `json:"zones"`
    Records RecordList `json:"recs"`
  } `json:"response"`
  Result string `json:"result"`
  Msg    string `json:"msg"`
}

type CloudFlareRecordResponse struct {
  Request struct {
    Act string `json:"act"`
  } `json:"request"`
  Response struct {
    Rec struct {
      Record Record `json:"obj"`
    } `json:"rec"`
  } `json:"response"`
  Result string `json:"result"`
  Msg    string `json:"msg"`
}
