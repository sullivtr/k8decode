# Auto decode kubernetes secrets 

If you are tired of piping your secrets manually through a base64 decoder, this tool is for you.
Download the binary from the /binaries dir (or from the Releases page) and place it in your path.

```bash
k8decode my-super-secret-secret
```

For `tab` autocompletion, download the `k8decode_completion.sh` bash completion script included in this repo

```bash
k8decode [tab]

# OR

k8decode secretsubstring[tab]
```

If you use ZSH, you can use the autocomplete script by doing the following

```bash
# RUN
autoload -U +X bashcompinit && bashcompinit
source /path/to/script
```
