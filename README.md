# socks5-go





# socks5      200行golang代码实现 翻墙梯



### #打包

```
go build ./src/main/socks5.go
```





#使用现有执行文件

```
git clone https://github.com/jackoka/socks5-go.git && cd socks5-go && git checkout prod && cd out && chmod +x socks5 && ./socks5 port=50000
```



## #运行

- ##### linux

  授权:

  ```
  chmod +x socks5
  ```

  指定端口号:

  ```
  ./socks5 port=50000
  ```

  
  
  
  
  后台运行:
  
  ```
  #不保存日志
  nohup ./socks5 >/dev/null 2>&1 &
  
  #保存日志
  nohup ./socks5 >./socks5.log  &
  ```



- windows

  ###### 双击 socks5.exe

  

  指定端口号运行:

  ```
  socks5.exe port=50000
  ```








# client

### 谷歌浏览器安装 SwitchyOmega 插件

代理协议：socks5

代理服务器：代理服务器外网ip地址

代理端口：上面程序启动用的端口，默认 50000

