## 关于erda的流水线如何实现和使用

### 例子
```yaml
version: 1.1
stages:
  # 克隆当前分支的代码
  - stage:
      - git-checkout:
          alias: clone-java-code
          params:
            depth: 1
  
  # 克隆maven的setting.xml，用作下面maven打包进行加速和代理，当然你也可以用wget或者任意方式   
  - stage:
      - git-checkout:
          alias: clone-maven-setting
          params:
            depth: 1
            uri: maven的setting的git地址
            branch: 分支 
            username: git的账号
            password: git的密码
  
  # 基于maven的容器对java代码进行打包        
  - stage:
      - custom-script:
          alias: java-mvn-build
          image: maven:3.6.3-openjdk-11
          commands:
            - mkdir ~/.m2
            # 将上面clone-maven-setting步骤克隆的setting.xml放入maven配置的位置，并变更名称为setting.xml
            - cat ${clone-maven-setting}/settings-docker.xml > ~/.m2/settings.xml
            # 进入clone-java-code步骤的工作目录
            - cd ${clone-java-code}
            - mvn package
            # 这里将打包好的文件放入当前目录下(也就是让上下文收录)，当前步骤只有$WORKDIR(环境变量)目录才会收录给上下文
            # 这里jar包虽然生成了，但是需要拷贝到 $WORKDIR 下，这样下文才能获取到生成的jar包
            - cp -r ${clone-java-code}/target/ $WORKDIR 
  
  # 执行docker的构建和上传
  - stage:
      - custom-script:
          alias: docker-build-push
          image: docker
          commands:
            # 设置镜像加速
            - echo "http://mirrors.aliyun.com/alpine/v3.6/main/" > /etc/apk/repositories && echo "http://mirrors.aliyun.com/alpine/v3.6/community/" >> /etc/apk/repositories
            # 将生成的jar包拷贝回clone-java-code步骤下
            - cp -r ${java-mvn-build}/target/ ${clone-java-code}/target/
            - cd ${clone-java-code}
            # 这里只是一些示例，下面构建的时候不设置cpu限制和内存限制也没有关系
            - cpu=$(echo "$PIPELINE_LIMITED_CPU*100000"|bc | awk '{sub(/\./,"",$1); print $1}')
            - memory=$(echo "$PIPELINE_LIMITED_MEM*1048576"|bc)
            # 构建image的名称
            - time=`date +%s%6N`
            - repo=""$BP_DOCKER_ARTIFACT_REGISTRY"/"$DICE_PROJECT_APPLICATION":"$PIPELINE_TASK_NAME"-"$time""
            # --build-arg和--cpu-quota不是必须，这里只是给个例子进行docker的image构建
            - docker build --build-arg DICE_VERSION=$DICE_VERSION --cpu-quota $cpu --memory $memory -t $repo .
            # login和registerAddr地址看用户自定义情况来定，目前dice自带docker的仓库，不需要登录，直接推送即可
            # - docker login --username= --password= registerAddr
            - echo $repo
            # 这里用的是dice的仓库，直接推送即可
            - docker push $repo
            # 写入image的名称给下文的release用
            - echo "image="$repo"" >> $METAFILE
            
  - stage:
      - release:
          params:
            dice_yml: ${clone-java-code}/dice.yml
            image:
              java-demo: ${docker-build-push:OUTPUT:image}

  - stage:
      - dice:
          params:
            release_id: ${release:OUTPUT:releaseID}
```

### 整体总结上面的流水线例子
1. 首先代码克隆，将当前仓库的代码克隆下来
2. java-mvn-build 原本的工作目录是 $WORKDIR (/.pipeline/actions/context/java-mvn-build) 然后主动 cd $(clone-java-code) (/.pipeline/actions/context/clone-java-code), 所以当前是在上个节点的 $WORKDIR 中
3. java-mvn-build 执行了构建命令，然后将生成的 jar 包拷贝到了自己的 $WORKDIR 下，因为 dice 只会上传各个任务 $WORKDIR 的内容，所以不上传到自己的目录那生成的文件就会丢失
4. docker-build-push 则是将之前拷贝的 jar 包重新拷贝回了 $(clone-java-code)，然后进入到 $(clone-java-code) 目录下，执行了一系列命令，最后用 echo "image="$repo"" >> $METAFILE 将镜像名称声明为当前任务的返回值
5. release 接收 ${docker-build-push:OUTPUT:image} 的返回值将其作为入参进行制品的构建

### 简短的说明下 erda 的 pipeline 实现原理

1. 用户通过 pipeline.yml 调用创建流水线的 api, 创建的时候会去校验各个占位符是否存在对应的任务，不存在就会报错，当校验完毕，就会将整个 pipeline.yml 转化为流水线表中的数据，stages 会转化为任务表中的数据，
刚开始所有数据都是初始化阶段，创建完后会往 etcd 中塞入 流水线ID，等待监听执行
2. pipeline 组件启动后会去监听 etcd, 当获取到 pipelineID 的时候，会调用 reconciler 方法逐步推进流水线
3. reconciler 进一步调用 reconcilerTask 去推进任务的执行， task 有一个生命周期调度，在 prepare 阶段的时候会去将占位符替换成对应的值，比如 ${{ outputs.actionName.actionResultName }}, 
就会遍历该任务之上的任务的出参，然后根据任务名称和出参名进行替换成对应的值，${{ configs.key }} 就是获取用户的流水线变量配置，将对应的 key 进行替换，${{ random.randomType }} 会根据需要随机的类型进行随机出数据然后替换，
${{ dirs.actionName }} 就会遍历该任务之上的任务，将各个任务的文件网络挂载并解压到对应的目录，例如有个任务 A 他的工作路径就是 /.pipeline/actions/context/A, 依次类推，各个任务就都在 context 目录下对应的任务名下，流水线中的占位符就会替换成该路径，从而可以访问任务的文件
4. create 当 prepare 状态执行代码后会变更任务状态为 create, create 则会根据 action 的 scheduler 去选择对应的调度器，比如 EDAS，k8s 或者内存调度器去创建任务
5. start，同理根据调度器去执行任务，queue 和 wait 会定时的调用各个调度器的 api 查询任务的状态来更新状态
6. 当一个任务执行完后，就会重新调度 reconciler 根据 DAG 进行计算出下个任务的调用，重新调度 reconcilerTask 去推进下个任务，并行任务则开启协程并发调度

### 扩展说明
1. 任务调度的流程，每个任务被 k8s 调度的时候，并不是直接用任务的镜像去运行启动命令，而是会启动一个叫 actionagent 的命令(代码在erda中)，actionagent 作为 1 号进程被启动，
当任务是 custom-script 的时候 actionagent 则是运行 command，而其他非自定义 action 都会运行其启动命令 /opt/action/run, 由此可见，用户需要开发一个自己的 action，
你只要构建出一个镜像，这个镜像中 /opt/action/run 应当是你的启动命令
2. actionagent 会做很多事情，首先他会帮你挂载 docker.sock 文件，让你在 action 中可以直接连接宿主机的 docker server(后续准备用buildkit来解决安全性问题)，
然后还要根据任务的先后顺序从网盘上下载任务的 tar 包，解压到对应的 /.pipeline/actions/context/actionName，从而实现任务之间文件共享，当子进程结束也就是 /opt/action/run 结束的时候，会去将 $WORKDIR 中的所有文件上传为 tar 包从而进行共享
3. 当子进程结束的时候还会做一件事情，解析用户的日志，当日志中存在 echo "resultName=resultValue" >> $METAFILE 就会将 resultName 作为任务的出参名称，resultValue 作为值存储到任务的表中，等待后续任务的使用


### 流水线其他关键的

#### 嵌套流水线实现

```yaml
- stage:
    snippet:
      alias: job
      params:
        name: value
      snippetConfig:
         name: "xxx"
         source: "xx"
         labels:
            key: value
```

snippetConfig 是 erda 定义的标准的 pipeline.yml 寻址的协议，客户端首先需要在 erda 的 dice_pipeline_snippet_clients 表中注册 source hosts prefixUrl, 然后客户端需要
实现 hots/prefixUrl/actions/batch-query-snippet-yml 这样的 http 服务，该服务接收的请求是 snippetConfig 请求体，然后返回 pipeline.yml 的文件字符串，
由此，source 可以唯一定位 clients 的地址，整体 snippetConfig 请求获取对应的 pipeline.yml

snippet 类型的 action 比较特殊，做了特别的判定，创建流水线时候不只会创建 snippet 的任务，还会根据 snippet 中的 snippetConfig 去唯一定位到一个 pipeline.yml，pipeline 会根据
请求获取的 yml 去生成新的流水线和任务，从而可以实现深层次的嵌套流水线，类似目前实现的 source 就有 local(嵌套应用流水线)，autotest(计划嵌套场景集嵌套场景)，configSheet(配置单互相引用)

pipeline.yml 中新加的 2 个字段就是为了服务嵌套流水线 params 和 output, params 就是声明该 snippet 有什么入参，对应可以在引用 snippet 的时候可以传递参数，output 则是出参，出参目前只能引用某个任务的出参，无法自定义
根据这2个定义，用户使用 snippet 就可以像原生的 erda 的 action 一样，从而可以构想 snippet 可以进一步扩展出市场功能，用户自己可以根据 custom-action 来实现一个 snippet，然后其他人就可以直接使用用户的 action，
用户可以像 dockerhub 那样选择自己喜欢的 action 来构建自己的任务

```yaml
version: 1.1
stages: []
params:
  - name: snippetParamsName
outputs:
  - name: snippetResultName
    ref: ${{ outputs.actionName.actionResult }}
```

额外说明:
目前嵌套流水线，还不支持 ${{ dirs.actionName }} 来使用内部任务的文件

#### 任务的重试和断言

任务的重试会在任务执行后判定任务的状态，当任务状态为失败的时候会根据重试设置的重试时间将任务先设置为初始状态，然后调用 reconcilerTask 进行重试

任务的断言执行，则是在任务 prepare 之后进行断言判断，判定断言字段中的值计算后是否成立，成立则让任务继续往下执行，不成立则跳过


#### 三方业务支持定制接入控制流水线执行逻辑

https://www.yuque.com/erda-project/design/sdvvwl 三方接入的文档

说明: 在 pipeline.yml 中声明好生命周期接入点，client 名称，然后传入的 labels(请求三方时会传入)
目前在 pipeline 中只开放了一个接入点那就是 before-run-check, 意思是任务在执行之前, 接入点的实现逻辑就是根据流水线的 aop 代码接入点进行请求三方 api
，然后根据标准协议获取请求三方的返回值来控制流水线的运行过程


