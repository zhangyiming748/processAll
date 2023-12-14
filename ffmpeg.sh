#!/bin/zsh
echo 在不符合golang版本的系统上临时使用批量ffmpeg命令的方法
# 定义源文件夹和目标文件夹
source_dir=/data/folder
target_dir=/data/folder/h265
# 判断文件夹是否存在
if [ ! -d $target_dir ]; then
    # 文件夹不存在,创建文件夹
    mkdir -p "$target_dir"
fi
# 遍历源文件夹中的所有文件
for file in "$source_dir"/*; do
    # 获取文件名
    filename=$(basename "$file")
    # 构造源文件的完整路径
    source_file="$source_dir/$filename"
    echo '源文件: $source_file'
    # 构造目标文件的完整路径
    target_file="$target_dir/$filename"
    echo '目标文件: $target_file'
    ffmpeg -i "$source_file" -c:v libx265 -c:a aac -ac 1 -tag:v hvc1 "$target_file"
done
