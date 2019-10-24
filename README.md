# trendrepo

Github Trending (https://github.com/trending) CLI

## How To Use

### Get

```
$ go get github.com/nmemoto/trendrepo
```

### Run

```
$ trendrepo
REPO NAME                               LANG       STARS FORKS STARS IN PERIOD HREF                                                       DESC
T-head-Semi/wujian100_open              Verilog    746   230   91              https://github.com/T-head-Semi/wujian100_open              IC design and development should be faster，simpler and more reliable
tinacms/tinacms                         TypeScript 2078  70    217             https://github.com/tinacms/tinacms                         Tina is a site editing toolkit for modern React-based sites (Gatsby and Next.js)
TheAlgorithms/Java                      Java       18061 6793  75              https://github.com/TheAlgorithms/Java                      All Algorithms implemented in Java
elastic/logstash                        Ruby       10686 2856  6               https://github.com/elastic/logstash                        Logstash - transport and process your logs, events, or other data
zeit/hyper                              JavaScript 31574 2627  104             https://github.com/zeit/hyper                              A terminal built on web technologies
keras-team/keras                        Python     44982 17093 35              https://github.com/keras-team/keras                        Deep Learning for humans
rancher/rancher                         Go         12917 1474  69              https://github.com/rancher/rancher                         Complete container management platform
Naotonosato/Blawn                       C++        621   23    167             https://github.com/Naotonosato/Blawn                       Pleasant Programming Language.
Netflix/mantis                          Java       214   20    37              https://github.com/Netflix/mantis                          A platform that makes it easy for developers to build realtime, cost-effective, operations-focused applications
cookieY/Yearning                        Go         2663  980   39              https://github.com/cookieY/Yearning                        受欢迎的 Mysql sql审核平台
knrt10/kubernetes-basicLearning         JavaScript 509   43    174             https://github.com/knrt10/kubernetes-basicLearning         Understand kubernetes step by step. A simple repo for beginners 🔥
bregman-arie/devops-interview-questions Groovy     1450  157   185             https://github.com/bregman-arie/devops-interview-questions Linux, Jenkins, AWS, SRE, Prometheus, Docker, Python, Ansible, Git, Kubernetes, Terraform, OpenStack, SQL, NoSQL, Azure, GCP
trimstray/nginx-admins-handbook         Shell      10378 741   156             https://github.com/trimstray/nginx-admins-handbook         How to improve NGINX performance, security, and other important things; @ssllabs A+ 100%, @mozilla A+ 120/100.
json-iterator/go                        Go         6253  533   89              https://github.com/json-iterator/go                        A high-performance 100% compatible drop-in replacement of "encoding/json"
quantopian/zipline                      Python     9927  2978  110             https://github.com/quantopian/zipline                      Zipline, a Pythonic Algorithmic Trading Library
kmario23/deep-learning-drizzle          Python     6302  1428  110             https://github.com/kmario23/deep-learning-drizzle          Drench yourself in Deep Learning, Reinforcement Learning, Machine Learning, Computer Vision, and NLP by learning from these exciting lectures!!
alibaba/flutter-go                      Dart       17788 2484  30              https://github.com/alibaba/flutter-go                      flutter 开发者帮助 APP，包含 flutter 常用 140+ 组件的demo 演示与中文文档
NobyDa/Script                           JavaScript 134   106   10              https://github.com/NobyDa/Script                           flutter 开发者帮助 APP，包含 flutter 常用 140+ 组件的demo 演示与中文文档
dotnetcore/WTM                          C#         1321  319   26              https://github.com/dotnetcore/WTM                          WTM框架是针对中小规模后台管理系统的开发利器。基于DotNetCore，实现0编码创建项目，0编码生成业务模块。框架严格遵循MVVM的开发模式，并深得MVVM的精髓。对于新手，可以快速上手搭建项目；对于高手，可以把那些繁琐重复的工作交给框架生成，专心攻克需求难点。框架经过数十个真实项目检测，可以极大提高开发效率，降低开发成本。
goldbergyoni/nodebestpractices          JavaScript 34942 2940  101             https://github.com/goldbergyoni/nodebestpractices          ✅ The largest Node.js best practices list (September 2019)
ArduPilot/ardupilot                     C++        4436  8414  26              https://github.com/ArduPilot/ardupilot                     ArduPlane, ArduCopter, ArduRover source
x899/chrome_password_grabber            Python     322   53    33              https://github.com/x899/chrome_password_grabber            Get unencrypted 'Saved Password' from Google Chrome
moby/moby                               Go         55286 15954 40              https://github.com/moby/moby                               Moby Project - a collaborative project for the container ecosystem to assemble container-based systems
evilsocket/pwnagotchi                   Python     1941  213   97              https://github.com/evilsocket/pwnagotchi                   (⌐■_■) - Deep Reinforcement Learning instrumenting bettercap for WiFi pwning.
apachecn/AiLearning                     Python     20520 7250  85              https://github.com/apachecn/AiLearning                     AiLearning: 机器学习 - MachineLearning - ML、深度学习 - DeepLearning - DL、自然语言处理 NLP
```

### Option

```
$ trendrepo -h
Usage of trendrepo:
  -b    Open Repository Web Page
  -browser
        Open Repository Web Page
  -f string
        List Format: text or json (default "text")
  -format string
        List Format: text or json (default "text")
  -l string
        Programming Language: go, typescript, ruby, .... anything is ok!
  -lang string
        Programming Language: go, typescript, ruby, .... anything is ok!
  -p string
        Date Range: today, weekly or monthly (default "today")
  -period string
        Date Range: today, weekly or monthly (default "today")
```

### example

```
$ trendrepo -l go -p weekly
REPO NAME                                 LANG STARS FORKS STARS IN PERIOD HREF                                                         DESC
dapr/dapr                                 Go   3287  130   2               https://github.com/dapr/dapr                                 Dapr is a portable, event-driven, runtime for building distributed applications across cloud and edge.
pingcap/tidb                              Go   21011 3117  93              https://github.com/pingcap/tidb                              TiDB is an open source distributed HTAP database compatible with the MySQL protocol
tinygo-org/tinygo                         Go   4744  203   179             https://github.com/tinygo-org/tinygo                         Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM.
helm/helm                                 Go   14441 4693  140             https://github.com/helm/helm                                 The Kubernetes Package Manager
bettercap/bettercap                       Go   6118  584   110             https://github.com/bettercap/bettercap                       The Swiss Army knife for 802.11, BLE and Ethernet networks reconnaissance and MITM attacks.
grpc/grpc-go                              Go   9801  1974  113             https://github.com/grpc/grpc-go                              The Go language implementation of gRPC. HTTP/2 based RPC
nats-io/nats-server                       Go   6754  685   99              https://github.com/nats-io/nats-server                       High-Performance server for NATS, the cloud native messaging system.
golang/go                                 Go   64974 9149  337             https://github.com/golang/go                                 The Go programming language
json-iterator/go                          Go   6253  533   233             https://github.com/json-iterator/go                          A high-performance 100% compatible drop-in replacement of "encoding/json"
fatedier/frp                              Go   28786 5234  280             https://github.com/fatedier/frp                              A fast reverse proxy to help you expose a local server behind a NAT or firewall to the internet.
go-redis/redis                            Go   7128  936   65              https://github.com/go-redis/redis                            Type-safe Redis client for Golang
helm/charts                               Go   10554 10908 125             https://github.com/helm/charts                               Curated applications for Kubernetes
btcsuite/btcd                             Go   3601  1410  14              https://github.com/btcsuite/btcd                             An alternative full node bitcoin implementation written in Go (golang)
hashicorp/terraform                       Go   19214 5184  105             https://github.com/hashicorp/terraform                       Terraform enables you to safely and predictably create, change, and improve infrastructure. It is an open source tool that codifies APIs into declarative configuration files that can be shared amongst team members, treated as code, edited, reviewed, and versioned.
iikira/BaiduPCS-Go                        Go   17786 2704  192             https://github.com/iikira/BaiduPCS-Go                        百度网盘客户端 - Go语言编写
astaxie/build-web-application-with-golang Go   32522 8993  104             https://github.com/astaxie/build-web-application-with-golang A golang ebook intro how to build a web with golang
mitchellh/mapstructure                    Go   2694  309   44              https://github.com/mitchellh/mapstructure                    Go library for decoding generic map values into native Go structures.
chai2010/advanced-go-programming-book     Go   9506  1453  116             https://github.com/chai2010/advanced-go-programming-book     📚 《Go语言高级编程》开源图书，涵盖CGO、Go汇编语言、RPC实现、Protobuf插件实现、Web框架实现、分布式系统等高阶主题(完稿)
kubernetes/kubernetes                     Go   59221 20761 308             https://github.com/kubernetes/kubernetes                     Production-Grade Container Scheduling and Management
rancher/rancher                           Go   12917 1474  187             https://github.com/rancher/rancher                           Complete container management platform
OWASP/Amass                               Go   2146  321   92              https://github.com/OWASP/Amass                               In-depth Attack Surface Mapping and Asset Discovery
go-gitea/gitea                            Go   16535 1854  149             https://github.com/go-gitea/gitea                            Git with a cup of tea, painless self-hosted git service
moby/moby                                 Go   55286 15954 135             https://github.com/moby/moby                                 Moby Project - a collaborative project for the container ecosystem to assemble container-based systems
Shopify/sarama                            Go   4948  890   28              https://github.com/Shopify/sarama                            Sarama is a Go library for Apache Kafka 0.8, and up.
cayleygraph/cayley                        Go   12926 1170  42              https://github.com/cayleygraph/cayley                        An open-source graph database
```

```
$ trendrepo -f text
[
    {
        "author": "T-head-Semi",
        "name": "wujian100_open",
        "href": "https://github.com/T-head-Semi/wujian100_open",
        "description": "IC design and development should be faster，simpler and more reliable",
        "language": "Verilog",
        "stars": 746,
        "forks": 230,
        "starsInPeriod": 91
    },
    {
        "author": "tinacms",
        "name": "tinacms",
        "href": "https://github.com/tinacms/tinacms",
        "description": "Tina is a site editing toolkit for modern React-based sites (Gatsby and Next.js)",
        "language": "TypeScript",
        "stars": 2078,
        "forks": 70,
        "starsInPeriod": 217
    },
    {
        "author": "TheAlgorithms",
        "name": "Java",
        "href": "https://github.com/TheAlgorithms/Java",
        "description": "All Algorithms implemented in Java",
        "language": "Java",
        "stars": 18062,
        "forks": 6793,
        "starsInPeriod": 75
    },
    ・・・
]
```
