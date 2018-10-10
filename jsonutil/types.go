package jsonutil

import (
	"fmt"
	"strconv"
	"strings"
)

type ByteJsonArray []byte

// Implement json.Marshaler
func (u ByteJsonArray) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = strings.Join(strings.Fields(fmt.Sprintf("%d", u)), ",")
	}
	return []byte(result), nil
}

// Implement json.Unmarshaler
func (u *ByteJsonArray) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `[]`)
	tok := strings.Split(str, ",")
	res := make([]byte, 0)
	for _, t := range tok {
		i, err := strconv.ParseInt(strings.TrimSpace(t), 0, 10)
		if err != nil {
			return err
		}
		if i < 0 || i > 255 {
			return fmt.Errorf("invalid byte value %d, must be between 0 and 255", i)
		}
		res = append(res, byte(i))
	}

	*u = ByteJsonArray(res)
	return nil
}

type ByteJsonString []byte

func (u ByteJsonString) String() string {
	return string(u)
}

func (u ByteJsonString) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = fmt.Sprintf(`"%s"`, u)
	}
	return []byte(result), nil
}

// Implement json.Unmarshaler
func (u *ByteJsonString) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = strings.Trim(str, `"`)
	*u = ByteJsonString(str)
	return nil
}
