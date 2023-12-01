#!/bin/bash
set -e;

add_input() {
    cd inputs;
    CURRENT=$(($(ls -la | wc -l) -  3));
    DIRNAME="day$((CURRENT + 1))";
    mkdir $DIRNAME;
    cd $DIRNAME;
    touch data.txt data1.txt;
    cd ../..;
}

create_package() {
    cd days;
    NAME="day$((CURRENT + 1))";
    mkdir $NAME;
    cd $NAME;
    FILE="solver.go"
    cp ../../raw/raw.go $FILE;
    sed -i "s/_raw_/$NAME/g" $FILE;
    cd ../..;
}

add_test() {
    cd days/$NAME;
    FILE="solver_test.go"
    cp ../../raw/raw_test.go $FILE;
    sed -i "s/_raw_/$NAME/g" $FILE;
    sed -i "s/var day = 0/var day = $((CURRENT + 1))/g" $FILE;
    cd ../..;
}

add_import() {
    LINE="\"aoc_2023/days/day$CURRENT\"";
    NEXT_LINE="\\\\t\"aoc_2023/days/day$((CURRENT + 1))\"";
    sed -i "\-$LINE-a $NEXT_LINE" main.go;
}

use_solver() {
    CURRENT_DAY="day$CURRENT";
    LINE="return &$CURRENT_DAY.Solver{}";
    NEXT_DAY=$((CURRENT + 1));
    NEXT_LINE_1="\\\\tcase $NEXT_DAY:";
    NEXT_LINE_2="\\t\\treturn &day$NEXT_DAY.Solver{}";
    sed -i "/$LINE/a $NEXT_LINE_1\\n$NEXT_LINE_2" main.go
}


add_input;
create_package;
add_test;
add_import;
use_solver;
echo "Succesfully created new day";
