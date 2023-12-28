# Overview
---
CLI tool to easily get differential parts of files built with buffio, os and flags libraries.

# Usage
---
Simply pass two text files as parameters like below (flags are optional):
```
godiffer [<flag>] <input_file1> <input_file2>
```
There are 4 options to choose from:
```
  -d	a difference set (default)
  -i	an intersection set
  -s	a symetric difference set
  -u	an union set
```
