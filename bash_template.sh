#!/bin/bash

###############################################################################
# Name: Replace
# Author:   Daniel Middleton <daniel-middleton.com>
# Description: Replace
# Usage: Replace
###############################################################################

# Global vars
PROG_NAME='Replace'

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

    echo -e "[$PROG_NAME][$status][$timestamp]: $message"
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
            screenOut "ERROR" "Error in checkStatus function. Exiting..."
            exit 1
            ;;
    esac
}

# New functions
newMethod() {
    screenOut "Inside new method!"
    echo "New Method!"
    checkStatus $? "Custom error message." \
                   "Custom success message."
}

echo && screenOut "$PROG_NAME script started..."

# Execute methods
newMethod

#End
screenOut "$PROG_NAME script ended." && echo
exit 0
