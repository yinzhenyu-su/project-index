# project-index
哲寻项目流水线列表

# 开发模式
将项目clone到本地后，在项目中打开终端执行 `go run . watch <gitlabtoken> -d <要监听的目录>`
然后在浏览器中打开 [http://localhost:8080](http://localhost:8080)

# 打包发布
打开终端执行 `./build.sh` （目前只配置了mac下打包linux可执行文件）