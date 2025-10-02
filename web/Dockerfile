# Stage 1: builder
FROM busybox:1.36 AS builder
WORKDIR /src
COPY index.html .
RUN mkdir -p /out && cp index.html /out/index.html

# Stage 2: runtime
FROM nginx:alpine AS runtime
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /out/index.html ./index.html
EXPOSE 80
