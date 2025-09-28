package vo

type Password string

func (p Password) Validate() error {

	// TODO
	// checking password strong level
	// will judge the password in range 0 - 100
	// 0 for very weak and 100 for very strong
	// below some threshold password will not allowed to created by outputing error

	return nil
}
