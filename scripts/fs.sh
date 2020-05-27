#!/usr/bin/env bash

set -e

alpine_x86_releases="http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64"
mini_root_fs="alpine-minirootfs-3.11.6-x86_64.tar.gz"

if [[ -d minirootfs && -f $mini_root_fs && -f $miniroot_fs.sha256 ]]
then
  echo 'minirootfs is already downloaded, valid, and configured'
fi

root_url="${alpine_x86_releases}/${mini_root_fs}"

wget "${root_url}"
wget "${root_url}.sha256"

given_sha_sum=$(cat $mini_root_fs.sha256)
inferred_sha_sum=$(shasum -a 256 $mini_root_fs)

if [[ $given_sha_sum != $inferred_sha_sum ]]
then
  echo 'sah256 sums do not match'

  rm $mini_root_fs*

  exit 1
else
  mkdir -p minirootfs

  tar -C minirootfs -xzf $mini_root_fs
fi
