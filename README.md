# dotato
Dotato is a dotfile manager.


![import](./example/import.gif)

![export](./example/export.gif)

(UI is not stable 😅)



## Installation
### With Go
Dotato is written in pure go.
```
> go install github.com/msisdev/dotato@latest
```


## Quickstart
Your dotato repository will look like this.
```
.
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
*         # ignore all
!.bashrc  # but .bashrc
```

Now you are ready to use dotato.

Copy files into dotato directory:
```
> dotato import group bash nux
```

Copy dotato files to remote directory:
```
> dotato export group bash nux
```
