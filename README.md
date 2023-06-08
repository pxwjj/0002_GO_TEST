添加 SSH 可以让你更方便高效的与 GitHub 交互,以下是详细步骤:

1. 生成 SSH key

```bash
ssh-keygen -t rsa -C "your_email@example.com"
```

将生成的两个文件：id_rsa和id_rsa.pub 保存到 `~/.ssh/`。

之后会要求你输入密码,密码可以为空回车。

2. 查看公钥

```bash
cat ~/.ssh/id_rsa.pub
```

会显示一个类似下面的文本:

```
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD3F6tyPEFEzV0LX3X8BsXdMsQz1x2cEikKDEY0aIj41qgxMCP/iteneqXSIFZBp5vizPvaoIR3Um9xK7PGoW8giupGn+EPuxIA4cDM4vzOqOkiMPhz5XK0whEjkVzTo4+S0puvDZuwIsdiW9mxhJc7tgBNL0cYlWSYVkz4G/fslNfRPW5mYAM49f4fhtxPb5ok4Q2Lg9dPKVHO/Bgeu5woMc7RY0p1ej6D4CKFE6lymSDJpW0YHX/wqE9+cfEauh7xZcG0q9t2ta6F6fmX0agvsS6T2EJZGeXSTGkwGAolH+lIMYn0gyaazbPN3aVmvUJgWysqA/nLZV5XhwpJdsa6XUlcC email@example.com
```

3. 添加公钥到 GitHub

登录GitHub -> Settings -> SSH and GPG keys -> New SSH key -> Title 任意填写一个名称 -> 将上一步生成的公钥内容粘贴到 Key 框中 -> Add SSH key

4. 测试 SSH 连接

```bash
ssh -T git@github.com
```  

会提示:

```
Hi username! You've successfully authenticated, but GitHub does not provide shell access.
```

表示 SSH Successfully!


设置完毕后,每次 `git clone` `git push` 就会用 SSH 方法推送数据到 GitHub 了，不需要输入用户名和密码。

