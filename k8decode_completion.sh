#!/bin/bash
__k8decode_parse_get()
{
local k8decode_output out
if kubectl_output=$(k8decode --no-headers "$1" 2>/dev/null); then
	out=($(echo "$(kubectl get secret)" | awk '{print $1}'))
	COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
fi
}
complete -F __k8decode_parse_get k8decode