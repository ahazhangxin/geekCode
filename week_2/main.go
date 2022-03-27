package main

import (
	"fmt"
)

// Q. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

// A:关于dao层遇到的 sql.ErrNoRows时，我觉得应该需要由wrap包装抛给上层。
//   因为当前调用者执行玩sql的语句后，语句报错，外部调用者可能想知道是哪个操作导致了sql.ErrNoRows,此时返回方可以使用wrap包装一层返回出去。
//   同时在真正处理错误的地方，可以打印出调用堆栈信息，方便更快地定位问题。

func main() {
	sql, err := InitSql()
	defer sql.Close()
	if err != nil {
		fmt.Println("initSql error:", err)
		fmt.Printf("%+v\n", err)
	}

	err = Query(sql, "select * from user where id=10001")
	if err != nil {
		fmt.Println("query sql error :", err)
		fmt.Printf("%+v\n", err)
	}
}
