/**
  @Author 1ch0
  @Date 16:51 2021/9/18
  @Description //TODO
 **/

package framework

import "net/http"

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) HttpHandler(rep http.Response, r http.Request) {
	//
}
