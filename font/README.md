# README

`@umijs/max` 模板项目，更多功能参考 [Umi Max 简介](https://umijs.org/docs/max/introduce)

# 准备
- pnpm安装
```bash
curl -fsSL https://get.pnpm.io/install.sh | sh -
$ pnpm -v
7.3.0
```
- nvm安装
```bash
$ curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
$ nvm -v
0.39.1
```
- node安装
```bash
$ nvm install 16
$ nvm use 16
$ node -v
v16.10.0
```

# 启动
```bash
pnpm dev
```

# 代理
找到`config/proxy.ts`文件，将`target`改成所要代理的IP

