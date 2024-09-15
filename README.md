# TextSearch

## Window环境

#### **Use Poses**
```
go run main.go -path=[PATH] -key=[KEY]
```
- -path: 查询根目录
- -key: 要查询的关键词

#### **Example**

```
go run main.go -path=D:\ -key=you
```
OR
```
./run.bat
```
## Performance
#### **建立倒排索引**


```
62029
搜索文件耗时: 6.0367245s
程序总耗时: 11.0824958s
```
6.2w个文件，建立倒排索引耗时约为5s（包括写入文件的时间）,结果保存在mr-out文件