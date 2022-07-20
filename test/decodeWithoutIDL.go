/**
* @Author:Dijiang
* @Description:
* @Date: Created in 11:54 2022/8/5
* @Modified By: Dijiang
 */

package main

import (
	"fmt"
	thrifter "github.com/thrift-iterator/go"
	"github.com/thrift-iterator/go/general"
)

func main() {
	//var msg general.Message
	//strMsg := "0000000500000012"
	//thriftEncodedBytes, _ := hex.DecodeString(strMsg)
	//err := thrifter.Unmarshal(thriftEncodedBytes, &msg)
	//if err != nil {
	//	fmt.Println("decode error", err.Error())
	//	return
	//}
	//// the RPC call method name, type is string
	//fmt.Println(msg.MessageName)
	//// the RPC call arguments, type is general.Struct
	//fmt.Println(msg.Arguments)
	//type CPing struct {
	//}
	//thrifterBytes, err := thrifter.Marshal(CPing{})
	//if err != nil {
	//	fmt.Println("decode error", err.Error())
	//}
	//fmt.Println(thrifterBytes)
	thriftEncodedBytes, _ := thrifter.Marshal([]int{1, 2, 3})
	// unmarshal back
	//var val []int
	var msg general.Message
	data := thrifter.Unmarshal(thriftEncodedBytes, &msg)
	fmt.Println(data)
}
