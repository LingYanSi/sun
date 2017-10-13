package model

import "fmt"

/**
 * 对于Model层的数据构建
 * 需要在Model内声明数据类型，并提供数据的增删改查方法
 */

//  Detail 详情
type Detail struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Details []Detail

// DetailSelect 获取详情页所有信息
func DetailSelect() Details {
	var details Details
	// 数据查询222
	rows, err := DB.Query("SELECT name,avatar FROM user")
	if err != nil {
		fmt.Println("数据查询失败", err)
	}
	defer rows.Close()
	// 需要注意的是Scan参数要和
	for rows.Next() {
		detail := Detail{}
		err := rows.Scan(&detail.Name, &detail.Avatar)
		if err != nil {
			fmt.Println("解析失败", err)
		} else {
			details = append(details, detail)
		}
	}
	return details
}
