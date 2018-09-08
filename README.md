# gitcd

Quickly navigate to your GitHub repositories.

## Usage

### 1) Install `gitcd`.

### 2) Add this function to your bash profile (`~/.bashrc` or `~/.bash_profile`):

```bash
[ -f ~/.bashrc ] && echo 'gcd() { gitcd "$@" && cd `gitcd "$@"`; }' >> ~/.bashrc && . ~/.bashrc
[ -f ~/.bash_profile ] && echo 'gcd() { gitcd "$@" && cd `gitcd "$@"`; }' >> ~/.bash_profile && . ~/.bash_profile
```

### 3) Use `gcd` to navigate to a repository.

```bash
gcd coollog/gitcd
```

`gitcd` clones the repository first if it does not exist.

## Configuration

Set `GITCD_HOME` to change the root directory for the cloned repositories.
