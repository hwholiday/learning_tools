### 安装 docker

### 安装 docker-compose

### 安装 harbor 镜像仓库 需要配置https证书

```
#安装 acme.sh 用来生成证书
curl https://get.acme.sh | sh 
source ~/.bashrc
acme.sh --version 测试是否安装成功
# 使用 acme.sh dns api 来生成证书 ,这个证书一定要能校验有效,不然 kaniko 会上传不上去

wget https://github.com/goharbor/harbor/releases/download/v2.1.2/harbor-online-installer-v2.1.2.tgz

tar -zxvf harbor-online-installer-v2.1.2.tgz

./install.sh
```

### 安装 gitlab

```
sudo docker run -d -h 主机IP -p 443:443 -p 80:80  --name gitlab --restart always --volume /srv/gitlab/config:/etc/gitlab --volume /srv/gitlab/logs:/var/log/gitlab --volume /srv/gitlab/data:/var/opt/gitlab gitlab/gitlab-ce
```

### 安装 gitlab-runner

```
docker run -d --name gitlab-runner --restart always -v /srv/runner/config:/etc/gitlab-runner -v /run/docker.sock:/var/run/docker.sock gitlab/gitlab-runner
```

### 测试

```
 git tag -a "1.1.1" -m "test k8s"
 git push origin 1.1.19
```

### 安装Rancher k8s集群

```
docker run --privileged -d --restart=unless-stopped -p 80:80 -p 443:443 rancher/rancher
```
