# dotato
A dotfile or symlink manager.

## In strategy
| state   | symlink mode | file mode |
|---------|--------------|-----------|
| symlink | 원본파일 추적, 새로운 링크 | 원본파일 추적, 링크는 그대로 |
| file    | stow 처럼     | 복사해서 가져오기 |


## Out strategy
| state   | symlink mode | file mode |
|---------|--------------|-----------|
| symlink | state 확인, overwrite | overwrite |
| file    | overwrite   | state확인, 덮어쓰기 |

## Cases
```
[import file]
dot     dtt

file    not exists
file    file, diff
file    file, same
link    not exists
link    file, diff
link    file, same


[import link]
dot     dtt

file    not exists
file    file, diff
file    file, diff
link    not exists
link    file, wrong
link    file, correct


[export file]
dot         dtt

not exists  file
exists      file


[export link]
dot         dtt

not exists  file
exists      file
```
