# stringer

Turns escaped raw strings into nice ones.

Usage:
```
$ echo "apiVersion: v1\nkind: Config\ncurrent-context: \"sit-sre\"" | stringer
apiVersion: v1
kind: Config
current-context: "sit-sre"
```
