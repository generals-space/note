# git token认证

参考文章

1. [Token authentication requirements for Git operations](https://github.blog/2020-12-15-token-authentication-requirements-for-git-operations/)
    - 废止密码认证, 使用token认证
2. [Creating a personal access token](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token)
    - 生成自己的token串以及使用方法
    - `gh auth login`
3. [Updating credentials from the macOS Keychain](https://docs.github.com/en/get-started/getting-started-with-git/updating-credentials-from-the-macos-keychain)
    - 清除本地缓存的密码
4. [gh安装方法](https://github.com/cli/cli#installation)

token类似于加密过的密码, 进行git操作时, 将需要输入用户名&token, 而不再是用户名&密码.

但是token比密码强大的地方在于, 我们可以通过token进行权限控制, 且定义过期时间.

按照参考文章2生成新的token后, 需要按照参考文章3中的步骤, 将本地缓存的密码认证清除.

注意: 这个token必须自行保存, 之后再也无法重新获取其内容了.

注意: linux环境下使用`gh auth login`登录, 可以只输入token, 不再需要用户名.
