/**
* @Author:Dijiang
* @Description:
* @Date: Created in 22:20 2022/7/17
* @Modified By: Dijiang
 */

package main

import (
	"context"
	"github.com/bndr/gojenkins"
)

func main() {
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, "http://10.95.129.129:8080/job/P7/", "dijiang", "11ccf12247e8b8fd8c7b4868bc2738d284")
	_, err := jenkins.Init(ctx)
	if err != nil {
		panic(err)
	}
	job, _ := jenkins.GetJob(ctx, "FightLandlord_Android_dev_new")
	lastBuild, _ := job.GetLastBuild(ctx)
	p := lastBuild.GetParameters()
	println(lastBuild.GetBuildNumber())
	parameter := make(map[string]string, len(p))
	for _, v := range p {
		parameter[v.Name] = v.Value
	}
	simple, err := job.InvokeSimple(ctx, parameter)
	if err != nil {
		panic(err)
	}
	println(simple)
}
