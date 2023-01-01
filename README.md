# SmartSuffix
Add ext to file those have no ext

### Usage
./ss [--default] [working directory]  
- --default: Default `ext` when failed to detect file type automatically
- working directory: Directory to scanning files, default is cwd

### Examples
```bash
./ss
./ss .
./ss --default=bin ~/Downloads
```

### Dependencies
- https://github.com/jedib0t/go-pretty
- https://github.com/h2non/filetype
