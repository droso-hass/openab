# Nabaztag V2

Code for the Nabaztag:tag

Communication is done over a TCP Socket: see [PROTOCOL.md](PROTOCOL.md).

When the nabaztag connects to the server, the servers uses its ip to connect to the nabaztag on port 5000.

Audio data is transmitted over UDP on port 4000 for nabaztag and server (recording and playback audio).

## Bootcode

Code for the nominal bootcode for the nabaztag:tag is located in `/tagtag/src`

### Compiling

Make sure that you have cloned this repo with the `--recurse-submodules` option.

Go to the repo root and run `make v2-fw`. To run the simulator use `make v2-simu`.

### Syntax Highlighting

There is a basic syntax highlighting vscode extension in `syntax_highlighting`.

To use it, make a symlink to your extensions dir `ln -s $(pwd)/tagtag/syntax_highlighting/ ~/.vscode/extensions/metal`.

## Resources

[Nabaztag:tag Creator Page](https://www.sylvain-huet.com/thema.php?lang=fr#nabv2)

[Mindscape Sources](http://code.google.com/p/nabaztag-source-code/source/browse/#svn%2Ftrunk%2FSources)
