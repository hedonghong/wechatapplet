package wechatapplet

import "fmt"

type CommonError struct {
	Code    int
	Message string
}

func (c *CommonError) Error() string {
	return fmt.Sprintf("wechat applet error: %s, %d", c.Message, c.Code)
}
