### 项目有需要需要配置私有仓库地址

```yaml
 - stage:
      - java-build:
          alias: java-build
          version: "1.0"
          params:
            build_cmd:
              - rm -rf /usr/share/maven/conf/setting.xml && mvn package -s ((maven_setting))
            jdk_version: 8
            workdir: ${git-checkout}
```

- build_cmd list 类型，可以输入多行 shell 命令来进行构建
- jdk_version jdk 的版本 暂时只支持 8 和 11
- workdir 工作目录

当项目需要配置私有仓库地址的时候，这时候可以用 java-build 进行打包，java-build 的 main 方法很简单，先根据 workDir 来 cd 到某个目录 
, 然后执行 build_cmd 参数的命令，最后将当前 cd 目录的文件保存到自己的 $WORKDIR 中, 这样下面的 action 就可以拿到当前镜像的文件上下文，
然后再打印几个出参到 $METAFILE 中，下面的 action 就可以拿到值上下文(不清楚action的原理可以看 action 文档)

((maven_setting)) 对应的是流水线的变量配置中上传的文件，具体可以看 pipeline 的实现

目前已知的问题，因为 mvn -s 的意思是 user repository 的备用，所以直接 -s 不会直接生效，和 user repository 同一优先级
要把 java-build action 原先初始化好的 /usr/share/maven/conf/setting.xml 给删除才能生效
