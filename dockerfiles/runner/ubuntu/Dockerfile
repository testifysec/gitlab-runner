FROM ubuntu:20.04

ARG DOCKER_MACHINE_VERSION=0.16.2
ARG DOCKER_MACHINE_ARCH=x86_64

ARG DUMB_INIT_VERSION=1.2.5
ARG DUMB_INIT_ARCH=x86_64

ARG GIT_LFS_ARCH=amd64
ARG GIT_LFS_VERSION=3.2.0

ENV DEBIAN_FRONTEND=noninteractive
# hadolint ignore=DL3008
RUN apt-get update -y && \
    apt-get install -y --no-install-recommends \
        apt-transport-https \
        ca-certificates \
        curl \
        git \
        wget \
        tzdata \
        openssh-client \
    && rm -rf /var/lib/apt/lists/*


#COPY ./bin/gitlab-runner /usr/bin/

#install docker-machine
RUN curl -L -o /usr/bin/docker-machine "https://github.com/docker/machine/releases/download/v${DOCKER_MACHINE_VERSION}/docker-machine-Linux-${DOCKER_MACHINE_ARCH}"

#install dumb-init
RUN curl -L -o /usr/bin/dumb-init "https://github.com/Yelp/dumb-init/releases/download/v${DUMB_INIT_VERSION}/dumb-init_${DUMB_INIT_VERSION}_${DUMB_INIT_ARCH}"

RUN curl -L "https://github.com/git-lfs/git-lfs/releases/download/v${GIT_LFS_VERSION}/git-lfs-linux-${GIT_LFS_ARCH}-v${GIT_LFS_VERSION}.tar.gz" | tar -xz -C /tmp \
    && mv /tmp/git-lfs-${GIT_LFS_VERSION}/git-lfs /usr/bin/ \
    && rm -rf /tmp/*

RUN chmod +x /usr/bin/docker-machine
RUN docker-machine version

RUN chmod +x /usr/bin/dumb-init
RUN dumb-init --version


RUN chmod +x /usr/bin/git-lfs

RUN git lfs version
RUN git-lfs install --skip-repo


COPY ./dockerfiles/runner/ubuntu/entrypoint /
RUN chmod 777 /entrypoint

STOPSIGNAL SIGQUIT
VOLUME ["/etc/gitlab-runner", "/home/gitlab-runner"]
ENTRYPOINT ["/usr/bin/dumb-init", "/usr/bin/entrypoint"]
CMD ["run", "--user=gitlab-runner", "--working-directory=/home/gitlab-runner"]