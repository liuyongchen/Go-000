学习笔记

sql.ErrNoRows 为空值结果，该结果应该返回给上层调用者，但是不应全盘使用wrap 将错误返回，
应该根据具体情况返回结果。



```go
func selectDB()(string, error) {
    res := ""
    err := sql.selectdb("sql words!",&res)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", nil
        }
        return "", errors.wrap(err, "get sql words err")
    }
    return res, nil
}
```

