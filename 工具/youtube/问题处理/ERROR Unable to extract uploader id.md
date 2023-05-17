参考文章

1. [Error: Unable to extract uploader id - Youtube, Discord.py](https://stackoverflow.com/questions/75495800/error-unable-to-extract-uploader-id-youtube-discord-py)

```
# youtube-dl --no-check-certificate 'https://www.youtube.com/watch?v=xxxxxx'
[youtube] Downloading login page
[youtube] Looking up account info
WARNING: Unable to look up account info: HTTP Error 400: Bad Request
[youtube] xxxxxx: Downloading webpage
WARNING: Unable to download webpage: HTTP Error 429: Too Many Requests
[youtube] xxxxxx: Downloading API JSON
[youtube] xxxxxx: Downloading MPD manifest
ERROR: Unable to extract uploader id; please report this issue on https://yt-dl.org/bug . Make sure you are using the latest version; see  https://yt-dl.org/update  on how to update. Be sure to call youtube-dl with the --verbose flag and include its complete output.
```
