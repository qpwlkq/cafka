## handler入口
当前存在的问题:  
1. 把整个字节流当成一个整体, 解析成一个request, 甚至把message_size也当做是一个字段, 根本原因还是混淆了Message和Request的概念  
Message = message_size + Header + Body  
可以把Header + Body合起来理解成API Request

改进: 
1. 首先从字节流中取出前四个字节


## 序列化
现在是每个response写一个序列化, 有很多重复代码, 是否能够动态解析struct, 根据字段类型序列化?