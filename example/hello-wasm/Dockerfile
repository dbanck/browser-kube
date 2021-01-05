FROM rustlang/rust:nightly as builder
WORKDIR /rust
RUN apt-get update && \
    apt-get install -y curl gcc python3

RUN curl -L https://sh.rustup.rs | sh -s -- -y --default-toolchain=nightly
ENV PATH=$PATH:~/.cargo/bin

RUN rustup update && \
    rustup target add wasm32-unknown-unknown --toolchain nightly && \
    cargo install wasm-pack

ADD . .
RUN wasm-pack build --target web --out-name wasm --out-dir ./static

FROM scratch
COPY --from=builder /rust/static /
