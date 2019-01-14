# godemo
golang特性使用

https://gobyexample.com/

## tcp 滑动窗口测试
```
tail -f stream.txt | nc localhost 3040
while true; do echo "foo" > stream.txt; done
```