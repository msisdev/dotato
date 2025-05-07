# dotato
Dotato is a dotfile manager.


## Installation
### With Go
Dotato is written in pure go.
```
> go install github.com/msisdev/dotato@latest
```


## Quickstart
Your dotato repository should look like this.
```
my-dtt
├── bash
│   └── .dotatoignore
├── .dotatoignore
└── dotato.yaml
```

Configure `dotato.yaml` to tell dotato where to order/deliver files from/to.
```yaml
# dotato.yaml
version: 1.0.0

mode: file

plans:
  all:  # empty plan means all groups

groups:
  bash:
    nux: "~"
```

Configure `bash/.dotatoignore` to tell which files to ignore or grab.
```
#.dotatoignore
*
!.bashrc
```

Now you can order/deliver your files!
```
> dotato import group all nux
> dotato export group all nux
```
