# socks5-go





# socks5      200行golang代码实现 翻墙梯



### #打包

```
go build ./src/main/socks5.go
```



## #运行

- ##### linux

  授权:

  ```
  chmod +x socks5
  ```

  指定端口号:

  ```
  ./socks5 port=8888
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
  socks5.exe port=8888
  ```

  
