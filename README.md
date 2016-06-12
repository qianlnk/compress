# compress
compress file or dir by tar and gzip. and uncompress it. 
## use
```golang
//tar cvzf helloXzj.tar.gz -C /Users/xiezhenjia/go/src/qianno.xie/superservice/superservicectl/programs/ helloLnk
Compress("/Users/xiezhenjia/go/src/qianno.xie/superservice/superservicectl/programs/helloLnk", "helloXzj.tar.gz")

//tar xzvf helloXzj.tar.gz /tmp
Uncompress("helloXzj.tar.gz", "/tmp")
```

