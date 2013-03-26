
splitline -l 500 -i infile.txt splits/chunk | xargs -I {} mean -i {} | sed -ne 'w file1' -e 'n;w file2' && paste file1 file2 | awk '{printf "%.15g\n"  $1 - $2}' > res.txt

awk 'NR%2==1 {print >> "file1"; next} {print >> "file2"}' infile 

paste file1 file2 | awk '{print $1 - $2}'

 nl -v 0 -i 10 -w 7 -n ln -s \,  file1

 