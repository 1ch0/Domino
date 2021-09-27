/**
  @Author 1ch0
  @Date 15:58 2021/9/27
  @Description //TODO
 **/

package main

import "hades/framework"

func registerRouter(core *framework.Core) {

	core.Get("foo", FooControllerHandler)
}
