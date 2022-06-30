/**
* @Author:Dijiang
* @Description:
* @Date: Created in 21:17 2022/6/17
* @Modified By: Dijiang
 */

package cryptx

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
