# give
Command line GitHub tools

# Usage

For example, if you see issues on GitHub

``` sh
$ give issue
#1   bug this is bug.
#2   enhancement this enhancement idea is ...
#3   bug bug report 2
.   .   ...
.   .   ...
.   .   ...
```

# options

## issue
If you get information about GitHub issue.

- num
The num option is restricted number of output

``` sh
$ give issue --list 3
#1   bug this is bug
#2   enhancement this enhancement idea is ...
#3   bug bug report 2
```

- show
The show option displays the details of the specified number issue.

``` sh
$ give issue --show 2
#2  Updated: 2019/05/09/    issue title
Labels:
Issue URL: https://github.com/homedm/give/issues/2
this issue is ...
```

- add
The add option add issue to repository.
Use this option, open the git editor and input issue body.
open the text editor which you use at `git commit`, typed issue body.

``` sh
$ give issue --add "issue title"
```

- close
The close option close issue.

``` sh
$ give issue --list
#1   bug this is bug
#2   enhancement this enhancement idea is ...
#3   bug bug report 2

$ give issue --close 3

$ give issue --list
#1   bug this is bug
#2   enhancement this enhancement idea is ...
```
