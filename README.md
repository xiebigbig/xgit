# xgit
一个纯go编写 版本控制工具

# 编译

    git clone https://github.com/xiebigbig/xgit.git

    cd xgit

    go mod tidy #自行启用go mod

    go build -o xgit main.go

    cp xgit /usr/local/bin/


### 提交版本
    xgit add

### 查看所有版本
    xgit list

### 回退版本
    xgit goto <commitId>



# todo

1.忽略文件和目录

2.版本差异比较

3.备份仓库远程s3、webdav、sftp、Alist

# 协议
MIT License
