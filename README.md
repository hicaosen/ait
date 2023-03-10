# ait
‘ait’  is a command line tool for managing git repository URLs

## Commands

### ait init

Create a `.ait/git-list.json` file in the user's home directory and initialize its content as `{}`.

### ait add repo url

Add `repo` and `url` as key-value pairs to `git-list.json`.

### ait del repo

Delete `repo` from `git-list.json`.

### ait ls

List all items in `git-list.json`, and `ait ls repo` can be used to view the corresponding URL of `repo`.

### ait clone repo

Clone the corresponding repository of `repo` to the current directory.

### ait clone all

If you input `y`, all the projects in `git-list.json` will be cloned to the current directory, otherwise they will not be cloned.

### ait update repo url

Update the URL corresponding to `repo` to `url`.

### ait search repo

Search `repo` in `git-list.json`, if found, display the corresponding URL of `repo`, otherwise prompt not found.

### ait help

Display help information.