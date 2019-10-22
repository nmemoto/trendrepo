# trendrepo

Github Trending CLI (https://github.com/trending)

## How To Use

### Build

```
$ git clone git@github.com:nmemoto/trendrepo.git
$ cd trendrepo
$ go build
```

### Run

```
$ ./trendrepo
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
$ ./trendrepo
Usage of ./trendrepo:
  -l string
        Programming Language
  -p string
        Date Range (default "today")
```

### example

```
$ ./trendrepo -l go
```

```
$ ./trendrepo -l go -p week
```
