package reponse

import "time"

type JsonTime time.Time

func (j *JsonTime) MarshalJSON() ([]byte, error) {
	t := time.Time(*j)
	return []byte(t.Format(`"2006-01-02"`)), nil
}

type UserResponse struct {
	Id       int32    `json:"id"`
	NickName string   `json:"name"`
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
}
