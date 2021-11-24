`编译文件生成pdf`

```shelll
latex test.tex
dvipdfmx test.dvi
```



```shell
xelatex test.tex
```



中文

设置编码格式utf-8

`\usepackage{ctex}`



```tex
% 导言区
\documentclass{article} %book,report,letter

\usepackage{ctex}
\title{\heiti 杂谈勾股定理}
\author{\kaishu 包龙}
\date{\today}

% 正文区(文稿区)
\begin{document}
	\maketitle
	Hello World!
    
    % 空一行表示换行
    Let$f(x)$be defined by the formula
    $$f(x)=3x^2+x-1$$
    % 一般用$符号包含起来称为数学模式 在其之外称为文本模式 数学模式亦可以使用双$$符号，表示另起一行写公式
    

\end{document}
```



```tex
% 导言区
\documentclass{article} %book,report,letter

\usepackage{ctex}
\newcommand\degree{^\circ}
\title{\heiti 杂谈勾股定理}
\author{\kaishu 包龙}
\date{\today}

% 正文区(文稿区)
\begin{document}
	\maketitle
	勾股定理可以用现代语言表述如下：
    
    直角三角形斜边的平方等于两腰的平方和。
    
    可以用符号语言表述为：设直角三角形$ABC$，其中$\angle C=9\degree$,则有：
    \begin{equation}
	AB^2 = BC^2 + AC^2.
    \end{equation}
    
\end{document}
```



手册查看

```shell
texdoc ctex
```

查找某个命令

```shell
texdoc lshort-zh
```



![image-20211124215431502](https://ccurj.oss-cn-beijing.aliyuncs.com/image-20211124215431502.png)