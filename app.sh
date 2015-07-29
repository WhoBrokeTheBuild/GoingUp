#!/bin/bash

case $1 in 
    "docker-build" )
    
        docker build -t goingup .
    
        ;;
    "docker-stop" )
    
        printf "Killing..\n"
        docker kill goingup-1 goingup-2 goingup-3 hap-goingup
        printf "Done\n\n"
        
        printf "Removing..\n"
        docker rm goingup-1 goingup-2 goingup-3 hap-goingup
        printf "Done\n\n"
        
        ;;
    "docker-run" )
    
        SCRIPT=$(readlink -f "$0")
        SCRIPTPATH=$(dirname "$SCRIPT")
        
        printf "Killing..\n"
        docker kill goingup-1 goingup-2 goingup-3 hap-goingup
        printf "Done\n\n"
        
        printf "Removing..\n"
        docker rm goingup-1 goingup-2 goingup-3 hap-goingup
        printf "Done\n\n"
        
        printf "Starting..\n"
        printf "goingup-1 "
        docker run -d  \
            -v $SCRIPTPATH/goingup-example/static:/go/bin/static:ro \
            -v $SCRIPTPATH/goingup-example/templates:/go/bin/templates:ro \
            --name goingup-1 goingup
            
        printf "goingup-2 "
        docker run -d  \
            -v $SCRIPTPATH/goingup-example/static:/go/bin/static:ro \
            -v $SCRIPTPATH/goingup-example/templates:/go/bin/templates:ro \
            --name goingup-2 goingup
        printf "goingup-3 "
        docker run -d  \
            -v $SCRIPTPATH/goingup-example/static:/go/bin/static:ro \
            -v $SCRIPTPATH/goingup-example/templates:/go/bin/templates:ro \
            --name goingup-3 goingup
        
        printf "hap-goingup "
        docker run -d -p 8080:8080 \
            -v $SCRIPTPATH/hap-goingup.cfg:/usr/local/etc/haproxy/haproxy.cfg:ro \
            --link goingup-1:goingup-1 \
            --link goingup-2:goingup-2 \
            --link goingup-3:goingup-3 \
            --name hap-goingup haproxy:1.5.14
        printf "Done\n\n"

        ;;
    "build" )
    
        cd goingup-example
        go build
        
        ;;
    "run" )
    
        cd goingup-example
        ./goingup-example
    
        ;;
esac