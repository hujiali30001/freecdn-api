package models

import (
	stringutil "github.com/iwind/TeaGo/utils/string"
	"golang.org/x/crypto/bcrypt"
)

// weakPasswordPlaintexts 弱密码明文列表（ORA-08：不再存 MD5，改为运行时 bcrypt 验证）
var weakPasswordPlaintexts = []string{
	"123", "1234", "12345", "123456", "12345678", "123456789",
	"000000", "111111", "666666", "888888", "654321",
	"password", "qwerty", "admin",
}

// weakPasswords 保留旧 MD5 列表，供 DAO 的 hasWeakPasswords 数据库查询使用（存量数据）
var weakPasswords = []string{}

func init() {
	for _, p := range weakPasswordPlaintexts {
		weakPasswords = append(weakPasswords, stringutil.Md5(p))
	}
}

// HasWeakPassword 检测弱密码（ORA-08：兼容 bcrypt 和旧 MD5）
func (this *Admin) HasWeakPassword() bool {
	if len(this.Password) == 0 {
		return false
	}

	// bcrypt 哈希：逐一验证弱密码明文
	if len(this.Password) > 32 && this.Password[0] == '$' {
		for _, plain := range weakPasswordPlaintexts {
			if bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(plain)) == nil {
				return true
			}
		}
		return false
	}

	// 旧 MD5 哈希：直接比对
	for _, md5hash := range weakPasswords {
		if md5hash == this.Password {
			return true
		}
	}
	return false
}
