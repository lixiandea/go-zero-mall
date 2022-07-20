/**
* @Author:Dijiang
* @Description:
* @Date: Created in 14:22 2022/7/27
* @Modified By: Dijiang
 */

package main

import "syscall"

func main() {
	dll := syscall.NewLazyDLL("exportgo.dll")
	printHello := dll.NewProc("PrintHello")
	printHello.Call()
}
