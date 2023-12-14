# ffmpeg

```shell
ffmpeg -i input.mp4 -ss 00:30:31.827 -to 00:44:58.194 -c:v copy -c:a copy -filter:a volume=3.0 -map_chapters -1 output.mp4
```