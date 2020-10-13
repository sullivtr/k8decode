#!/bin/bash
__k8decode_parse_get()
{
local out
	out=($(echo "$(kubectl get secret --all-namespaces)" | awk '{print $2}'))
	COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
}
complete -F __k8decode_parse_get k8decode