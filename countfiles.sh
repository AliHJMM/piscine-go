# a="$(ls -1R | egrep -c '^-|^d')" 
# ((a++)) 
# echo $a
echo $(find . -type f | grep -v .git | wc -l)