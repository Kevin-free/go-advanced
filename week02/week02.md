Q:1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A:当遇到一个 sql.ErrNoRows 的时候，不应该 Wrap 这个 error。因为作为第三方调用的基础库，不应该包装错误，应该返回它的根因。

sql.ErrNoRows 的实现如下：

var ErrNoRows = errors.New("sql: no rows in result set")