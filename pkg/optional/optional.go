package optional

import (
	"encoding/json"
	"time"
)

type optionalValue interface {
	string | bool | int | float64 | time.Time
}

type Optional[T optionalValue] struct {
	Set    bool
	IsNull bool
	Value  T
}

func New[T optionalValue](value T) Optional[T] {
	return Optional[T]{
		Value:  value,
		Set:    true,
		IsNull: false,
	}
}

func (i *Optional[T]) HasValue() bool {
	return i.Set && !i.IsNull
}

func (i *Optional[T]) GetValuePointer() *T {
	if !i.HasValue() {
		return nil
	}
	return &i.Value
}

func (i *Optional[T]) UnmarshalJSON(data []byte) error {
	i.Set = true

	if string(data) == "null" {
		i.IsNull = true
		return nil
	}

	var temp T
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	i.Value = temp
	return nil
}

func (i *Optional[T]) MarshalJSON() ([]byte, error) {
	if i.IsNull {
		body, err := json.Marshal(nil)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	body, err := json.Marshal(i.Value)
	if err != nil {
		return nil, err
	}

	return body, nil
}
