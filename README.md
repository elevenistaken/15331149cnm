# 15331149cnm
服务计算用go语言开发selpg命令

使用方法： -s Start from Page . 
-e End to Page .  
-l [options]Specify the number of line per page.Default is72.format:-l=number  
-f [options]Specify that the pages are sperated by \f.  
[filename] [options]Read input from the file.  
If no file specified, selpg will read input from stdin. Control-D to end.  
   
已达成的形参的输入事例：  
./selpg -s=1 -e=1 in.txt  
./selpg -s=1 -e=1 < in.txt  
./selpg -s=1 -e=1 -l=2 in.txt    
./selpg -s=1 -e=1 -f in.txt  
./selpg -s=1 -e=1 in >out.txt     
