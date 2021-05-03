####################################
# docker
####################################
docker_build_deploy_image:
	docker build ./ -t todo-list-deploy -f docker/for_deploy/Dockerfile
