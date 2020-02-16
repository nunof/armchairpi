# *** AARCH64 / ARM64 **
#export CC=/usr/bin/aarch64-linux-gnu-gcc
#export CGO_CFLAGS="--sysroot=/home/nunof/Downloads/aaaaa/sysroot-glibc-linaro-2.25-2019.02-aarch64-linux-gnu"
#export CGO_LDFLAGS="--sysroot=/home/nunof/Downloads/aaaaa/sysroot-glibc-linaro-2.25-2019.02-aarch64-linux-gnu"

# *** ARM ***
export CC=/usr/bin/arm-linux-gnu-gcc
export CGO_CFLAGS="--sysroot=/home/nunof/Downloads/aaaaa/sysroot-glibc-linaro-2.25-2019.02-arm-linux-gnueabihf"
export CGO_LDFLAGS="--sysroot=/home/nunof/Downloads/aaaaa/sysroot-glibc-linaro-2.25-2019.02-arm-linux-gnueabihf"

env \
PKG_CONFIG_DIR="" \
PKG_CONFIG_LIBDIR="/home/nunof/Downloads/aaaaa/sysroot-glibc-linaro-2.25-2019.02-arm-linux-gnueabihf/usr/lib/pkgconfig:/home/nunof/Downloads/aaaaa/sysroot-glibc-linaro-2.25-2019.02-arm-linux-gnueabihf/usr/share/pkgconfig:/home/nunof/Downloads/aaaaa/sysroot-glibc-linaro-2.25-2019.02-aarch64-linux-gnu/usr/lib:/usr/lib64/pkgconfig" \
PKG_CONFIG_SYSROOT_DIR="/home/nunof/Downloads/aaaaa/sysroot-glibc-linaro-2.25-2019.02-arm-linux-gnueabihf" \
GOARCH=arm \
GOARM=7 \
GOOS=linux \
CGO_ENABLED=1 \
go build -x -o armchairpi