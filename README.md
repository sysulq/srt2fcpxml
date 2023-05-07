# srt2fcpxml

Convert srt subtitle file to final cut pro subtitle file(fcpxml)

This software uses final cut pro X 10.4.6 version fcpxml file as template development, if there is any problem, please upgrade to the corresponding version.

srt 字幕文件转为final cut pro 字幕文件(fcpxml)

本软件使用 final cut pro X 10.4.6 版本的 fcpxml 文件作为模版开发，如果有问题请升级到对应版本

## Install

```bash
go install github.com/hnlq715/srt2fcpxml
```

复制`自适应背景.moti`文件到`~/Motion Templates/Titles`目录下

## Use

```bash
$ ./srt2fcpxml
Usage of srt2fcpxml:
  -fd string
        frame rate currently supported 23.98、24、25、29.97、30、50、59.94、60 (default "25")
  -height int
        high resolution default 1080 (default 1080)
  -logger-app-name string
        the logger's app name
  -logger-filename string
        the logger's filename
  -srt string
        srt subtitle file
  -v    if true, then log level is debug
  -width int
        width resolution default 1920 (default 1920)
```

## Execution (执行)

```bash
srt2fcpxml -srt /tmp/test.srt
```

the `fcpxml` file named with srt file name will be generated automatically in the directory of srt file.

在 srt 文件的目录中会自动生成以srt文件名命名的`fcpxml`文件。
