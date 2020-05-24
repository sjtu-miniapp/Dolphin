#!/bin/bash

do_usage() {
    if [ $# -eq 1 ]; then
        echo ${1}
    fi
    echo 
    cat >&2 <<EOF
Usage:
./execute.sh hostname case [Options]

  hostname: Target host url of sockshop, e.g. http://localhost:30000
  case: test case, e.g. case1

Options:
  -l stress load
  -t run time

Description:
  Runs a Locust load simulation against specified host.

EOF
  exit 1
}

do_correctness() {
    locust --host=$TARGET_HOST -f correctness.py --clients=2 --hatch-rate=5 --run-time=15s --no-web --only-summary
}

do_loadtest() {
    INSTANCE_CNT=$(( ($STRESS+49)/50 ))
    if [ $INSTANCE_CNT -eq 0 ]; then
        INSTANCE_CNT=1
    fi
    PERLOCUST_CLIENT_CNT=$(( ($STRESS + $INSTANCE_CNT - 1) / ($INSTANCE_CNT) ))
    LASTLOCUST_CLIENT_CNT=$(( $STRESS - $PERLOCUST_CLIENT_CNT*($INSTANCE_CNT-1) ))
    for (( i=0; i<$(( $INSTANCE_CNT-1 )); i++ )) ; do
        locust --host=$TARGET_HOST -f background-stress.py --clients $PERLOCUST_CLIENT_CNT --hatch-rate=1000000000 --run-time=$RUNTIME  --no-web --only-summary 2>background-result$i.txt &
    done
    locust --host=$TARGET_HOST -f background-stress.py --clients $LASTLOCUST_CLIENT_CNT --hatch-rate=1000000000 --run-time=$RUNTIME  --no-web --only-summary 2>background-result$(( $INSTANCE_CNT - 1 )).txt &
    locust --host=$TARGET_HOST -f $CASE.py --clients 20 --hatch-rate=1000000000 --run-time=$RUNTIME --no-web --only-summary
}

do_verbose() {
    wait < <(jobs -p)
    INSTANCE_CNT=$(( ($STRESS+49)/50 ))
    if [ $INSTANCE_CNT -eq 0 ]; then
        INSTANCE_CNT=1
    fi
    for (( i=0; i<$(( $INSTANCE_CNT )); i++ )) ; do
        cat background-result$i.txt
    done
}

STRESS=20
RUNTIME=30

while [ $# -gt 0 ] ; do
    OPTIND=1
    while getopts ":l:t:h:v" o; do
        case "${o}" in
            h)
                do_usage
                ;;
            l)
                STRESS=${OPTARG}
                ;;
            t)
                RUNTIME=${OPTARG}
                ;;
            v)
                VERBOSE=true
                ;;
            *)
                do_usage "Unrecognized option: ${OPTARG}"
                ;;
        esac
    done
    shift $((OPTIND-1))
    if [ $# -eq 0 ]; then
        break
    fi

    if [ -z $TARGET_HOST ]; then
        TARGET_HOST=${1}
    else 
        if [ -z $CASE ]; then
            CASE=${1}
        else
            do_usage "Too much arguments: ${1}"
        fi
    fi
    shift 1
done

if [ -z $TARGET_HOST ] || [ -z $CASE ]; then
    do_usage
fi

case $CASE in
    case1)
        do_loadtest
        ;;
    case2)
        do_loadtest
        ;;
    correctness)
        do_correctness
        ;;
    *)
        do_usage "Unrecognized test case: $CASE"
        ;;
esac

if [ ! -z $VERBOSE ]; then
    do_verbose
fi

exit 0
