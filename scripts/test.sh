#!/usr/bin/env bash 
set -e

optspec=":hv-:"
while getopts "$optspec" optchar; do
    case "${optchar}" in

        -)
            case "${OPTARG}" in
                version)
                    val="${!OPTIND}"; OPTIND=$(( $OPTIND + 1 ))
                    #echo "Parsing option: '--${OPTARG}', value: '${val}'" >&2;
                    version=$val
                    ;;
                #loglevel=*)
                    #val=${OPTARG#*=}
                    #opt=${OPTARG%=$val}
                    #echo "Parsing option: '--${opt}', value: '${val}'" >&2
                    #;;
                *)
                    if [ "$OPTERR" = 1 ] && [ "${optspec:0:1}" != ":" ]; then
                      echo "Unknown option --${OPTARG}" >&2
                    fi
                    ;;
            esac;;
        h)
          echo "usage: $0 [-v <value>] or [--version <value>]" >&2
            exit 2
            ;;
        v)
            #echo "Parsing option: '-${optchar}'" >&2
            val="${!OPTIND}"; OPTIND=$(( $OPTIND + 1 ))
            version=$val
            ;;
        *)
            if [ "$OPTERR" != 1 ] || [ "${optspec:0:1}" = ":" ]; then
                echo "Non-option argument: '-${OPTARG}'" >&2
            fi
            ;;
    esac
done

docker build --force-rm -t libdna:$version `pwd`
docker run --rm libdna:$version sh -c "go test github.com/wmiller848/libdna/..."
