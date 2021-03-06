//go:binary-only-package
//---------------------------------------------------------------------//
//                    GNU GENERAL PUBLIC LICENSE                       //
//                       Version 2, June 1991                          //
//                                                                     //
// Copyright (C) Wells Hsu, wellshsu@outlook.com, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies        //
// of this license document, but changing it is not allowed.           //
//                  SEE LICENSE.md FOR MORE DETAILS.                   //
//---------------------------------------------------------------------//
package xobj

// 基类
type OBJECT struct {
	CHILD interface{} `orm:"-" json:"-"` // 实例指针对象
}

// 构造函数
func (this *OBJECT) CTOR(CHILD interface{}) {
	return
}
