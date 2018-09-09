# gitcd

Quickly navigate to your GitHub repositories.

## Examples

Without `gitcd`, you may need to manage multiple each clone manually:

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
gcd coollog/gitcd
# Some days later.
gcd foo/bar
...
```

## Usage

### 1) Install `gitcd`.

*TODO: Release binaries.*

For now, you have to `go get -u github.com/coollog/gitcd`. `gitcd` will be at `$GOPATH/bin/gitcd` (default `~/go/bin/gitcd`). Make sure `gitcd` is on your `PATH`.

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

Set `GITCD_HOME` to change the root directory for the cloned repositories. By default, `gitcd` uses `~/gitcd`.
