# gitcd

Quickly navigate to your GitHub repositories.

[![Gitter chat](https://badges.gitter.im/coollog/gitcd.png)](https://gitter.im/coollog/gitcd)

## Examples

Without `gitcd`, you may need to manage multiple GitHub repo clones manually:

```bash
mkdir -p ~/github/coollog
cd ~/github/coollog
git clone https://github.com/coollog/gitcd.git
cd gitcd
# Now make some commits.
# Time to work on another repo.
mkdir -p ~/github/foo
git clone -C ~/github/foo/bar https://github.com/foo/bar.git
cd ~/github/foo/bar
# Make some edits.
# Now time to go back to coollog/gitcd. Where is it again? Oh, right.
cd ~/github/coollog/gitcd
# Some days later. Did I clone foo/bar already?
ls ~/github/foo/bar
# Ah okay, time to go work on that.
cd ~/github/foo/bar
...
```

With `gitcd`, this becomes just:

```bash
gcd coollog/gitcd # Clones https://github.com/coollog/gitcd.git
# Make some commits.
gcd foo/bar       # Clones https://github.com/foo/bar.git
# Make some commits.
gcd gitcd
# Some days later.
gcd bar
...
```

## Usage

### 1) Install `gitcd`.

#### Linux

```bash
curl -Lo gitcd https://storage.googleapis.com/gitcd/gitcd-linux-amd64 && \
    chmod +x gitcd && sudo mv gitcd /usr/local/bin
```

#### macOS

```bash
curl -Lo gitcd https://storage.googleapis.com/gitcd/gitcd-darwin-amd64 && \
    chmod +x gitcd && sudo mv gitcd /usr/local/bin
```

#### Windows

Download the latest Windows build: https://storage.googleapis.com/gitcd/gitcd-windows-amd64.exe

#### Build from source

```bash
go get -u github.com/coollog/gitcd
# `gitcd` will be at `$GOPATH/bin/gitcd`
```

### 2) Add `gcd` to your bash profile (`~/.bashrc` or `~/.bash_profile`):

This adds `gcd` as a `bash` function.

```bash
[ -f ~/.bashrc ] && echo 'gcd() { GITCD_GCD=1 gitcd "$@" && cd `gitcd "$@"`; }' >> ~/.bashrc && . ~/.bashrc
[ -f ~/.bash_profile ] && echo 'gcd() { GITCD_GCD=1 gitcd "$@" && cd `gitcd "$@"`; }' >> ~/.bash_profile && . ~/.bash_profile
```

### 3) Use `gcd` to navigate to a repository.

```bash
gcd coollog/gitcd
```

`gitcd` clones the repository first if it does not exist.

## Configuration

Set `GITCD_HOME` to change the root directory for the cloned repositories. By default, `gitcd` uses `~/gitcd`.

## How it works

```bash
# These all navigate to the directory for the cloned repo, cloning the repo if necessary.
gcd https://github.com/coollog/gitcd.git
gcd http://github.com/coollog/gitcd.git
gcd https://github.com/coollog/gitcd
gcd git@github.com:coollog/gitcd.git
gcd github.com/coollog/gitcd
gcd coollog/gitcd.git
gcd coollog/gitcd
gcd gitcd # If you have used repos under coollog/ before.
```

When the name is ambiguous (just the repo name like `gitcd` rather than `coollog/gitcd`), `gitcd` tries to find the name under owners in the order in which they were last used. For example, if `gitcd` had used `foo/`, `bar/`, and `cat/` (in that order), `gcd dog` would try to find `dog` in `cat/dog`, then `bar/dog`, then `foo/dog`. 
