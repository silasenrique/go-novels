package model

/*
{
  "error": {
    "code": 403,
      "message": "@ClassroomApiDisabled The user is not permitted to access the Classroom API.",
      "errors": [
        {
          "message": "@ClassroomApiDisabled The user is not permitted to access the Classroom API.",
          "domain": "global",
          "reason": "forbidden"
        }
      ],
      "status": "PERMISSION_DENIED"
  }
}
*/

type internalError struct {
	Err err `json:"error"`
}

type err struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Errors  []*detailErr `json:"errors"`
}

type detailErr struct {
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
}

func NewDetailError(message, domain, reason string) *detailErr {

}

func NewError(message string) *internalError {
	return &internalError{
		Err: err{
			Message: message,
		},
	}
}

func (e *internalError) Error() string {
	return ""
}
