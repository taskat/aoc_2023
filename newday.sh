#!/bin/bash
set -e;

add_input() {
    cd inputs;
    CURRENT=$(($(ls -la | wc -l) -  3));
    DIRDAYNUMBER="day$((CURRENT + 1))";
    mkdir $DIRDAYNUMBER;
    cd $DIRDAYNUMBER;
    touch data.txt data1.txt;
    cd ../..;
}

create_package() {
    cd days;
    DAYNUMBER="$((CURRENT + 1))";
    mkdir day$DAYNUMBER;
    cd day$DAYNUMBER;
    FILE="solver.go"
    cp ../../raw/raw.go $FILE;
    sed -i "s/_daynumber_/$DAYNUMBER/g" $FILE;
    cd ../..;
}

add_test() {
    cd days/day$DAYNUMBER;
    FILE="solver_test.go"
    cp ../../raw/raw_test.go $FILE;
    sed -i "s/_daynumber_/$DAYNUMBER/g" $FILE;
    sed -i "s/day    = 0/day    = $((CURRENT + 1))/g" $FILE;
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
