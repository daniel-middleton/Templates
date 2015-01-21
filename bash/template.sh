#!/bin/bash

###############################################################################
# Name:         Replace
# Author:       Daniel Middleton <daniel-middleton.com>
# Description:  Replace
# Usage:        Replace
###############################################################################

# Global vars
PROG_NAME="$(basename $0)"

# Usage: screenOut STATUS message
screenOut() {
    timestamp=$(date +"%H:%M:%S")
    
    if [ "$#" -ne 2 ]; then
        status='INFO'
        message="$1"
    else
        status="$1"
        message="$2"
    fi

    echo -e "[$PROG_NAME][$timestamp][$status]: $message"
    logger -t $PROG_NAME "[$status]: $message"
}

# Usage: checkStatus $? "Error message" "Success message"
checkStatus() {
    case $1 in
        0)
            screenOut "SUCCESS" "$3"
            ;;
        1)
            screenOut "ERROR" "$2 - Exiting..."
            exit 1
            ;;
        *)
            screenOut "ERROR" "Unhandled return code. Exiting..."
            exit 1
            ;;
    esac
}

# Usage: 
newMethod() {
    screenOut "Inside new method!"
    echo "New Method!"
    checkStatus $? "Custom error message." \
                   "Custom success message."
}

echo && screenOut "$PROG_NAME script started..."
screenOut "Logging to syslog and stdout."

# Execute methods
newMethod

#End
screenOut "$PROG_NAME script ended." && echo
exit 0
