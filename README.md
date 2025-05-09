# dotato
🥔dotato is a lightweight dotfile manager.


![import](./example/import.gif)

![export](./example/export.gif)



## Introduction
🥔dotato is simple.
- Select dotfiles with config file
- Move dotfiles to your directory

Config files are clear:
- `dotato.yaml` - declare original path of dotfiles
- `.dotatoignore` - select files that you are interested

Choose move behavior:
- file mode - copy dotfiles into your directory (like snapshot)
- link mode - move dotfiles into your directory and leave symlink instead. (like [stow](https://www.gnu.org/software/stow/))



## Installation
### With Go
dotato is written in pure go. If you have go, it is easy:
```cmd
> go install github.com/msisdev/dotato@latest
```


## Tutorial (file mode)
Your dotato directory will look like this.
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

Configure `bash/.dotatoignore` to tell which files to ignore/grab.
```gitignore
#.dotatoignore
*         # ignore all
!.bashrc  # grab .bashrc
```

Copy files into dotato directory:
```cmd
> dotato import group bash nux
```

Copy dotato files to remote directory:
```cmd
> dotato export group bash nux
```
