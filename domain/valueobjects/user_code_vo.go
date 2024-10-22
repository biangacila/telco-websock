package valueobjects

// UserCode Value Object representing a user code
type UserCode struct {
	Code string
}

func NewUserCode(code string) UserCode {
	return UserCode{
		Code: code,
	}
}
