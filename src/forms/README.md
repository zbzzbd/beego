forms对象层
=====

form层用于映射http提交的表单, 功能有如下:

1. 基本的数据类型验证
2. 基本的业务逻辑数据验证
> 比如时间验证, 是否必要, 数字, 字母等; 此验证不涉及下一层的数据验证, 比如邮箱是否被使用, 用户名是否存在
3. 返回已经验证过的数据对象
4. 基本的安全保护, 比如防止重复提交等

使用库
----
[binding](https://github.com/mholt/binding)
> 将http request里的数据绑定到一个struct上
