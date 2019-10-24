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
[
    {
        "author": "tinacms",
        "name": "tinacms",
        "href": "https://github.com/tinacms/tinacms",
        "description": "Tina is a site editing toolkit for modern React-based sites (Gatsby and Next.js)",
        "language": "TypeScript",
        "stars": 1585,
        "forks": 45,
        "starsInPeriod": 379
    },
    {
        "author": "evilsocket",
        "name": "pwnagotchi",
        "href": "https://github.com/evilsocket/pwnagotchi",
        "description": "(⌐■_■) - Deep Reinforcement Learning instrumenting bettercap for WiFi pwning.",
        "language": "Python",
        "stars": 1738,
        "forks": 200,
        "starsInPeriod": 90
    },
    {
        "author": "ArduPilot",
        "name": "ardupilot",
        "href": "https://github.com/ArduPilot/ardupilot",
        "description": "ArduPlane, ArduCopter, ArduRover source",
        "language": "C++",
        "stars": 4396,
        "forks": 8403,
        "starsInPeriod": 6
    },
    {
        "author": "ripienaar",
        "name": "free-for-dev",
        "href": "https://github.com/ripienaar/free-for-dev",
        "description": "A list of SaaS, PaaS and IaaS offerings that have free tiers of interest to devops and infradev",
        "language": "HTML",
        "stars": 26005,
        "forks": 2774,
        "starsInPeriod": 540
    },
    ・・・・・
]
```

### Option

```
$ trendrepo -h
Usage of trendrepo:
  -f string
        List Format: json or text (default "json")
  -format string
        List Format: json or text (default "json")
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
$ trendrepo -l go
```

```
$ trendrepo -l go -p weekly
[
    {
        "author": "pingcap",
        "name": "tidb",
        "href": "https://github.com/pingcap/tidb",
        "description": "TiDB is an open source distributed HTAP database compatible with the MySQL protocol",
        "language": "Go",
        "stars": 21010,
        "forks": 3117,
        "starsInPeriod": 93
    },
    {
        "author": "bettercap",
        "name": "bettercap",
        "href": "https://github.com/bettercap/bettercap",
        "description": "The Swiss Army knife for 802.11, BLE and Ethernet networks reconnaissance and MITM attacks.",
        "language": "Go",
        "stars": 6117,
        "forks": 584,
        "starsInPeriod": 110
    },
    {
        "author": "danicat",
        "name": "pacgo",
        "href": "https://github.com/danicat/pacgo",
        "description": "A Pac Man clone written in Go (with emojis!)",
        "language": "Go",
        "stars": 705,
        "forks": 68,
        "starsInPeriod": 148
    },
    ・・・・
]
```

```
$ trendrepo -f text
REPO NAME                       LANG             STARS FORKS STARS IN PERIOD HREF                                               DESC
StreisandEffect/streisand       Shell            19263 1672  102             https://github.com/StreisandEffect/streisand       Streisand sets up a new server running your choice of WireGuard, OpenConnect, OpenSSH, OpenVPN, Shadowsocks, sslh, Stunnel, or a Tor bridge. It also generates custom instructions for all of these services. At the end of the run you are given an HTML file with instructions that can be shared with friends, family members, and fellow activists.
Naotonosato/Blawn               C++              616   22    201             https://github.com/Naotonosato/Blawn               Pleasant Programming Language.
vulhub/vulhub                   Shell            4487  1597  43              https://github.com/vulhub/vulhub                   Pre-Built Vulnerable Environments Based on Docker-Compose
quantopian/zipline              Python           9921  2977  110             https://github.com/quantopian/zipline              Zipline, a Pythonic Algorithmic Trading Library
electron/electron               C++              77856 10315 35              https://github.com/electron/electron               Build cross-platform desktop apps with JavaScript, HTML, and CSS
magento/magento2                PHP              7952  6959  23              https://github.com/magento/magento2                All Submissions you make to Magento Inc. ("Magento") through GitHub are subject to the following terms and conditions: (1) You grant Magento a perpetual, worldwide, non-exclusive, no charge, royalty free, irrevocable license under your applicable copyrights and patents to reproduce, prepare derivative works of, display, publically perform, subli…
HongqingCao/GitDataV            Vue              656   232   51              https://github.com/HongqingCao/GitDataV            基于Vue框架构建的github数据可视化平台
neex/phuip-fpizdam              Go               452   77    286             https://github.com/neex/phuip-fpizdam              Exploit for CVE-2019-11043
chineseGLUE/chineseGLUE         Python           362   45    75              https://github.com/chineseGLUE/chineseGLUE         Language Understanding Evaluation benchmark for Chinese: datasets, baselines, pre-trained models,corpus and leaderboard
chiphuyen/python-is-cool        Jupyter Notebook 685   128   241             https://github.com/chiphuyen/python-is-cool        Cool Python features for machine learning that I used to be too afraid to use
youfou/wxpy                     Python           11063 1947  42              https://github.com/youfou/wxpy                     微信机器人 / 可能是最优雅的微信个人号 API ✨✨
Jguer/yay                       Go               3671  143   30              https://github.com/Jguer/yay                       Yet another Yogurt - An AUR Helper written in Go
trailofbits/algo                Python           15247 1322  130             https://github.com/trailofbits/algo                Set up a personal VPN in the cloud
ageron/handson-ml2              Jupyter Notebook 3537  1172  60              https://github.com/ageron/handson-ml2              A series of Jupyter notebooks that walk you through the fundamentals of Machine Learning and Deep Learning in Python using Scikit-Learn, Keras and TensorFlow 2.
florinpop17/app-ideas           Jupyter Notebook 7277  769   60              https://github.com/florinpop17/app-ideas           A Collection of application ideas which can be used to improve your coding skills.
barry-ran/QtScrcpy              C++              868   163   60              https://github.com/barry-ran/QtScrcpy              Android实时投屏控制软件
jacobeisenstein/gt-nlp-class    TeX              3685  875   72              https://github.com/jacobeisenstein/gt-nlp-class    Course materials for Georgia Tech CS 4650 and 7650, "Natural Language"
apachecn/AiLearning             Python           20513 7247  85              https://github.com/apachecn/AiLearning             AiLearning: 机器学习 - MachineLearning - ML、深度学习 - DeepLearning - DL、自然语言处理 NLP
trimstray/nginx-admins-handbook Shell            10366 740   156             https://github.com/trimstray/nginx-admins-handbook How to improve NGINX performance, security, and other important things; @ssllabs A+ 100%, @mozilla A+ 120/100.
LisaDziuba/Awesome-Design-Tools JavaScript       17044 1123  23              https://github.com/LisaDziuba/Awesome-Design-Tools The best design tools and plugins for everything 👉
jtablesaw/tablesaw              Java             1750  355   22              https://github.com/jtablesaw/tablesaw              Java dataframe and visualization library
google-research/google-research Python           3922  644   48              https://github.com/google-research/google-research Google AI Research
2227324689/gpmall               Java             2188  810   98              https://github.com/2227324689/gpmall               【咕泡学院实战项目】-基于SpringBoot+Dubbo构建的电商平台-微服务架构、商城、电商、微服务、高并发、kafka、Elasticsearch
jet-admin/jet-bridge            Python           424   30    32              https://github.com/jet-admin/jet-bridge            Jet Bridge (Universal) for Jet Admin – API-based Admin Panel Framework for your application
cookieY/Yearning                Go               2656  980   39              https://github.com/cookieY/Yearning                受欢迎的 Mysql sql审核平台
```
