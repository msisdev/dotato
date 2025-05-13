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
- **file mode**: copy dotfiles into your backup directory (like snapshot)
- **link mode**: move dotfiles into your backup directory and leave symlink instead. (like [stow](https://www.gnu.org/software/stow/))

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
### With binaries
Download binary file in the dotato release page [here](https://github.com/msisdev/dotato/releases).

### With Go
If you have [go](https://go.dev/dl/), it is easy:
```console
go install github.com/msisdev/dotato@latest
```
Some systems may require C library.

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

Write `dotato.yaml`.
```yaml
# dotato.yaml
version: v1

mode: file

groups:
  bash:        # same name of your group directory
    nux: "~"   # write directory of your dotfile
```

Write `bash/.dotatoignore`. It applies to both import/export.
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
│   ├── 🚚.bashrc        # dotato will copy this to ~/.bashrc
│   └── 📄.dotatoignore
└── ⚙️dotato.yaml
```

## Tips
[CLI](https://github.com/msisdev/dotato/wiki/Commands)
- Use flag `-h` to read hints.

[`dotato.yaml`](https://github.com/msisdev/dotato/wiki/Configuration#dotatoyaml)
- There is another entity 'plan' — select multiple groups.
- Create groups as many as you like.
- Create duplicate groups to maintain different versions.
- You can save many directories in one group — manage different machines.

[`.dotatoignore`](https://github.com/msisdev/dotato/wiki/Configuration#dotatoignore)
- It works same with [gitignore](https://git-scm.com/docs/gitignore).
- You can define global rule - one ignore rule applied to all groups.
- Nest many dotatoignore files under group directory.
- Remember it is applied on both import/export command.





## Advanced
Currently dotato doesn't provide advanced features.
- Templating
- Install script
- Password encryption

You can...
1. Create issue and wait for support
2. Use another tools like [chezmoi](https://www.chezmoi.io/), [etc](https://github.com/topics/dotfiles-manager)
3. Use dotato API — [engine](https://pkg.go.dev/github.com/msisdev/dotato/pkg/engine) — if you love dotato and ready to go down the rabbit hole.

If you decided to use dotato API, golang standard libraries will help.
- Use [template](https://pkg.go.dev/text/template) to change file content.
- Use [os.Exec](https://pkg.go.dev/os/exec) to run external commands
- Use [crypto](https://pkg.go.dev/crypto) to encrypt/decrypt file content.
