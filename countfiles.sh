a="$(ls -1R | egrep -c '^-|^d')" 
((a++)) 
echo $a