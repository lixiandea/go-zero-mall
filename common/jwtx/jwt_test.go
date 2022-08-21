/**
* @Author:Dijiang
* @Description:
* @Date: Created in 21:46 2022/6/17
* @Modified By: Dijiang
 */

package jwtx

import (
	"testing"
	"time"
)

func TestGetUserId(t *testing.T) {
	type args struct {
		secretKey string
		// tokenString string
		iat     int64
		seconds int64
		uid     int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "tes get token",
			args: args{
				secretKey: "asdasdsdasd",
				iat:       time.Now().Unix(),
				seconds:   1000,
				uid:       421224199609213711,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetToken(tt.args.secretKey, tt.args.iat, tt.args.seconds, tt.args.uid)
			if err != nil {
				t.Errorf("Get token error %v", err)
				return
			}
			gotUid, err := GetUserId(tt.args.secretKey, token)
			if err != nil {
				t.Errorf("GetUserId() error = %v", err)
				return
			}
			if gotUid != tt.args.uid {
				t.Errorf("GetUserId() = %v, want %v", gotUid, tt.args.uid)
			}
		})
	}
}
