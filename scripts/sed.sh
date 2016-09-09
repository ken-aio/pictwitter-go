#!/bin/sh

if [ $# -ne 1 ]; then
  echo "引数にプロジェクト名を入力してください"
  echo "Usage: sh $0 your_pjname"
  exit 1
fi

grep -rl "pictwitter-go" ./ | xargs perl -i -pe "s/pictwitter-go/$1/g"
