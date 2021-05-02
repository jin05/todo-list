###################
# STAGE 1: builder
###################

FROM 387402613703.dkr.ecr.ap-northeast-1.amazonaws.com/alpine:3.13.5 as builder

WORKDIR /app/source

RUN apk update && apk add --no-cache go \
        git \
        binutils-gold \
        curl \
        g++ \
        gcc \
        gnupg \
        libgcc \
        linux-headers \
        make \
        python
ADD . .
RUN go build && strip todo-list

###################
# STAGE 2: runner
###################

FROM 387402613703.dkr.ecr.ap-northeast-1.amazonaws.com/alpine:3.13.5 as runner

WORKDIR /app

RUN apk update && apk add --no-cache libc6-compat ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /app/source/todo-list /app/
COPY --from=builder /app/source/cmd/start /app/cmd/

# run it
ENTRYPOINT ["/app/cmd/start"]
