<div align="center">
  
  <h1 align="center">K8decode: decode kubernetes secrets</h1>

  [![test][test-image]][test-url]
  [![coverage][cov-image]][cov-url]
  
  [test-image]: https://github.com/sullivtr/k8decode/workflows/test/badge.svg
  [test-url]: https://github.com/sullivtr/k8decode/actions?query=workflow%3Atest
  [cov-image]: https://codecov.io/gh/sullivtr/k8decode/branch/master/graph/badge.svg?token=6HISPGX35J
  [cov-url]: https://github.com/sullivtr/k8decode/actions?query=workflow%3Atest

</div>

<img src="secretdecode.gif" width="1000"/>
<!-- ![Demo](secretdecode.gif) -->

```bash
k8decode {secret-name} [-n] {namespace}
```
The namespace flag's default value is `default`. Use `-n` to specify an alternate namespace for the secret. 

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
