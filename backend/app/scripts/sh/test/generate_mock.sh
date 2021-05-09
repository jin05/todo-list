#!/bin/sh

WORK_DIR=$(pwd); cd $(dirname $0)
PROG_DIR=$(pwd); cd $WORK_DIR

BASE_DIR=$PROG_DIR/../../../app
DATA_DIR=$PROG_DIR/../../../app/library/test/mock

cd $BASE_DIR

dirs=
dirs+=" interfaces"
dirs+=" usecase"
dirs+=" infrastructure"

for target_dir in $dirs; do
    SEARCH_DIR=$target_dir
    src_files=$(find $SEARCH_DIR -type f -name *.go -not -name '*_test.go')
    for src_file in ${src_files[@]}; do
        grep "type.*interface" $src_file > /dev/null
        if [ $? -ne 0 ]; then continue; fi
        echo $src_file
        mkdir -p $DATA_DIR/$(dirname $src_file)
        mockgen -source $src_file > $DATA_DIR/$src_file &
    done
done

