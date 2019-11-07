# nettonet-adapter

本项目继承了ethereum-adapter，主要修改了如下内容：
   
- 重写了Symbol = "NTN"。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，NTN.ini文件，编辑如下内容：

```ini


#wallet api url
ServerAPI = "http://127.0.0.1:1111"
#block chain ID
ChainID = 
# Cache data file directory, default = "", current directory: ./data
dataDir = ""


```

## 官方浏览器

http://18.162.205.20/home