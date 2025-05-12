<div align="center">

# 🥔 dotato
dotato is a lightweight dotfile manager.

<sup>[What is dotfile manager?](https://github.com/msisdev/dotato/wiki/Background)</sup>

<img src="./example/import.gif" alt="import" width="75%">

Import dotfiles

<img src="./example/export.gif" alt="export" width="75%">

Export dotfiles

</div>


## Introduction

🥔 dotato is simple.
- Write config file.
- Run dotato.

✏️ Config files are clear.
- Define directories in `dotato.yaml`
- Filter unnecessary files with `.dotatoignore`

🚚 Choose your mode.
- **file mode**: copy dotfiles into your directory (like snapshot)
- **link mode**: move dotfiles into your directory and leave symlink instead. (like [stow](https://www.gnu.org/software/stow/))

## OS Support
<table>
  <tr>
    <th>OS</th>
    <th>Support</th>
  </tr>
  <tr>
    <td>Linux</td>
    <td>✅</td>
  </tr>
  <tr>
    <td>MacOS</td>
    <td>✅</td>
  </tr>
  <tr>
    <td>Windows</td>
    <td>
      <p>File mode: ⚠️</p>
      <ul>
        <li>It works with local drive. e.g. <code>C:\</code></li>
        <li>It doesn't work with network path. e.g. <code>\\wsl.localhost\</code></li>
      </ul>
      <p>Link mode: ❌</p>
      <ul>
        <li>Any command that deletes symlink will fail.</li>
      </ul>
    </td>
  </tr>
  
</table>



## Installation
### With Go
dotato is written in pure go. If you have [go](https://go.dev/dl/), it is easy:
```console
go install github.com/msisdev/dotato@latest
```

And make sure you have `~/go/bin` in `PATH` env var.



## Tutorial
Let's copy `~/.bashrc` file into your backup directory.

Prepare your backup directory like this.
```
📁
├── 📁bash
│   └── 📄.dotatoignore
└── ⚙️dotato.yaml
```

Write `dotato.yaml` to tell dotato where your dotfiles are.
```yaml
# dotato.yaml
version: 1.0.0

mode: file

groups:
  bash:        # same name of your group directory
    nux: "~"   # write directory of your dotfile
```

Write `bash/.dotatoignore` to tell which files to ignore/grab.
```gitignore
#bash/.dotatoignore
*         # ignore all
!.bashrc  # but include .bashrc
```

Copy files into backup directory:
```console
dotato import group bash nux
```
```
📁
├── 📁bash
│   ├── ✨.bashrc        # dotato created this
│   └── 📄.dotatoignore
└── ⚙️dotato.yaml
```

Copy dotato files back to their original place:
```console
dotato export group bash nux
```
```
📁
├── 📁bash
│   ├── ✨.bashrc        # dotato will copy this to ~/.bashrc
│   └── 📄.dotatoignore
└── ⚙️dotato.yaml
```

## Tips
[`dotato.yaml`](https://github.com/msisdev/dotato/wiki/Configuration#dotatoyaml)
- There is another entity 'plan' - for selecting multiple groups.
- Create groups as many as you like.
- Create duplicate groups to maintain different versions.
- You can have many directories in one group - address with many machines.

[`.dotatoignore`](https://github.com/msisdev/dotato/wiki/Configuration#dotatoignore)
- It works same with gitignore.
- You can define global rule - one ignore rule applied to all groups.
- Nest many dotatoignore files under group directory.
- Remember it is applied on both imprt/export command.
