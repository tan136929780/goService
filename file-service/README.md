docker build --rm -t vfile .

测试环境配置:
docker run -d -p 8200:8200 -v /Users/tanxianchen/coding/docker/www/vfile:/vfile --name vfile vfile

项目首次运行时启动后需要执行（执行模型创建） /preload/init
