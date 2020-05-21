package web

type IMarshaler interface {
	Marshal() ([]byte, error)
}

type WebHeaders map[string]string

type WebResponse struct {
	Status  int
	Headers WebHeaders
	Body    []byte
}

func NewWebResponse(status int, headers WebHeaders, body []byte) *WebResponse {
	return &WebResponse{
		Status:  status,
		Headers: headers,
		Body:    body,
	}
}

func (wb *WebResponse) WriteStruct(status int, headers WebHeaders, m IMarshaler) error {
	body, err := m.Marshal()
	if err != nil {
		return err
	}

	wb.Status = status
	wb.Headers = headers
	wb.Body = body

	return nil
}
