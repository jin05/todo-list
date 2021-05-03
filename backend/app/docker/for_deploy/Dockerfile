###################
# STAGE 1: builder
###################

FROM 387402613703.dkr.ecr.ap-northeast-1.amazonaws.com/alpine:3.13.5 as builder

WORKDIR /app/source

RUN apk add --update --no-cache curl \
        unzip \
        make \
        curl
ADD . .
RUN make db_install_schemalex

###################
# STAGE 2: runner
###################

FROM 387402613703.dkr.ecr.ap-northeast-1.amazonaws.com/alpine:3.13.5 as runner

WORKDIR /app

ARG pip_installer="https://bootstrap.pypa.io/pip/2.7/get-pip.py"

# Install dependent packages
RUN apk add --update --no-cache \
        make \
        mysql-client \
        python2 \
        curl \
        groff \
        git \
        gettext \
        wget \
        npm

# Install awscli
RUN curl ${pip_installer} | python && \
    pip install awscli

# Install sls
RUN npm install -g serverless && npm install --g serverless-prune-plugin

# Install binary
COPY --from=builder /usr/local/bin/schemalex /usr/local/bin/schemalex
